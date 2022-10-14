package routes

import (
	"encoding/json"
	"errors"
	"github.com/everFinance/goar/types"
	"github.com/everFinance/goar/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/warp-contracts/sequencer/ar"
	"github.com/warp-contracts/sequencer/config"
	"github.com/warp-contracts/sequencer/db/interactiondb"
	"github.com/warp-contracts/sequencer/db/sequencerdb"
	"github.com/warp-contracts/sequencer/sortkey"
	"github.com/warp-contracts/sequencer/tagHelper"
	"net/http"
	"regexp"
	"strconv"
	"sync"
	"time"
)

func RegisterSequencer(c *gin.Context) {
	transaction := new(types.Transaction)
	err := c.BindJSON(transaction)
	//utils.VerifyTransaction(*transaction)

	if checkError(c, err, http.StatusBadRequest) {
		return
	}

	cachedNetworkData := ar.GetCachedInfo()
	jwk := config.GetArConnectAsJwkKey()
	if checkError(c, err, http.StatusInternalServerError) {
		return
	}
	if checkError(c, err, http.StatusInternalServerError) {
		return
	}
	originalOwner := transaction.Owner
	originalAddress, err := utils.OwnerToAddress(originalOwner)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	millis := time.Now().UnixMilli()
	currentHeight := cachedNetworkData.NetworkInfo.Height
	currentBlockId := cachedNetworkData.NetworkInfo.Current
	sortKey, err := sortkey.CreateSortKey(jwk, []byte(currentBlockId), millis, []byte(transaction.ID), currentHeight)
	if checkError(c, err, http.StatusInternalServerError) {
		return
	}
	//print(sortKey, originalAddress)
	contractTag, inputTag, internalWrites, decodedTags, tags, vrfData, err := tagHelper.PrepareTags(
		transaction,
		originalAddress,
		millis,
		sortKey,
		currentHeight,
		currentBlockId,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	bundlrResp, err := ar.GetBundlr().UploadToBundlr(
		transaction,
		tags...,
	)
	if checkError(c, err, http.StatusInternalServerError) {
		return
	}

	functionInput, err := parseFunctionInput(inputTag)
	if err != nil || functionInput.Function == "" {
		logrus.WithField("input", inputTag).
			Error("Could not parse function input", err)
	}
	var evolve string
	if functionInput != nil &&
		functionInput.Function == "evolve" &&
		functionInput.Value != "" &&
		isTxIdValid(functionInput.Value) {
		evolve = functionInput.Value
	}

	interaction := createInteraction(
		transaction,
		originalAddress,
		decodedTags,
		currentHeight,
		currentBlockId,
		cachedNetworkData.CurrentBlock,
		sortKey,
		vrfData,
	)

	bundlerRespJson, err := json.Marshal(bundlrResp)
	if checkError(c, err, http.StatusInternalServerError) {
		return
	}

	interactionJson, err := json.Marshal(interaction)
	if checkError(c, err, http.StatusInternalServerError) {
		return
	}

	errs := saveResultsInDb(transaction, originalOwner, originalAddress, currentBlockId, currentHeight, millis, sortKey, bundlrResp, bundlerRespJson, interactionJson, contractTag, functionInput, inputTag, internalWrites, evolve)

	if len(errs) > 0 {
		var msg string
		for _, e := range errs {
			logrus.Error(e)
			msg += e.Error() + "\n"
		}
		checkError(c, errors.New(msg), http.StatusInternalServerError)
		return
	}

	c.JSON(200, bundlrResp)
}

func saveResultsInDb(transaction *types.Transaction, originalOwner string, originalAddress string, currentBlockId string, currentHeight int64, millis int64, sortKey string, bundlrResp *types.BundlrResp, bundlerRespJson []byte, interactionJson []byte, contractTag string, functionInput *FunctionInput, inputTag string, internalWrites string, evolve string) []error {
	var wg sync.WaitGroup
	wg.Add(2)
	var lock sync.Mutex
	var errs []error
	go func() {
		err := sequencerdb.Save(&sequencerdb.Sequence{
			OriginalSig:           transaction.Signature,
			OriginalOwner:         originalOwner,
			OriginalAddress:       originalAddress,
			SequenceBlockId:       currentBlockId,
			SequenceBlockHeight:   currentHeight,
			SequenceTransactionId: transaction.ID,
			SequenceMillis:        strconv.FormatInt(millis, 10),
			SequenceSortKey:       sortKey,
			BundlerTxId:           bundlrResp.Id,
			BundlerResponse:       string(bundlerRespJson),
		})
		if err != nil {
			lock.Lock()
			errs = append(errs, err)
			lock.Unlock()
		}
		wg.Done()
	}()
	go func() {
		err := interactiondb.Save(&interactiondb.Interaction{
			InteractionId:      transaction.ID,
			Interaction:        string(interactionJson),
			BlockHeight:        currentHeight,
			BlockId:            currentBlockId,
			ContractId:         contractTag,
			Function:           functionInput.Function,
			Input:              inputTag,
			ConfirmationStatus: "confirmed",
			ConfirmingPeer:     viper.GetString("arweave.bundlrUrls"),
			Source:             "redstone-sequencer",
			BundlerTxId:        bundlrResp.Id,
			InteractWrite:      []string{internalWrites},
			SortKey:            sortKey,
			Evolve:             evolve,
		})
		if err != nil {
			lock.Lock()
			errs = append(errs, err)
			lock.Unlock()
		}
		wg.Done()
	}()
	wg.Wait()
	return errs
}

func checkError(c *gin.Context, err error, returnCode int) bool {
	if err != nil {
		if returnCode > 499 && returnCode < 600 {
			logrus.Error(err)
		} else {
			logrus.Debug(err)
		}
		c.JSON(returnCode, err)
		return true
	}
	return false
}

func createInteraction(transaction *types.Transaction,
	originalAddress string,
	decodedTags []types.Tag,
	height int64,
	currentBlockId string,
	blockData *types.Block,
	sortKey string,
	vrfData tagHelper.VrfData,
) *Interaction {
	return &Interaction{
		Id:        transaction.ID,
		Owner:     owner{Address: originalAddress},
		Recipient: transaction.Target,
		Tags:      decodedTags,
		Block: block{
			Height:    height,
			Id:        currentBlockId,
			Timestamp: blockData.Timestamp,
		},
		Fee: fee{
			Winston: transaction.Reward,
		},
		Quantity: quantity{
			Winston: transaction.Quantity,
		},
		SortKey: sortKey,
		Source:  "redstone-sequencer",
		Vrf:     vrfData,
	}
}

func isTxIdValid(txId string) bool {
	r, _ := regexp.Compile("/[a-z0-9_-]{43}/i")
	return r.Match([]byte(txId))
}

func parseFunctionInput(input string) (functionInput *FunctionInput, err error) {
	err = json.Unmarshal([]byte(input), &functionInput)
	return
}

type FunctionInput struct {
	Function string
	Value    string
}

type Interaction struct {
	Id        string            `json:"id"`
	Owner     owner             `json:"owner"`
	Recipient string            `json:"recipient"`
	Tags      []types.Tag       `json:"tags"`
	Block     block             `json:"block"`
	Fee       fee               `json:"fee"`
	Quantity  quantity          `json:"quantity"`
	SortKey   string            `json:"sortkey"`
	Source    string            `json:"source"`
	Vrf       tagHelper.VrfData `json:"vrf"`
}
type owner struct {
	Address string `json:"address"`
}
type block struct {
	Height    int64  `json:"height"`
	Id        string `json:"id"`
	Timestamp int64  `json:"timestamp"`
}
type fee struct {
	Winston string `json:"winston"`
}
type quantity struct {
	Winston string `json:"winston"`
}

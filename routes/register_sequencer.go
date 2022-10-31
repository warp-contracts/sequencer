package routes

import (
	"encoding/json"
	"errors"
	"github.com/everFinance/goar/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/warp-contracts/sequencer/ar"
	"github.com/warp-contracts/sequencer/config"
	"github.com/warp-contracts/sequencer/db/interactiondb"
	"github.com/warp-contracts/sequencer/db/sequencerdb"
	"github.com/warp-contracts/sequencer/measure"
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

	start := time.Now()

	jwk := config.GetArConnectAsJwkKey()
	if checkError(c, err, http.StatusInternalServerError) {
		return
	}
	if checkError(c, err, http.StatusInternalServerError) {
		return
	}
	originalOwner := transaction.Owner
	if checkError(c, err, http.StatusBadRequest) {
		return
	}

	millis := time.Now().UnixMilli()
	currentHeight := cachedNetworkData.NetworkInfo.Height
	currentBlockId := cachedNetworkData.NetworkInfo.Current
	sortKey, err := sortkey.CreateSortKey(jwk, []byte(currentBlockId), millis, []byte(transaction.ID), currentHeight)
	if checkError(c, err, http.StatusInternalServerError) {
		return
	}

	contractTag, inputTag, originalAddress, internalWrites, decodedTags, tags, vrfData, isEvmSigner, err := tagHelper.PrepareTags(
		transaction,
		originalOwner,
		millis,
		sortKey,
		currentHeight,
		currentBlockId,
	)
	if checkError(c, err, http.StatusBadRequest) {
		return
	}
	if inputTag == "" {
		checkError(c, errors.New("input tag is required"), http.StatusBadRequest)
	}

	startBundlrUpload := time.Now()
	bundlrResp, confirmPeer, err := ar.GetBundlr().UploadToBundlr(
		transaction,
		tags...,
	)
	measure.LogDurationFrom(logrus.DebugLevel, startBundlrUpload, "Uploading to bundlr")
	if checkError(c, err, http.StatusInternalServerError) {
		return
	}
	logrus.Debugf("Bundlr response id %s", bundlrResp.Id)

	functionInput, err := parseFunctionInput(inputTag)
	if checkError(c, err, http.StatusInternalServerError) {
		return
	}
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

	var sign string
	if isEvmSigner {
		sign = transaction.Signature
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
		sign,
	)

	bundlerRespJson, err := json.Marshal(bundlrResp)
	if checkError(c, err, http.StatusInternalServerError) {
		return
	}

	interactionJson, err := json.Marshal(interaction)
	if checkError(c, err, http.StatusInternalServerError) {
		return
	}

	sequence := &sequencerdb.Sequence{
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
	}
	interact := &interactiondb.Interaction{
		InteractionId:      transaction.ID,
		Interaction:        string(interactionJson),
		BlockHeight:        currentHeight,
		BlockId:            currentBlockId,
		ContractId:         contractTag,
		Function:           functionInput.Function,
		Input:              inputTag,
		ConfirmationStatus: "confirmed",
		ConfirmingPeer:     confirmPeer,
		Source:             "redstone-sequencer",
		BundlerTxId:        bundlrResp.Id,
		InteractWrite:      internalWrites,
		SortKey:            sortKey,
		Evolve:             evolve,
	}

	errs := saveResultsInDb(sequence, interact)

	if len(errs) > 0 {
		var msg string
		for _, e := range errs {
			logrus.Error(e)
			logrus.Error(transaction)
			logrus.Error(tags)
			logrus.Error(bundlrResp)
			msg += e.Error() + "\n"
		}
		checkError(c, errors.New(msg), http.StatusInternalServerError)
		return
	}

	c.JSON(200, bundlrResp)
	measure.LogDurationFrom(logrus.InfoLevel, start, "Total sequencer processing")
}

func saveResultsInDb(sequence *sequencerdb.Sequence, interaction *interactiondb.Interaction) []error {
	var wg sync.WaitGroup
	wg.Add(2)
	var lock sync.Mutex
	var errs []error
	start := time.Now()
	go func() {
		defer wg.Done()
		err := sequencerdb.Save(sequence)
		if err != nil {
			lock.Lock()
			errs = append(errs, err)
			lock.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err := interactiondb.Save(interaction)
		if err != nil {
			lock.Lock()
			errs = append(errs, err)
			lock.Unlock()
		}
	}()
	wg.Wait()
	measure.LogDurationFrom(logrus.InfoLevel, start, "Inserting into tables")
	return errs
}

func checkError(c *gin.Context, err error, returnCode int) bool {
	if err != nil {
		if returnCode > 499 && returnCode < 600 {
			logrus.Error(err)
		} else {
			logrus.Info(err)
		}
		c.JSON(returnCode, err.Error())
		return true
	}
	return false
}

func createInteraction(
	transaction *types.Transaction,
	originalAddress string,
	decodedTags []types.Tag,
	height int64,
	currentBlockId string,
	blockData *types.Block,
	sortKey string,
	vrfData tagHelper.VrfData,
	signature string,
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
		SortKey:   sortKey,
		Source:    "redstone-sequencer",
		Vrf:       vrfData,
		Signature: signature,
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
	Signature string            `json:"signature"`
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

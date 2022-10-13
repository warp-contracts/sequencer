package routes

import (
	"encoding/json"
	"github.com/everFinance/goar/types"
	"github.com/everFinance/goar/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/warp-contracts/sequencer/ar"
	"github.com/warp-contracts/sequencer/ar/tagHelper"
	"github.com/warp-contracts/sequencer/config"
	"github.com/warp-contracts/sequencer/db/interactiondb"
	"github.com/warp-contracts/sequencer/db/sequencerdb"
	"github.com/warp-contracts/sequencer/sortkey"
	"net/http"
	"regexp"
	"sync"
	"time"
)

//arweave.net/{txId}

func RegisterSequencer(c *gin.Context) {
	transaction := new(types.Transaction)
	// I know that it wasn't done in the original version
	// but it would be wise to add transaction validation here
	// i.e. https://github.com/everFinance/goar/blob/main/utils/transaction.go#L143

	// what is the point of this operation?
	err := c.BindJSON(transaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	cachedNetworkData := ar.GetCachedInfo()
	jwk := config.GetArConnectAsJwkKey()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	originalOwner := transaction.Owner
	originalAddress, err := utils.OwnerToAddress(originalOwner)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// we need to be super-careful here when this code will be deployed on prod
	// ie be sure that the timezone etc is the same as in the current - js version.
	millis := time.Now().UnixMilli()
	currentHeight := cachedNetworkData.NetworkInfo.Height
	currentBlockId := cachedNetworkData.NetworkInfo.Current
	sortKey, err := sortkey.CreateSortKey(jwk, []byte(currentBlockId), millis, []byte(transaction.ID), currentHeight)
	// shouldn't we - in general - log errors here?
	// at least in the beginning state of running of the new version
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
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

	// please log here with warn if senging to bundlr took longer than 500ms
	bundlrResp, err := ar.GetBundlr().UploadToBundlr(
		transaction,
		tags...,
	)
	if err != nil {
		// please log here with some placeholder that would be easy to search
		c.JSON(http.StatusInternalServerError, err)
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
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	interactionJson, err := json.Marshal(interaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		sequencerdb.Save(&sequencerdb.Sequence{
			OriginalSig:           transaction.Signature,
			OriginalOwner:         originalOwner,
			OriginalAddress:       originalAddress,
			SequenceBlockId:       currentBlockId,
			SequenceBlockHeight:   currentHeight,
			SequenceTransactionId: transaction.ID,
			SequenceMillis:        millis,
			SequenceSortKey:       sortKey,
			BundlerTxId:           bundlrResp.Id,
			BundlerResponse:       string(bundlerRespJson),
		})
		wg.Done()
	}()
	go func() {
		interactiondb.Save(&interactiondb.Interaction{
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
			InteractWrite:      internalWrites,
			SortKey:            sortKey,
			Evolve:             evolve,
		})
		wg.Done()
	}()
	wg.Wait()

	c.JSON(200, bundlrResp)
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
		Owner:     struct{ address string }{address: originalAddress},
		Recipient: transaction.Target,
		Tags:      decodedTags,
		Block: struct {
			Height    int64
			Id        string
			Timestamp int64
		}{
			Height:    height,
			Id:        currentBlockId,
			Timestamp: blockData.Timestamp,
		},
		Fee: struct {
			Winston string
		}{
			Winston: transaction.Reward,
		},
		Quantity: struct {
			Winston string
		}{
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
	Id        string
	Owner     struct{ address string }
	Recipient string
	Tags      []types.Tag
	Block     struct {
		Height    int64
		Id        string
		Timestamp int64
	}
	Fee      struct{ Winston string }
	Quantity struct {
		Winston string
	}
	SortKey string
	Source  string
	Vrf     tagHelper.VrfData
}

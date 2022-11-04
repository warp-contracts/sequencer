package tagHelper

import (
	"encoding/base64"
	"encoding/hex"
	"github.com/everFinance/goar/types"
	"github.com/everFinance/goar/utils"
	"github.com/warp-contracts/sequencer/ar/smartweave"
	"github.com/warp-contracts/sequencer/crypt"
	"math/big"
	"strconv"
)

func PrepareTags(transaction *types.Transaction,
	originalOwner string,
	millis int64,
	sortKey string,
	currentHeight int64,
	currentBlockId string) (contractTag, inputTag, originalAddress string,
	internalWrites []string,
	decodedTags, tags []types.Tag,
	vrfData VrfData,
	isEvmSigner bool,
	testnetVersion string,
	err error) {
	decodedTags, err = utils.TagsDecode(transaction.Tags)
	if err != nil {
		return
	}
	tags = []types.Tag{
		{
			Name:  "Sequencer",
			Value: "RedStone",
		},
		{
			Name:  "Sequencer-Owner",
			Value: originalOwner,
		},
		{
			Name:  "Sequencer-Mills",
			Value: strconv.FormatInt(millis, 10),
		},
		{
			Name:  "Sequencer-Sort-Key",
			Value: sortKey,
		},
		{
			Name:  "Sequencer-Tx-Id",
			Value: transaction.ID,
		},
		{
			Name:  "Sequencer-Block-Height",
			Value: strconv.FormatInt(currentHeight, 10),
		},
		{
			Name:  "Sequencer-Block-Id",
			Value: currentBlockId,
		},
	}

	var requestVrfTag = false
	for _, tag := range decodedTags {
		switch tag.Name {
		case smartweave.TagContractTxId:
			contractTag = tag.Value
		case smartweave.TagInput:
			inputTag = tag.Value
		case smartweave.TagInteractWrite:
			internalWrites = append(internalWrites, tag.Value)
		case smartweave.TagRequestVrf:
			requestVrfTag = true
		case smartweave.TagSignatureType:
			if tag.Value == "ethereum" {
				originalAddress = originalOwner
				isEvmSigner = true
			}
		case smartweave.TagWarpTestnet:
			testnetVersion = tag.Value
		}
	}
	if originalAddress == "" {
		originalAddress, err = utils.OwnerToAddress(originalOwner)
		if err != nil {
			return
		}
	}

	tags = append(tags, decodedTags...)
	if requestVrfTag {
		var vrfTags []types.Tag
		vrfTags, vrfData = getVrfTags(sortKey)
		tags = append(tags, vrfTags...)
	}
	return
}

func getVrfTags(sortKey string) (vrfTags []types.Tag, vrfData VrfData) {
	k := crypt.GetKey()
	index, proof := k.Evaluate([]byte(sortKey))
	arrayIndex := index[:]
	vrfData = VrfData{
		Index:  base64.RawURLEncoding.EncodeToString(arrayIndex),
		Proof:  base64.RawURLEncoding.EncodeToString(proof[:]),
		Bigint: indexToBigint(arrayIndex).String(),
		Pubkey: crypt.GetCompactPublicKey(k),
	}
	vrfTags = []types.Tag{
		{
			Name:  "vrf-index",
			Value: vrfData.Index,
		}, {
			Name:  "vrf-proof",
			Value: vrfData.Proof,
		}, {
			Name:  "vrf-bigint",
			Value: vrfData.Bigint,
		}, {
			Name:  "vrf-pubkey",
			Value: vrfData.Pubkey,
		},
	}
	return
}

func indexToBigint(index []byte) *big.Int {
	b := new(big.Int)
	b.SetString(hex.EncodeToString(index), 16)
	return b
}

type VrfData struct {
	Index  string `json:"index"`
	Proof  string `json:"proof"`
	Bigint string `json:"bigint"`
	Pubkey string `json:"pubkey"`
}

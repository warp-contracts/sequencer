package ar

import (
	"github.com/everFinance/goar/types"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/warp-contracts/sequencer/config"
	"os"
	"testing"
)

func TestBundlr(t *testing.T) {
	t.Parallel()

	key, err := os.ReadFile("../_tests/arweavekeys/5SUBakh_R97MbHoX0_wNarVUw6DH0TziW5rG2K1vc6k.json")
	assert.NoError(t, err)
	viper.Set("arweave.walletJwk", key)

	config.Init()

	bundl := GetBundlr()
	t.Run("should return non-nil bundlr", func(t *testing.T) {
		t.Parallel()
		assert.NotNil(t, bundl)
	})

	t.Run("should successfully UploadToBundlr", func(t *testing.T) {
		t.Parallel()

		assert.NoError(t, err)
		transaction := &types.Transaction{}
		resp, confirmNode, err := bundl.UploadToBundlr(transaction)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.NotEmpty(t, resp.Id)
		assert.NotEmpty(t, resp.Signature)
		assert.NotEmpty(t, resp.Block)
		assert.NotEmpty(t, resp.Public)
		assert.True(t, "https://node.bundlr.network" == confirmNode ||
			"https://node2.bundlr.network" == confirmNode)
	})

	t.Run("should return information about correct confirm peer", func(t *testing.T) {
		viper.Set("arweave.bundlrUrls", []string{
			"some-unexisting-node",
			"https://node2.bundlr.network",
			"https:/google.com",
		})
		_, confirmNode, err := bundl.UploadToBundlr(&types.Transaction{})
		assert.NoError(t, err)
		assert.True(t, "https://node2.bundlr.network" == confirmNode)
	})

}

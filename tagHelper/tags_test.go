package tagHelper

import (
	"crypto/ecdsa"
	"encoding/base64"
	"github.com/everFinance/goar"
	"github.com/everFinance/goar/types"
	"github.com/everFinance/goar/utils"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/warp-contracts/sequencer/ar/smartweave"
	"github.com/warp-contracts/sequencer/config"
	"github.com/warp-contracts/sequencer/crypt"
	"github.com/warp-contracts/sequencer/crypt/p256/secp256k1/vrf/secp256k1VRF"
	"os"
	"strings"
	"testing"
)

func TestPrepareTags(t *testing.T) {
	t.Parallel()

	kBytes, err := os.ReadFile("../_tests/arweavekeys/vrf-example")
	assert.NoError(t, err)
	key, err := crypt.UnmarshalKey(string(kBytes))
	assert.NoError(t, err)

	viper.Set("vrf.privateKey", crypt.MarshalKey(key))
	config.Init()

	t.Run("should create tags", func(t *testing.T) {
		t.Parallel()

		originalOwner := "addr"
		var millis int64 = 123
		sortKey := "000001026265,1664446281798,dd9a9dc0d898a93bb00e278d4c7fa8840fa3a04363c7ae4089b2c3d1ac56ecad"
		sourceTags := []types.Tag{
			{
				Name:  "transaction tag name 1",
				Value: "transaction tag value 1",
			},
			{
				Name:  "transaction tag name 2",
				Value: "transaction tag value 2",
			},
			{
				Name:  smartweave.TagContractTxId,
				Value: "contract tag value",
			},
			{
				Name:  smartweave.TagInput,
				Value: "input tag value",
			},
			{
				Name:  smartweave.TagInteractWrite,
				Value: "internalWrites tag value",
			},
			{
				Name:  smartweave.TagRequestVrf,
				Value: "true",
			},
		}
		transaction := &types.Transaction{
			ID:   "tx id",
			Tags: utils.TagsEncode(sourceTags),
		}
		var currentHeight int64 = 123123
		currentBlockId := "qweasd"
		contractTag, inputTag, originalAddress, internalWrites, decodedTags, tags, vrfData, err := PrepareTags(
			transaction,
			originalOwner,
			millis,
			sortKey,
			currentHeight,
			currentBlockId,
		)
		assert.NoError(t, err)

		print(contractTag, inputTag, internalWrites, tags, &vrfData)
		t.Run("Tags", func(t *testing.T) {
			t.Parallel()

			t.Run("should create non-empty tags", func(t *testing.T) {
				assert.NotEmpty(t, tags)
			})

			t.Run("should contain sequencer tags", func(t *testing.T) {
				t.Parallel()

				assertTag(t, tags, "Sequencer", "RedStone")
				assertTag(t, tags, "Sequencer-Owner", originalOwner)
				assertTag(t, tags, "Sequencer-Mills", "123")
				assertTag(t, tags, "Sequencer-Sort-Key", sortKey)
				assertTag(t, tags, "Sequencer-Tx-Id", transaction.ID)
				assertTag(t, tags, "Sequencer-Block-Height", "123123")
				assertTag(t, tags, "Sequencer-Block-Id", currentBlockId)
			})

			t.Run("should contain vrf tags", func(t *testing.T) {
				t.Parallel()

				assertTag(t, tags, "vrf-index", "s7cDEZ5ZfkbkN0NfN5jsRMiKnMGN4IHtdG3Nr2QAYyU")

				vrfProof := findTag("vrf-proof", tags)
				assert.NotNil(t, vrfProof)
				assertTag(t, tags, "vrf-bigint", "81287354089493419043609974266174296544041075436869961695329565234059829863205")
				assertTag(t, tags, "vrf-pubkey", "03afb216678c386eeb1ceb0f5fdcfe7db3a9f7480ba5bd63695e5226f9bbb75b58")

				verifier, err := secp256k1VRF.NewVRFVerifier(key.Public().(*ecdsa.PublicKey))
				assert.NoError(t, err)

				decodedProof, err := base64.RawURLEncoding.DecodeString(vrfProof.Value)
				assert.NoError(t, err)
				hash, err := verifier.ProofToHash([]byte(sortKey), decodedProof)
				assert.NoError(t, err)
				assert.NotNil(t, hash)
				assert.Equal(t, "tzKvhpZV4CkcXosFdbLzBKG8rQvJxMSQkZer74fuRhc", originalAddress)
			})

			t.Run("should have transaction's tags", func(t *testing.T) {
				t.Parallel()
				for _, tag := range sourceTags {
					assertTag(t, decodedTags, tag.Name, tag.Value)
				}
			})
		})

		t.Run("return parameters", func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, contractTag, "contract tag value")
			assert.Equal(t, inputTag, "input tag value")
			assert.Equal(t, internalWrites, []string{"internalWrites tag value"})
			assert.Equal(t, vrfData, VrfData{
				Index:  findTag("vrf-index", tags).Value,
				Proof:  findTag("vrf-proof", tags).Value,
				Bigint: findTag("vrf-bigint", tags).Value,
				Pubkey: findTag("vrf-pubkey", tags).Value,
			})
		})
	})
	t.Run("should not return vrf tags when Request Vrf tag not exists", func(t *testing.T) {
		transaction := &types.Transaction{
			Tags: utils.TagsEncode([]types.Tag{}),
		}
		_, _, _, _, _, tags, _, err := PrepareTags(transaction, "", 1, "", 1, "")
		assert.NoError(t, err)

		for _, tag := range tags {
			assert.False(t, strings.HasPrefix(tag.Name, "vrf-"), "tag %s is unexpected", tag)
		}
	})
	t.Run("Signature-Type", func(t *testing.T) {
		t.Run("should return same address when Signature-Type equals ethereum", func(t *testing.T) {
			transaction := &types.Transaction{
				Tags: utils.TagsEncode([]types.Tag{
					{
						Name:  "Signature-Type",
						Value: "ethereum",
					},
				}),
			}
			_, _, originalAddress, _, _, _, _, err := PrepareTags(transaction, "owner", 1, "", 1, "")
			assert.NoError(t, err)
			assert.Equal(t, "owner", originalAddress)
		})
		t.Run("should convert owner to address when Signature-Type is not equals to ethereum", func(t *testing.T) {
			transaction := &types.Transaction{
				Tags: utils.TagsEncode([]types.Tag{
					{
						Name:  "Signature-Type",
						Value: "some-type",
					},
				}),
			}
			wallet, err := goar.NewWalletFromPath("../_tests/arweavekeys/5SUBakh_R97MbHoX0_wNarVUw6DH0TziW5rG2K1vc6k.json", viper.GetString("arweave.url"))
			assert.NoError(t, err)

			_, _, originalAddress, _, _, _, _, err := PrepareTags(transaction, wallet.Owner(), 1, "", 1, "")
			assert.NoError(t, err)
			ownerToAddress, err := utils.OwnerToAddress(wallet.Owner())
			assert.NoError(t, err)
			assert.Equal(t, ownerToAddress, originalAddress)
		})
	})
}

func assertTag(t *testing.T, tags []types.Tag, tagName string, expected string) {
	vrfPubkey := findTag(tagName, tags)
	assert.NotNil(t, vrfPubkey, "Tag %s not found", tagName)
	if vrfPubkey == nil {
		return
	}
	assert.Equal(t, expected, vrfPubkey.Value, "Wrong value for tag: %s", tagName)
}

func findTag(s string, tags []types.Tag) *types.Tag {
	for _, tag := range tags {
		if tag.Name == s {
			return &tag
		}
	}
	return nil
}

package sortkey

import (
	"fmt"
	"github.com/everFinance/goar/utils"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/stretchr/testify/assert"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

var lastNow int64
var blockHeightCount int64 = 0

func TestCreateSortKey(t *testing.T) {
	t.Parallel()

	t.Run("should create a non-empty key", func(t *testing.T) {
		key, err := CreateSortKey(newParams())
		assert.NoError(t, err)
		assert.NotEmpty(t, key)
	})
	t.Run("should generate different keys", func(t *testing.T) {
		key1, err := CreateSortKey(newParams())
		assert.NoError(t, err)
		key2, err := CreateSortKey(newParams())
		assert.NotEqual(t, key1, key2)
	})
	t.Run("should have 3 parts, split by comma", func(t *testing.T) {
		jwkKey, blockId, mills, transactionId, blockHeight := newParams()
		key, _ := CreateSortKey(jwkKey, blockId, mills, transactionId, blockHeight)
		splitKey := strings.Split(key, ",")
		assert.Len(t, splitKey, 3)
		t.Run("should have blockHeight in first part", func(t *testing.T) {
			keyPart := splitKey[0]
			blockHeightString := strconv.FormatInt(blockHeight, 10)
			t.Run("should end with blockHeight", func(t *testing.T) {
				assert.True(t, strings.HasSuffix(keyPart, blockHeightString))
			})
			t.Run("should have leading zeroes", func(t *testing.T) {
				requiredLen := 12
				assert.Len(t, keyPart, requiredLen)
				zeroes := generateMissedZeroes(requiredLen, blockHeightString)
				assert.Equal(t, zeroes+blockHeightString, keyPart)
			})
		})
		t.Run("should have mills as second part", func(t *testing.T) {
			i, err := strconv.ParseInt(splitKey[1], 10, 64)
			assert.NoError(t, err)
			assert.Equal(t, i, mills)
		})
		t.Run("should have hash in the third part", func(t *testing.T) {
			type testParam struct {
				testName      string
				key           jwk.Key
				blockId       []byte
				transactionId []byte
			}
			tests := []testParam{
				{
					testName:      "different jwk",
					key:           getJwkKey2(),
					blockId:       blockId,
					transactionId: transactionId,
				}, {
					testName:      "different blockId",
					key:           jwkKey,
					blockId:       []byte(strconv.Itoa(math.MaxInt)),
					transactionId: transactionId,
				}, {
					testName:      "different transactionId",
					key:           jwkKey,
					blockId:       blockId,
					transactionId: []byte(strconv.Itoa(math.MaxInt)),
				},
			}
			for _, test := range tests {
				t.Run(test.testName, func(t *testing.T) {
					sortKey, err := CreateSortKey(test.key, test.blockId, mills, test.transactionId, blockHeight)
					assert.NoError(t, err)
					splitTestKey := strings.Split(sortKey, ",")
					assert.NotEqual(t, splitKey[2], splitTestKey[2])
					assert.NotContains(t, splitTestKey[2], test.blockId)
				})
			}
		})
	})

	t.Run("should be able to generate same key with js implementations", func(t *testing.T) {
		originKey := "000001020352,1663684217401,1e8f524466584f490d9ca865a357e53b49ed064fe9416afe2ba338102c568509"
		var blockHeight int64 = 1020352
		blockId, err := utils.Base64Decode("KFOkPVliGG-KnunORRQVc1hj-OBQOIvc2g4tmaWmWYpz7jC9BwDlz4a7WD8ylqaU")
		assert.NoError(t, err)
		//blockId := []byte("KFOkPVliGG-KnunORRQVc1hj-OBQOIvc2g4tmaWmWYpz7jC9BwDlz4a7WD8ylqaU")
		var mills int64 = 1663684217401
		transactionId, err := utils.Base64Decode("6kPEEwd0GW_i1FVa6nA0hYYDFMYYckswkxTS4KZwNVg")
		assert.NoError(t, err)

		key, err := CreateSortKey(getJwkKey3(), blockId, mills, transactionId, blockHeight)
		assert.NoError(t, err)
		assert.Equal(t, originKey, key)
	})
}

//func TestQwe(t *testing.T) {
//wallet, _ := goar.NewWalletFromPath("../_tests/arweavekeys/5SUBakh_R97MbHoX0_wNarVUw6DH0TziW5rG2K1vc6k.json", "")
//d := wallet.Signer.PrvKey.D.Bytes()
//pem, err := jwk.EncodePEM(d)
//assert.NoError(t, err)
//print(base64.URLEncoding.EncodeToString(pem))
//}

func generateMissedZeroes(requiredLen int, blockHeight string) string {
	zeroes := make([]int, requiredLen-len(blockHeight))
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(zeroes)), ""), "[]")
}

func newParams() (jwk.Key, []byte, int64, []byte, int64) {
	now := time.Now().UnixMilli()
	if now == lastNow {
		now += 1
	}
	lastNow = now
	blockHeightCount++
	return getJwkKey(), nil, now, nil, blockHeightCount
}

func getJwkKey() jwk.Key {
	keyBytes, err := os.ReadFile("../_tests/arweavekeys/5SUBakh_R97MbHoX0_wNarVUw6DH0TziW5rG2K1vc6k.json")
	if err != nil {
		return nil
	}
	key, err := jwk.ParseKey(keyBytes)
	if err != nil {
		panic(err)
	}
	return key
}

func getJwkKey2() jwk.Key {
	keyBytes, err := os.ReadFile("../_tests/arweavekeys/axJNcs4-2yv5-yihgmwuDyuRql_06mhLa0PtrwP3PQo.json")
	if err != nil {
		return nil
	}
	key, err := jwk.ParseKey(keyBytes)
	if err != nil {
		panic(err)
	}
	return key
}
func getJwkKey3() jwk.Key {
	keyBytes, err := os.ReadFile("../_tests/arweavekeys/tst.json")
	if err != nil {
		return nil
	}
	key, err := jwk.ParseKey(keyBytes)
	if err != nil {
		panic(err)
	}
	return key
}

package routes

import (
	_ "encoding/base64"
	_ "github.com/everFinance/goar/types"
	"github.com/gin-gonic/gin"
	"github.com/warp-contracts/sequencer/config"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterSequence(t *testing.T) {
	config.Init("..")
	c := GetTestGinContext()
	c.Request.Method = http.MethodPost
	//wallet, err := goar.NewWalletFromPath("../_tests/arweavekeys/5SUBakh_R97MbHoX0_wNarVUw6DH0TziW5rG2K1vc6k.json", viper.GetString("arweave.url"))
	//assert.NoError(t, err)
	//transaction := &types.Transaction{}
	//assert.NoError(t, wallet.Signer.SignTx(transaction))
	//jsonTransaction, err := json.Marshal(transaction)
	//assert.NoError(t, err)

	//c.Request.Body = io.NopCloser(bytes.NewReader(jsonTransaction))

	RegisterSequencer(c)
}
func GetTestGinContext() *gin.Context {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	return ctx
}

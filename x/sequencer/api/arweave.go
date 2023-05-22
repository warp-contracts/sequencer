package api

import (
	"encoding/json"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gorilla/mux"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

type arweaveHandler struct {
	ctx client.Context
}

const ArweaveEndpoint = "/api/v1/arweave"

// The endpoint that accepts the Arweave transaction in the form of JSON, 
// wraps it with a Cosmos transaction and broadcasts it to the network.
func RegisterArweaveAPIRoute(clientCtx client.Context, router *mux.Router) {
	router.Handle(ArweaveEndpoint, arweaveHandler{ctx: clientCtx}).Methods("POST")
}

func (h arweaveHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// parse JSON
	var msg types.MsgArweave
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&msg)
	if err != nil {
		badRequest(w, err.Error())
		return
	}

	// wrap Arweave message with Cosmos transaction
	txBytes, err := createTxWithArweaveMsg(h.ctx, msg)
	if err != nil {
		badRequest(w, err.Error())
		return
	}

	// validate and broadcast transaction
	response, err := h.ctx.BroadcastTxSync(txBytes)
	if err != nil {
		badRequest(w, err.Error())
		return
	}
	if response.Code != 0 {
		badRequest(w, response.RawLog)
		return
	}

	w.Write([]byte(response.TxHash))
}

func badRequest(w http.ResponseWriter, errorMessage string) {
	http.Error(w, errorMessage, http.StatusBadRequest)
}

func createTxWithArweaveMsg(ctx client.Context, msg types.MsgArweave) ([]byte, error) {
	txBuilder := ctx.TxConfig.NewTxBuilder()
	txBuilder.SetMsgs(&msg)
	return ctx.TxConfig.TxEncoder()(txBuilder.GetTx())
}

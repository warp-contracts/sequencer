package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
)

type txBySenderNonceHandler struct {
	ctx      client.Context
	validate *validator.Validate
}

type TxBySenderRequest struct {
	Sender string `json:"sender" validate:"required,alphanum,startswith=warp,min=12,max=94"`
	Nonce  uint64 `json:"nonce" validate:"number,min=0"`
}

type TxBySenderResponse struct {
	Hash string `json:"tx_hash"`
}

// The endpoint that returns the transaction hash for a given sender and nonce.
// If such a transaction is not found, a response with a 404 status is returned.
func RegisterTxAPIRoute(clientCtx client.Context, router *mux.Router) {
	router.Handle("/api/v1/tx", txBySenderNonceHandler{ctx: clientCtx, validate: validator.New()}).Methods("POST")
}

func (h txBySenderNonceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request TxBySenderRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		BadRequestError(w, err, "request decoding error")
		return
	}

	err = h.validate.Struct(request)
	if err != nil {
		BadRequestError(w, err, "invalid request")
		return
	}

	events := []string{
		fmt.Sprintf("%s.%s='%s/%d'", sdk.EventTypeTx, sdk.AttributeKeyAccountSequence, request.Sender, request.Nonce),
	}
	txs, err := tx.QueryTxsByEvents(h.ctx, events, query.DefaultPage, query.DefaultLimit, "")
	if err != nil {
		InternalServerError(w, err, "query txs error")
		return
	}

	if len(txs.Txs) == 0 {
		NotFoundResponse(w, "transaction not found for the given sender and nonce")
		return
	}

	if len(txs.Txs) > 1 {
		InternalServerErrorString(w, "more than one transaction matches the criteria", "query txs error")
		return
	}

	OkResponse(w, TxBySenderResponse{txs.Txs[0].TxHash})
}

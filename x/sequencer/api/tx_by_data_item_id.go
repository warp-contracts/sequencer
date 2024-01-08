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

	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

type txByDataItemIdHandler struct {
	ctx      client.Context
	validate *validator.Validate
}

type TxByDataItemIdRequest struct {
	DataItemId string `json:"data_item_id" validate:"required,base64rawurl"`
}

type TxByDataItemIdResponse struct {
	Hash string `json:"tx_hash"`
}

// The endpoint that returns the transaction hash for a given data item id.
// If such a transaction is not found, a response with a 404 status is returned.
func RegisterTxByDataItemIdAPIRoute(clientCtx client.Context, router *mux.Router) {
	router.Handle("/api/v1/tx-data-item-id", &txByDataItemIdHandler{ctx: clientCtx, validate: validator.New()}).Methods("POST")
}

func (h *txByDataItemIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request TxByDataItemIdRequest
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

	eventQuery := fmt.Sprintf("%s.%s='%s'", sdk.EventTypeTx, types.AttributeKeyDataItemId, request.DataItemId)
	txs, err := tx.QueryTxsByEvents(h.ctx, query.DefaultPage, query.DefaultLimit, eventQuery, "")
	if err != nil {
		InternalServerError(w, err, "query txs error")
		return
	}

	if len(txs.Txs) == 0 {
		NotFoundResponse(w, "transaction not found for the given data item id")
		return
	}

	if len(txs.Txs) > 1 {
		InternalServerErrorString(w, "more than one transaction matches the criteria", "query txs error")
		return
	}

	OkResponse(w, TxBySenderResponse{txs.Txs[0].TxHash})
}

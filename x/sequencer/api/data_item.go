package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gorilla/mux"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

type dataItemHandler struct {
	ctx client.Context
}

type dataItemResponse struct {
	SequencerTxHash string `json:"sequencer_tx_hash"`
	DataItemId      string `json:"data_item_id"`
}

// The endpoint that accepts the DataItems as described in AND-104
// wraps it with a Cosmos transaction and broadcasts it to the network.
func RegisterDataItemAPIRoute(clientCtx client.Context, router *mux.Router) {
	router.Handle("/api/v1/dataitem", dataItemHandler{ctx: clientCtx}).Methods("POST")
}

func (h dataItemHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var msg types.MsgDataItem

	// Parse DataItem from request body
	err := msg.DataItem.UnmarshalFromReader(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to parse data item: %s", err.Error()), http.StatusBadRequest)
		return
	}

	// Wrap message with Cosmos transaction, validate and broadcast transaction
	response, err := types.BroadcastDataItem(h.ctx, msg)
	if err != nil {
		http.Error(w, "failed to broadcast transaction", http.StatusInternalServerError)
		return
	}
	if response.Code != 0 {
		http.Error(w, "failed to broadcast transaction: " + response.RawLog, http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(dataItemResponse{
		SequencerTxHash: response.TxHash,
		DataItemId:      msg.DataItem.Id.Base64(),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = fmt.Fprintf(w, "%s", jsonResponse)
	if err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
}

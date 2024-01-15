package api

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gorilla/mux"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

type dataItemHandler struct {
	ctx client.Context
}

// The endpoint that accepts the DataItems as described in AND-104
// wraps it with a Cosmos transaction and broadcasts it to the network.
func RegisterDataItemAPIRoute(clientCtx client.Context, router *mux.Router) {
	router.Handle("/api/v1/data-item", dataItemHandler{ctx: clientCtx}).Methods("POST")
}

func (h dataItemHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var msg types.MsgDataItem

	// Parse DataItem from request body
	// FIXME 
	// Consider whether this unmarshalling is necessary, 
	// since the `BroadcastDataItem` function is about to encode the transaction to bytes
	err := msg.DataItem.UnmarshalFromReader(r.Body)
	if err != nil {
		BadRequestError(w, err, "parse data item error")
		return
	}

	// Get broadcast mode from request header
	mode := r.Header.Get("X-Broadcast-Mode")
	switch mode {
	case flags.BroadcastSync, flags.BroadcastAsync:
		h.ctx.BroadcastMode = mode
	default:
		h.ctx.BroadcastMode = flags.BroadcastSync
	}

	// Wrap message with Cosmos transaction, validate and broadcast transaction
	response, err := types.BroadcastDataItem(h.ctx, &msg)
	if err != nil {
		InternalServerError(w, err, "broadcast transaction error")
		return
	}
	if response.Code != 0 {
		ErrorWithStatus(w, response, "broadcast response with non-zero code", responseCodeToStatus(response.Code))
		return
	}

	OkResponse(w, msg.GetInfo())
}

func responseCodeToStatus(responseCode uint32) int {
	switch responseCode {
	case sdkerrors.ErrWrongSequence.ABCICode():
		return http.StatusConflict
	case sdkerrors.ErrMempoolIsFull.ABCICode():
		return http.StatusServiceUnavailable
	default:
		return http.StatusBadRequest
	}
}

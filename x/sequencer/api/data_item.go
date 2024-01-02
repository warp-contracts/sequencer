package api

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gorilla/mux"
	limitermodulekeeper "github.com/warp-contracts/sequencer/x/limiter/keeper"
	"github.com/warp-contracts/sequencer/x/sequencer"
	"github.com/warp-contracts/sequencer/x/sequencer/types"
)

type dataItemHandler struct {
	ctx           client.Context
	limiterKeeper *limitermodulekeeper.Keeper
}

// The endpoint that accepts the DataItems as described in AND-104
// wraps it with a Cosmos transaction and broadcasts it to the network.
func RegisterDataItemAPIRoute(clientCtx client.Context, router *mux.Router, limiterKeeper *limitermodulekeeper.Keeper) {
	router.Handle("/api/v1/data-item", dataItemHandler{ctx: clientCtx, limiterKeeper: limiterKeeper}).Methods("POST")
}

func (h dataItemHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var msg types.MsgDataItem

	// Parse DataItem from request body
	err := msg.DataItem.UnmarshalFromReader(r.Body)
	if err != nil {
		BadRequestError(w, err, "parse data item error")
		return
	}

	// Check if there isn't too many requests for this contract
	contractId, err := msg.GetContractFromTags()
	if err != nil {
		BadRequestError(w, err, "parse contract id error")
		return
	}

	if h.limiterKeeper.GetCount(sequencer.LIMITER_CONTRACT_ID, []byte(contractId)) > sequencer.LIMITER_CONTRACT_ID_MAX_REQUESTS {
		TooManyRequestsError(w, err, "parse contract id error")
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

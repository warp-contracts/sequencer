package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/warp-contracts/sequencer/x/sequencer/types"

	"github.com/warp-contracts/syncer/src/utils/bundlr"
)

type nonceHandler struct {
	ctx      client.Context
	validate *validator.Validate
}

type nonceRequest struct {
	SignatureType int    `json:"signature_type" validate:"required,oneof=1 3"`
	Owner         string `json:"owner" validate:"required,base64rawurl"`
}

type nonceResponse struct {
	Address string `json:"address"`
	Nonce   uint64 `json:"nonce"`
}

// The endpoint that returns the account address and nonce for the given fields of the DataItem:
// owner (in Base64URL format) and signature type.
func RegisterNonceAPIRoute(clientCtx client.Context, router *mux.Router) {
	router.Handle("/api/v1/nonce", nonceHandler{ctx: clientCtx, validate: validator.New()}).Methods("POST")
}

func (h nonceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request nonceRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.validate.Struct(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	publicKey, err := getPublicKey(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := getAddressWithNonce(h.ctx, publicKey)
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", jsonResponse)
}

func getPublicKey(request nonceRequest) (key cryptotypes.PubKey, err error) {
	ownerBytes, err := base64.RawURLEncoding.DecodeString(request.Owner)
	if err != nil {
		return
	}

	signatureType := bundlr.SignatureType(request.SignatureType)
	return types.GetPublicKey(signatureType, ownerBytes)
}

func getAddressWithNonce(ctx client.Context, key cryptotypes.PubKey) nonceResponse {
	address := sdk.AccAddress(key.Address())
	response := nonceResponse{Address: address.String()}

	acc, err := ctx.AccountRetriever.GetAccount(ctx, address)
	if acc == nil || err != nil {
		// account does not exist
		response.Nonce = 0
	} else {
		response.Nonce = acc.GetSequence()
	}

	return response
}

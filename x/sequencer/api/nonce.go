package api

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/warp-contracts/sequencer/x/sequencer/types"

	"github.com/warp-contracts/syncer/src/utils/bundlr"
)

type nonceHandler struct {
	ctx client.Context
}

type nonceResponse struct {
	Address string `json:"address"`
	Nonce   uint64 `json:"nonce"`
}

// The endpoint that returns the account address and nonce for the given fields of the DataItem: 
// owner (in Base64URL format) and signature type.
func RegisterNonceAPIRoute(clientCtx client.Context, router *mux.Router) {
	router.Handle("/api/v1/nonce", nonceHandler{ctx: clientCtx})
}

func (h nonceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	publicKey, err := getPublicKeyFromParameters(r)
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

func getPublicKeyFromParameters(r *http.Request) (key cryptotypes.PubKey, err error) {
	params := r.URL.Query()

	owner := params.Get("owner")
	if len(owner) == 0 {
		err = errors.New("no owner parameter")
		return
	}
	ownerBytes, err := base64.URLEncoding.DecodeString(owner)
	if err != nil {
		return
	}

	signatureTypeStr := params.Get("signature_type")
	if len(signatureTypeStr) == 0 {
		err = errors.New("no signature_type parameter")
		return
	}
	signatureTypeInt, err := strconv.Atoi(signatureTypeStr)
	if err != nil {
		err = errors.New("signature_type should be numeric")
		return
	}
	signatureType := bundlr.SignatureType(signatureTypeInt)
	if signatureType != bundlr.SignatureTypeArweave && signatureType != bundlr.SignatureTypeEthereum {
		err = fmt.Errorf("invalid signature_type value, should be %d or %d",
			bundlr.SignatureTypeArweave, bundlr.SignatureTypeEthereum)
		return
	}

	key, err = types.GetPublicKey(signatureType, ownerBytes)
	if err != nil {
		return
	}

	return
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

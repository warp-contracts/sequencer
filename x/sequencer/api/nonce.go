package api

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/warp-contracts/sequencer/x/sequencer/types"

	"github.com/warp-contracts/syncer/src/utils/bundlr"
)

type AccountProvider interface {
	GetAccount(address sdk.AccAddress) (sdk.AccountI, error)
}

type nonceHandler struct {
	accountProvider AccountProvider
	validate        *validator.Validate
}

type NonceRequest struct {
	SignatureType int    `json:"signature_type" validate:"required,oneof=1 3"`
	Owner         string `json:"owner" validate:"required,base64rawurl,min=87,max=683"`
}

type NonceResponse struct {
	Address string `json:"address"`
	Nonce   uint64 `json:"nonce"`
}

// The endpoint that returns the account address and nonce for the given fields of the DataItem:
// owner (in Base64URL format) and signature type.
func RegisterNonceAPIRoute(accountProvider AccountProvider, router *mux.Router) {
	router.Handle("/api/v1/nonce", nonceHandler{accountProvider: accountProvider, validate: validator.New()}).Methods("POST")
}

func (h nonceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request NonceRequest
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

	publicKey, err := getPublicKey(request)
	if err != nil {
		BadRequestError(w, err, "public key problem")
		return
	}

	address := sdk.AccAddress(publicKey.Address())
	account, err := h.accountProvider.GetAccount(address)
	if err != nil {
		InternalServerError(w, err, "query nonce error")
		return
	}

	response := NonceResponse{Address: address.String()}
	if account != nil {
		response.Nonce = account.GetSequence()
	}
	OkResponse(w, response)
}

func getPublicKey(request NonceRequest) (key cryptotypes.PubKey, err error) {
	ownerBytes, err := base64.RawURLEncoding.DecodeString(request.Owner)
	if err != nil {
		return
	}

	signatureType := bundlr.SignatureType(request.SignatureType)
	return types.GetPublicKey(signatureType, ownerBytes)
}

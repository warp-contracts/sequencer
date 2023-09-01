package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/warp-contracts/sequencer/crypto/keys/arweave"
)

const TypeMsgArweaveTransaction = "arweave_transaction"

var _ sdk.Msg = &MsgArweaveTransaction{}

func (msg *MsgArweaveTransaction) Route() string {
	return RouterKey
}

func (msg *MsgArweaveTransaction) Type() string {
	return TypeMsgArweaveTransaction
}

func (msg *MsgArweaveTransaction) GetCreator() sdk.AccAddress {
	return sdk.AccAddress(arweave.FromOwner(msg.Transaction.Owner).Address())
}

func (msg *MsgArweaveTransaction) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.GetCreator()}
}

func (msg *MsgArweaveTransaction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgArweaveTransaction) ValidateBasic() error {
	return msg.Transaction.Verify()
}

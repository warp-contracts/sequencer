package ante

import (
	storetypes "cosmossdk.io/store/types"
	txsigning "cosmossdk.io/x/tx/signing"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

type HandlerOptions struct {
	AccountKeeper          authkeeper.AccountKeeper
	BankKeeper             authtypes.BankKeeper
	ExtensionOptionChecker ante.ExtensionOptionChecker
	FeegrantKeeper         ante.FeegrantKeeper
	SignModeHandler        *txsigning.HandlerMap
	SigGasConsumer         func(meter storetypes.GasMeter, sig signing.SignatureV2, params authtypes.Params) error
	TxFeeChecker           ante.TxFeeChecker
}

// TODO
func NewAnteHandler(AccountKeeper authkeeper.AccountKeeper,
	BankKeeper authtypes.BankKeeper,
	FeegrantKeeper ante.FeegrantKeeper,
	txConfig client.TxConfig) sdk.AnteHandler {
	anteDecorators := []sdk.AnteDecorator{
		ante.NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
		NewSetInfiniteGasMeterDecorator(),
		NewDataItemTxDecorator(&AccountKeeper),
		NewArweaveBlockTxDecorator(),
		ante.NewExtensionOptionsDecorator(nil),
		ante.NewValidateBasicDecorator(),
		ante.NewTxTimeoutHeightDecorator(),
		ante.NewValidateMemoDecorator(AccountKeeper),
		ante.NewConsumeGasForTxSizeDecorator(AccountKeeper),
		ante.NewDeductFeeDecorator(AccountKeeper, BankKeeper, FeegrantKeeper, nil),
		ante.NewSetPubKeyDecorator(AccountKeeper), // SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewValidateSigCountDecorator(AccountKeeper),
		ante.NewSigGasConsumeDecorator(AccountKeeper, SigVerificationGasConsumer),
		ante.NewSigVerificationDecorator(AccountKeeper, txConfig.SignModeHandler()),
		ante.NewIncrementSequenceDecorator(AccountKeeper),
	}

	return sdk.ChainAnteDecorators(anteDecorators...)
}

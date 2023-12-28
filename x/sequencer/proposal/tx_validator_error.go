package proposal

type InvalidTxErrorType int

const (
	INVALID_ARWEAVE InvalidTxErrorType = iota
	INVALID_DATA_ITEM
)

type InvalidTxError struct {
	err       error
	errorType InvalidTxErrorType
}

func InvalidArweaveError(err error) *InvalidTxError {
	return &InvalidTxError{err, INVALID_ARWEAVE}
}

func InvalidDataItemError(err error) *InvalidTxError {
	return &InvalidTxError{err, INVALID_DATA_ITEM}
}

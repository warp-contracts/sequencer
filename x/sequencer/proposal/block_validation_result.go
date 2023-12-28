package proposal

import "sync"

// The validation result is protected by a mutex, ensuring error is sent exactly once to the output channel.
// The sending occurs when the first error appears, or nil is sent at the end.
type validationResult struct {
	// has the result been sent to the channel?
	sent bool
	// the channel to which the validation result is sent
	output chan *InvalidTxError
	mtx    sync.RWMutex
}

func newValidationResult(output chan *InvalidTxError) *validationResult {
	return &validationResult{
		sent:   false,
		output: output,
	}
}

func (vr *validationResult) isNotSent() bool {
	vr.mtx.RLock()
	defer vr.mtx.RUnlock()

	return !vr.sent
}

func (vr *validationResult) sendFirstError(err *InvalidTxError) {
	if err != nil {
		vr.mtx.Lock()
		defer vr.mtx.Unlock()

		if !vr.sent {
			vr.output <- err
			vr.sent = true
		}
	}
}

func (vr *validationResult) sendIfNoError() {
	vr.mtx.Lock()
	defer vr.mtx.Unlock()

	if !vr.sent {
		vr.output <- nil
		vr.sent = true
	}
}

package proposal

import "sync"

// The validation result is protected by a mutex, ensuring error is sent exactly once to the output channel.
// The sending occurs when the first error appears, or nil is sent at the end.
type validationResult struct {
	valid  bool
	mtx    sync.RWMutex
	output chan error
}

func newValidationResult(output chan error) *validationResult {
	return &validationResult{
		valid:  true,
		output: output,
	}
}

func (vr *validationResult) isValid() bool {
	vr.mtx.RLock()
	defer vr.mtx.RUnlock()

	return vr.valid
}

func (vr *validationResult) sendFirstError(err error) {
	if err != nil {
		vr.mtx.Lock()
		defer vr.mtx.Unlock()

		if vr.valid {
			vr.output <- err
			vr.valid = false	
		}
	}
}

func (vr *validationResult) sendIfValid() {
	vr.mtx.RLock()
	defer vr.mtx.RUnlock()

	if vr.valid {
		vr.output <- nil
	}
}

package proposal

import "sync"

// The validation result is protected by a mutex, ensuring it is sent exactly once to the output channel. 
// Sending occurs upon the first incorrect transaction or at the very end if everything is fine.
type validationResult struct {
	result bool
	mtx    sync.RWMutex
	output chan bool
}

func newValidationResult(output chan bool) *validationResult {
	return &validationResult{
		result: true,
		output: output,
	}
}

func (vr *validationResult) isValid() bool {
	vr.mtx.RLock()
	defer vr.mtx.RUnlock()

	return vr.result
}

func (vr *validationResult) sendIfValid() {
	vr.mtx.RLock()
	defer vr.mtx.RUnlock()

	if vr.result {
		vr.output <- true
	}
}

func (vr *validationResult) sendFirstInvalid(newResult bool) {
	if !newResult && vr.isValid() {
		vr.mtx.Lock()
		defer vr.mtx.Unlock()

		vr.output <- false
		vr.result = false
	}
}

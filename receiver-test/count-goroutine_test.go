package receivertest

import "testing"

func TestCountGoroutine(t *testing.T) {
	//
	if err := CountGoroutine(); err != nil {
		t.Errorf("CountGoroutine() error = %v", err)
	}
}

package receivertest

import "testing"

func TestCartX(t *testing.T) {
	if err := CartX(); err != nil {
		t.Errorf("CartX() error = %v", err)
	}
}

// file: pkg/snettool/whoislookup_test.go
package snettool

import (
	"testing"
)

func TestWhoisLookup(t *testing.T) {
	result, err := WhoisLookup("google.com")
	if err != nil {
		t.Errorf("WhoisLookup failed: %v", err)
	}
	if len(result) == 0 {
		t.Errorf("Expected non-empty result, got %v", result)
	}
}

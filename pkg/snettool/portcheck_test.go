// file: pkg/snettool/portcheck_test.go
package snettool

import (
    "testing"
)

func TestCheckPort(t *testing.T) {
    result, err := CheckPort("google.com", "80")
    if err != nil {
        t.Errorf("CheckPort failed: %v", err)
    }
    if !result {
        t.Errorf("Expected true, got %v", result)
    }
}

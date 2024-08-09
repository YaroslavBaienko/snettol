// file: pkg/snettool/ping_test.go
package snettool

import (
    "testing"
)

func TestPing(t *testing.T) {
    result, err := Ping("google.com")
    if err != nil {
        t.Errorf("Ping failed: %v", err)
    }
    if !result {
        t.Errorf("Expected true, got %v", result)
    }
}

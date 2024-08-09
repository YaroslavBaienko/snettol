// file: pkg/snettool/ping.go
package snettool

import (
    "os/exec"
    "strings"
)

// Ping sends a ping to the specified host and returns true if the host is reachable.
func Ping(host string) (bool, error) {
    cmd := exec.Command("ping", "-c", "4", host)
    output, err := cmd.CombinedOutput()
    if err != nil {
        return false, err
    }
    return strings.Contains(string(output), "0% packet loss"), nil
}

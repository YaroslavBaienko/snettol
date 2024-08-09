// file: pkg/snettool/portcheck.go
package snettool

import (
    "net"
    "time"
)

// CheckPort tries to connect to the specified host and port and returns true if the port is open.
func CheckPort(host string, port string) (bool, error) {
    timeout := 3 * time.Second
    conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
    if err != nil {
        return false, err
    }
    defer conn.Close()
    return true, nil
}

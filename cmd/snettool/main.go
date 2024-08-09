// file: cmd/snettool/main.go
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/YaroslavBaienko/snettool/pkg/snettool"
)

type Command struct {
	Action string `json:"action"`
	Host   string `json:"host"`
	Port   string `json:"port,omitempty"`
}

func main() {
	var cmd Command
	decoder := json.NewDecoder(os.Stdin)
	err := decoder.Decode(&cmd)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	switch cmd.Action {
	case "ping":
		result, err := snettool.Ping(cmd.Host)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Ping result:", result)
		}
	case "checkport":
		result, err := snettool.CheckPort(cmd.Host, cmd.Port)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Port check result:", result)
		}
	case "dnslookup":
		result, err := snettool.DNSLookup(cmd.Host)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("DNS Lookup result:", result)
		}
	case "whoislookup":
		result, err := snettool.WhoisLookup(cmd.Host)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("WHOIS Lookup result:", result)
		}
	default:
		fmt.Println("Unknown action:", cmd.Action)
	}
}

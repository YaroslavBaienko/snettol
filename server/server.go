// file: server/server.go
package main

import (
	"encoding/json"
	"github.com/YaroslavBaienko/snettool/pkg/snettool"
	"net/http"
)

type Command struct {
	Action string `json:"action"`
	Host   string `json:"host"`
	Port   string `json:"port,omitempty"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	var cmd Command
	err := json.NewDecoder(r.Body).Decode(&cmd)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	switch cmd.Action {
	case "ping":
		result, err := snettool.Ping(cmd.Host)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			json.NewEncoder(w).Encode(result)
		}
	case "checkport":
		result, err := snettool.CheckPort(cmd.Host, cmd.Port)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			json.NewEncoder(w).Encode(result)
		}
	case "dnslookup":
		result, err := snettool.DNSLookup(cmd.Host)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			json.NewEncoder(w).Encode(result)
		}
	case "whoislookup":
		result, err := snettool.WhoisLookup(cmd.Host)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			json.NewEncoder(w).Encode(result)
		}
	default:
		http.Error(w, "Unknown action", http.StatusBadRequest)
	}
}

func main() {
	http.HandleFunc("/api", handler)
	http.ListenAndServe(":8080", nil)
}

// Package api contains the HTTP API used by the delivery lab.
package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Greeting is deliberately small so failures are easy to investigate.
func Greeting(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		name = "enemy"
	}
	return fmt.Sprintf("Hello, %s!", name)
}

// NewHandler constructs an isolated handler without opening a network port.
func NewHandler(version string) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, map[string]string{"status": "ok"})
	})
	mux.HandleFunc("GET /version", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, map[string]string{"version": version})
	})
	mux.HandleFunc("GET /greet", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, map[string]string{"message": Greeting(r.URL.Query().Get("name"))})
	})
	return mux
}

func writeJSON(w http.ResponseWriter, body map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	// These string maps are always JSON-encodable. A disconnected client can
	// still cause a write error; there is no response left to send in that case.
	_ = json.NewEncoder(w).Encode(body)
}

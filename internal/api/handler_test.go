package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGreeting(t *testing.T) {
	cases := []struct{ name, input, want string }{
		{"named", "Micah", "Hello, Micah!"},
		{"empty", "", "Hello, friend!"},
		{"whitespace", "   ", "Hello, friend!"},
		{"trimmed", "  Micah  ", "Hello, Micah!"},
		{"unicode", "世界", "Hello, 世界!"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Greeting(tc.input); got != tc.want {
				t.Fatalf("Greeting(%q) = %q; want %q", tc.input, got, tc.want)
			}
		})
	}
}

func TestEndpoints(t *testing.T) {
	cases := []struct{ path, key, want string }{
		{"/health", "status", "ok"},
		{"/version", "version", "test-version"},
		{"/greet?name=Micah", "message", "Hello, Micah!"},
		{"/greet", "message", "Hello, friend!"},
	}
	for _, tc := range cases {
		t.Run(tc.path, func(t *testing.T) {
			rec := httptest.NewRecorder()
			NewHandler("test-version").ServeHTTP(rec, httptest.NewRequest(http.MethodGet, tc.path, nil))
			if rec.Code != http.StatusOK {
				t.Fatalf("status = %d; want 200", rec.Code)
			}
			if got := rec.Header().Get("Content-Type"); got != "application/json" {
				t.Fatalf("Content-Type = %q; want application/json", got)
			}
			var body map[string]string
			if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
				t.Fatal(err)
			}
			if body[tc.key] != tc.want {
				t.Fatalf("body = %v; want %s=%q", body, tc.key, tc.want)
			}
		})
	}
}

func TestRoutingErrors(t *testing.T) {
	cases := []struct {
		method, path string
		want         int
	}{
		{http.MethodGet, "/missing", http.StatusNotFound},
		{http.MethodPost, "/health", http.StatusMethodNotAllowed},
		{http.MethodPost, "/greet", http.StatusMethodNotAllowed},
	}
	for _, tc := range cases {
		t.Run(tc.method+tc.path, func(t *testing.T) {
			rec := httptest.NewRecorder()
			NewHandler("dev").ServeHTTP(rec, httptest.NewRequest(tc.method, tc.path, nil))
			if rec.Code != tc.want {
				t.Fatalf("status = %d; want %d", rec.Code, tc.want)
			}
		})
	}
}

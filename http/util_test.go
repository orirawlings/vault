package http

import (
	"net/http"
	"testing"
)

func testSeverAuth(t *testing.T, addr string, token string) {
	if _, err := http.Get(addr + "/_test/auth?token=" + token); err != nil {
		t.Fatalf("error authenticating: %s", err)
	}
}

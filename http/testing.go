package http

import (
	"fmt"
	"net"
	"net/http"

	"github.com/hashicorp/vault/vault"
)

func TestListener(fail func(format string, args ...interface{})) (net.Listener, string) {
	if fail == nil {
		fail = func(format string, args ...interface{}) {
			panic(fmt.Sprintf(format, args...))
		}
	}

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		fail("err: %s", err)
	}
	addr := "http://" + ln.Addr().String()
	return ln, addr
}

func TestServerWithListener(ln net.Listener, addr string, core *vault.Core) {
	// Create a muxer to handle our requests so that we can authenticate
	// for tests.
	mux := http.NewServeMux()
	mux.Handle("/_test/auth", http.HandlerFunc(testHandleAuth))
	mux.Handle("/", Handler(core))

	server := &http.Server{
		Addr:    ln.Addr().String(),
		Handler: mux,
	}
	go server.Serve(ln)
}

func TestServer(fail func(format string, args ...interface{}), core *vault.Core) (net.Listener, string) {
	ln, addr := TestListener(fail)
	TestServerWithListener(ln, addr, core)
	return ln, addr
}

func testHandleAuth(w http.ResponseWriter, req *http.Request) {
	respondOk(w, nil)
}

package redirecttest

import (
	"fmt"
	"net/http"
	"testing"
)

func RedirectTo(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Hello Redirect")
}

func RedirectFrom(writer http.ResponseWriter, request *http.Request) {
	// logic
	http.Redirect(writer, request, "/redirect-to", http.StatusMovedPermanently)
}

func RedirectOut(writer http.ResponseWriter, request *http.Request) {
	http.Redirect(writer, request, "https://www.programmerzamannow.com", http.StatusMovedPermanently)
}

func TestRedirect(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect-from", RedirectFrom)
	mux.HandleFunc("/redirect-to", RedirectTo)
	mux.HandleFunc("/redirect-out", RedirectOut)

	http.ListenAndServe(":8080", mux)
}

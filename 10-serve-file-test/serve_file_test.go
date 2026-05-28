package servefiletest

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		http.ServeFile(writer, request, "./resources/ok.html")
	} else {
		http.ServeFile(writer, request, "./resources/notfound.html")
	}
}

func TestServeFileServer(t *testing.T) {
	server := http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// jika menggunakan embed maka ambil string html nya lalu buat file baru dengan nama ok.html dan notfound.html lalu isi dengan string html yang sudah diambil tadi

//go:embed resources/ok.html
var okHTML string

//go:embed resources/notfound.html
var notFoundHTML string

func ServeFileEmbed(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		// writer.Write([]byte(okHTML)) // ini lebih boros karena harus mengubah string ke byte
		// pakai writer.Write hanya jika kita sudah memiliki byte, jika kita sudah memiliki string maka lebih efisien menggunakan fmt.Fprint karena langsung menulis string tanpa harus mengubahnya ke byte terlebih dahulu
		// pzn
		// jadi intinya jika kita sudah memiliki string maka lebih efisien menggunakan fmt.Fprint
		// karena langsung menulis string tanpa harus mengubahnya ke byte terlebih dahulu,
		// sedangkan jika kita sudah memiliki byte maka lebih efisien menggunakan writer.Write
		// karena langsung menulis byte tanpa harus mengubahnya ke string terlebih dahulu
		fmt.Fprint(writer, okHTML)
	} else {
		// writer.Write([]byte(notFoundHTML))
		// pzn
		fmt.Fprint(writer, notFoundHTML)
	}
}

func TestServeFileEmbed(t *testing.T) {
	server := http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

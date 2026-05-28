package fileservertest

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {
	directory := http.Dir("./resources")
	fileServer := http.FileServer(directory)

	mux := http.NewServeMux()
	// strip prefix digunakan untuk menghapus bagian "/static/" dari URL sebelum mencari file di direktori
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		t.Fatal(err)
	}
}

//go:embed resources
var resources embed.FS

func TestFileServerGolangEmbed(t *testing.T) {
	// di golang embed nanti jadinya /static/resources/isinya
	directory, _ := fs.Sub(resources, "resources") // nah ini untuk mengambil sub direktori "resources" dari embed.FS
	fileServer := http.FileServer(http.FS(directory))

	mux := http.NewServeMux()
	// strip prefix digunakan untuk menghapus bagian "/static/" dari URL sebelum mencari file di direktori
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		t.Fatal(err)
	}
}

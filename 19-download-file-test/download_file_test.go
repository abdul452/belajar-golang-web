package downloadfiletest

import (
	"fmt"
	"net/http"
	"testing"
)

func DownloadFile(writer http.ResponseWriter, request *http.Request) {
	file := request.URL.Query().Get("file")

	if file == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(writer, "Bad Request")
		return
	}

	// ini buat auto download file nya
	writer.Header().Add("Content-Disposition", "attachment; filename=\""+file+"\"")

	http.ServeFile(writer, request, "./resources/"+file)
}

func TestDownloadFile(t *testing.T) {
	server := http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(DownloadFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

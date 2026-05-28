package uploadfiletest

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

//go:embed templates/*.gohtml
var templates embed.FS

var myTemplates = template.Must(template.ParseFS(templates, "templates/*.gohtml"))

func UploadForm(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "upload_form.gohtml", nil)
}

// running server
func TestUploadForm(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func Upload(writer http.ResponseWriter, request *http.Request) {
	// request.ParseMultipartForm(100 >> 20) // ini maks file nya 100MB
	// ambil file nya
	file, fileHeader, err := request.FormFile("file") // defaultnya 32MB
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// create destination nya
	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}
	// masukan ke path destination
	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}
	name := request.FormValue("name")
	fmt.Println("name : ", name)
	fmt.Println("File : ", fileHeader.Filename)

	data := map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	}
	myTemplates.ExecuteTemplate(writer, "upload_success.gohtml", data)
}

//go:embed resources/pzn.png
var uploadFileTest []byte

func TestUploadFile(t *testing.T) {
	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)
	writer.WriteField("name", "Abdul Salim")
	file, _ := writer.CreateFormFile("file", "test.png")
	file.Write(uploadFileTest)
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/upload", body)
	request.Header.Set("Content-type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	Upload(recorder, request)

	response := recorder.Result()
	bodyRespon, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(bodyRespon))
}

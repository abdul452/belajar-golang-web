package main

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SimpleHTML(writer http.ResponseWriter, request *http.Request) {
	templateText := `<html><body>{{.}}</body></html>`
	// cara manual
	// t, err := template.New("SIMPLE").Parse(templateText)
	// if err != nil {
	// 	panic(err)
	// }

	t := template.Must(template.New("SIMPLE").Parse(templateText))
	err := t.ExecuteTemplate(writer, "SIMPLE", "Hello HTML Template")
	if err != nil {
		panic(err)
	}

}

func TestSimpleHTML(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTML(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func SimpleHTMLFile(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/simpple.gohtml"))
	err := t.ExecuteTemplate(writer, "simpple.gohtml", "Hello HTML Template")
	if err != nil {
		panic(err)
	}
}

func TestSimpleHTMLFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLFile(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateDirektory(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))
	err := t.ExecuteTemplate(writer, "simpple.gohtml", "Hello HTML Template")
	if err != nil {
		panic(err)
	}
}

func TestTemplateDirektory(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDirektory(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

//go:embed templates/*.gohtml
var templateFS embed.FS

func TemplateEmbed(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFS(templateFS, "templates/*.gohtml"))
	err := t.ExecuteTemplate(writer, "simpple.gohtml", "Hello HTML Template")
	if err != nil {
		panic(err)
	}
}

func TestTemplateEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateEmbed(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

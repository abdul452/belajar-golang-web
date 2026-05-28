package templatecachingtest

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

//go:embed templates/*.gohtml
var templates embed.FS

var myTemplates = template.Must(template.ParseFS(templates, "templates/*.gohtml"))

func TemplateCaching(writer http.ResponseWriter, request *http.Request) {
	// Use the pre-parsed templates
	err := myTemplates.ExecuteTemplate(writer, "simple.gohtml", "Hello Template Caching")
	if err != nil {
		panic(err)
	}
}

func TestTemplateCaching(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateCaching(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

// Xss test
func TemplateAutoEscaping(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Hello Template Caching",
		"Body":  "<p>Ini adalah body<script>alert('Anda di hack')</script></p>",
	})
}

func TestTemplateAutoEscaping(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscaping(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TestTemplateAutoEscapingCheck(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscaping(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	// 🎯 PANAH BUKTI: Kita cek apakah ada teks "&lt;" di dalam hasil datanya
	hasilTeks := string(body)
	fmt.Println("Apakah mengandung '&lt;'? :", strings.Contains(hasilTeks, "&lt;"))
	fmt.Println("Apakah mengandung '<p>'?  :", strings.Contains(hasilTeks, "<p>"))
}

// mematikan auto escaping, contohnya harus dari database, karena kita sudah yakin data tersebut aman untuk di render
func TemplateAutoEscapingOff(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Hello Template Caching",
		"Body":  template.HTML("<p>Ini adalah body</p>"),
	})
}

func TestTemplateAutoEscapingOff(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscapingOff(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

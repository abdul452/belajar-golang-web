package templatefunctiontest

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + ", My Name is " + myPage.Name
}

func TemplateFunction(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{.SayHello "John"}}`))

	data := MyPage{
		Name: "Abdul Salim",
	}

	err := t.ExecuteTemplate(writer, "FUNCTION", data)
	if err != nil {
		panic(err)
	}
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TemplateFunctionGlobal(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{len .Name}}`))

	data := MyPage{
		Name: "Abdul Salim",
	}

	err := t.ExecuteTemplate(writer, "FUNCTION", data)
	if err != nil {
		panic(err)
	}
}

func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobal(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TemplateFunctionGlobalMap(writer http.ResponseWriter, request *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse(`{{upper .Name}}`))

	data := MyPage{
		Name: "Abdul Salim",
	}
	err := t.ExecuteTemplate(writer, "FUNCTION", data)
	if err != nil {
		panic(err)
	}
}

func TestTemplateFunctionGlobalMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobalMap(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TemplateFunctionGlobalPipeline(writer http.ResponseWriter, request *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"sayHello": func(value string) string {
			return "Hello " + value
		},
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse(`{{sayHello .Name | upper}}`))

	data := MyPage{
		Name: "Abdul Salim",
	}
	err := t.ExecuteTemplate(writer, "FUNCTION", data)
	if err != nil {
		panic(err)
	}
}

func TestTemplateFunctionGlobalPipeline(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobalPipeline(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TemplateFunctionGlobalPipelineFiles(writer http.ResponseWriter, request *http.Request) {
	// 1. Inisialisasi template kosongan dan beri nama yang sama dengan nama file target
	t := template.New("name.gohtml")

	// 2. DAFTARKAN FUNGSI TERLEBIH DAHULU (Wajib sebelum ParseFiles)
	t = t.Funcs(map[string]interface{}{
		"sayHello": func(value string) string {
			return "Hello " + value
		},
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	// 3. BARU BACA FILE NYA SEKARANG
	// Gunakan template.Must untuk membungkus parsing file fisik
	t = template.Must(t.ParseFiles("./templates/name.gohtml"))

	data := MyPage{
		Name: "Abdul Salim",
	}

	// 4. Eksekusi template (panggil nama filenya sebagai target)
	err := t.ExecuteTemplate(writer, "name.gohtml", data)
	if err != nil {
		panic(err)
	}
}

func TestTemplateFunctionGlobalPipelineFiles(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobalPipelineFiles(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

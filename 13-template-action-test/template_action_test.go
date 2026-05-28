package templateactiontest

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Data struct {
	Title string
	Name  string
}

func TemplateActionIf(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))

	data := Data{
		Title: "Template Data Struct",
		// Name:  "Abdul Salim",
	}

	err := t.ExecuteTemplate(writer, "name.gohtml", data)
	if err != nil {
		panic(err)
	}
}

func TestTemplateActionIf(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionIf(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TemplateActionOperator(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))

	data := map[string]interface{}{
		"Title":      "TemplateActionOperator",
		"FinalValue": 65,
	}

	err := t.ExecuteTemplate(writer, "comparator.gohtml", data)
	if err != nil {
		panic(err)
	}
}

func TestTemplateActionOperator(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionOperator(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TemplateActionRange(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))

	data := map[string]interface{}{
		"Title":   "TemplateActionRange",
		"Hobbies": []string{"Coding", "Gaming", "Traveling"},
	}

	err := t.ExecuteTemplate(writer, "range.gohtml", data)
	if err != nil {
		panic(err)
	}
}

func TestTemplateActionRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionRange(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TemplateActionWith(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/address.gohtml"))

	data := map[string]interface{}{
		"Title": "TemplateActionWith",
		"Name":  "Abdul Salim",
		"Address": map[string]interface{}{
			"Street": "Jl. Merdeka No. 123",
			"City":   "Jakarta",
		},
	}

	err := t.ExecuteTemplate(writer, "address.gohtml", data)
	if err != nil {
		panic(err)
	}
}

func TestTemplateActionWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionWith(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

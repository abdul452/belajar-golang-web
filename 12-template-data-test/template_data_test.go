package templatedatatest

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateDataMap(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	data := map[string]interface{}{
		"Title": "Template Data Map",
		"Name":  "Abdul Salim",
		// "Address": map[string]interface{}{
		// 	"Street": "Jl. Merdeka No. 123",
		// },
	}
	err := t.ExecuteTemplate(writer, "name.gohtml", data)
	if err != nil {
		panic(err)
	}
}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

type Address struct {
	Street string
}

type Data struct {
	Title   string
	Name    string
	Address Address
}

func TemplateDataStruct(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))

	data := Data{
		Title: "Template Data Struct",
		Name:  "Abdul Salim",
		Address: Address{
			Street: "Jl. Merdeka No. 123",
		},
	}

	err := t.ExecuteTemplate(writer, "name.gohtml", data)
	if err != nil {
		panic(err)
	}
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

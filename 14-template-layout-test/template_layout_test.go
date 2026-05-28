package templatelayouttest

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateLayout(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles(
		"./templates/layout.gohtml",
		"./templates/header.gohtml",
		"./templates/footer.gohtml",
	))

	data := map[string]interface{}{
		"Title": "Template Layout",
		"Name":  "Abdul Salim",
	}

	err := t.ExecuteTemplate(writer, "layout", data)
	if err != nil {
		panic(err)
	}
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

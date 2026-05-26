package httptest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello World")
}

func AboutHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "About Us")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

}

// dari gemini test api
type CustomerRequest struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func CreateCustomerHandler(writer http.ResponseWriter, request *http.Request) {
	// Membaca Request Body berbentuk JSON
	var payload CustomerRequest
	decoder := json.NewDecoder(request.Body)
	_ = decoder.Decode(&payload)

	// Set header agar client tahu ini balasan JSON
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated) // Status 201 Created

	// Kirim balik data sebagai tanda sukses
	json.NewEncoder(writer).Encode(payload)
}

func TestCreateCustomerJSON(t *testing.T) {
	// 1. Siapkan data JSON mentah yang biasa kamu ketik di body Postman
	requestBody := `{"id": "CUST-001", "name": "Abdul"}`

	// 2. Buat request tiruan (Method POST, bawa body JSON)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/customer", strings.NewReader(requestBody))

	// PENTING: Set header agar handler tahu data yang dikirim adalah JSON
	request.Header.Set("Content-Type", "application/json")

	// 3. Siapkan perekam respon seperti kodemu
	recorder := httptest.NewRecorder()

	// 4. Eksekusi Handlernya langsung
	CreateCustomerHandler(recorder, request)

	// 5. Ambil hasil rekaman untuk divalidasi (Assertion)
	response := recorder.Result()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	// Cek status code (Ekspektasi kita: 201 Created)
	if response.StatusCode != http.StatusCreated {
		t.Errorf("Harusnya status 201, tapi dapet %d", response.StatusCode)
	}

	// Cetak hasilnya ke terminal test
	println("Response Body dari Test:", string(body))
}

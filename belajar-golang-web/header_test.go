package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Response Header adalah header yang dikirim oleh server
// Request Header adalah header yang dikirim oleh client

func RequestHeader(w http.ResponseWriter, r *http.Request) {
	// Header tipenya sama seperti query yaitu map slice atau url.Values
	contentType := r.Header.Get("content-type")

	fmt.Fprint(w, contentType)
}

func ResponseHeader(writer http.ResponseWriter, request *http.Request) {
	// Untuk menambahkan response header gunakan ini
	writer.Header().Add("X-Powered-By", "Hamam Priyatmoko")

	fmt.Fprint(writer, "OK")
}
func TestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

	fmt.Println(response.Header.Get("X-Powered-By"))
}



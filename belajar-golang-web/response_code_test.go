package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		// Http menyediakan data status code juga, jadi klo kita tidak pasti
		// berapa codenya kita bisa gunakan nama codenya
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "name: must have")
	} else {
		writer.WriteHeader(http.StatusOK)
		fmt.Fprintf(writer, "Hallo %s", name)
	}

}

func TestResponseCodeInvalid(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	// Ini mengembalikan object ResponseRecorder
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	// Ini mengembalikan object Response
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}

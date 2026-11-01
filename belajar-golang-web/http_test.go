package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

// NewRequest(method, url, body) merupakan function yang digunakan untuk membuat http.Request
// Kita bisa menentukan method, url dan body yang akan kita kirim sebagai simulasi unit test
// Selain itu, kita juga bisa menambahkan informasi tambahan lainnya pada request yang ingin kita kirim,
// seperti header, cookie, dan lain-lain

// NewRecorder() akan mengembalikan sebuah recorder yang mengimplementasikan ResponseWriter
// ResponWriter/ResponseRecorder merupakan struct bantuan yang merekam HTTP RESPONSE dari hasil
// testing yang kita lakukan, bisa nangkep response body, status dll
func TestHttp(t *testing.T) {
	// Dengan test ini kita tidak perlu menjalankan servernya bisa langsung test
	request := httptest.NewRequest("GET", "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()
	// Body akan membalikan io.ReadCloser
	body, _ := io.ReadAll(response.Body)
	fmt.Println(response)
	fmt.Println(body)
	fmt.Println(string(body))
}

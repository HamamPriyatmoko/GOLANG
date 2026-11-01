package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// kita tahu bahwa saat kita membuat form, kita bisa submit datanya dengan method GET atau POST
// Jika menggunakan method GET, maka hasilnya semua data di form akan menjadi query parameter
// Sedangkan jika menggunakan POST, maka semua data di form akan dikirim via body HTTP request

// Semua data form post secara otomatis akan disimpan dalam attribute Request.PostForm
// Namun sebelum kita bisa mengambil data di attribute PostForm kita wajib memanggil
// method Request.ParseForm() terlebih dahulu, method ini digunakan untuk melakukan
// parsing data body apakah bisa di parsing menjadi form data atau tidak

func FormPost(writer http.ResponseWriter, request *http.Request) {
	// sebelum kita bisa mengambil data di attribute PostForm kita wajib memanggil
	// method Request.ParseForm() terlebih dahulu, ini akan melakukan parsing data body menjadi form data
	if err := request.ParseForm(); err != nil {
		panic(err)
	}

	// Jika menggunakan PostFormValue dia akan otomatis ParseFormnya, karena dari logicnya,
	// dia sudah melakukan ParseForm, jadi kita tidak perlu lagi secara manual melakukan ParseForm
	// request.PostFormValue("first_name")

	first_name := request.PostForm.Get("first_name")
	last_name := request.PostForm.Get("last_name")

	fmt.Fprintf(writer, "Hallo %s %s, good morning", first_name, last_name)
}

func TestFormPost(t *testing.T) {
	// Karena body itu bertipe (Reader) maka cara membuat body adalah harus membuat object Readernya
	requestBody := strings.NewReader("first_name=Hamam&last_name=Priyatmoko")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", requestBody)
	// Harus tambahkan content bertipe form agar web server tahu bahwa ini mengirim body
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

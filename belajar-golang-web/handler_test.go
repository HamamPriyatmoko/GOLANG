package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	// Server untuk bisa menerima Http Request masuk ke server kita perlu yang namanya HANDLER
	// HANDLER digolang itu merupakan sebuah interface
	// HANDLER sendiri memiliki function yang wajib dibuat yaitu ServeHTTP() dimana berfungsi sebagai
	// function yang akan dieksekusi ketika menerima HTTP Request
	// Karena HANDLER itu sendiri hanya berupa kontrak, tidak bisa langsung kita pakai tetapi golang menyediakan
	// type function HandlerFunc sebagai implementasi dari interface HANDLER, kita bisa menggunakanya
	// untuk membuat function Handler HTTP

	// Writer disini sebagai response yang akan diberikan ketika ada sebuah permintaan/request
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		// Ini akan ngeprint ke Writernya/ lempar nilainya ke writer
		fmt.Fprint(w, "Hello World")
	}

	var server http.Server = http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

	// Selain HandlerFunc implementasi dari Handler yang lain adalah ServeMux
	// Bisa lihat lebih detail di servemux_test.go 
}

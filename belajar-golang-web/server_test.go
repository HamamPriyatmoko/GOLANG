package belajar_golang_web

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	// Server ini hanya untuk membuat Web Servernya
	server := http.Server{
		Addr: "localhost:8080",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

	// Agar server ini bisa menerima request diperlukan yang namanya Handler, namun
	// digolang Handler merupakan interface jadi kita perlu membuat fungsi yang memenuhi kontrak Handler
	// akan tetapi golang sudah menyediakan implementasi Handler itu, jadi kita tidak perlu
	// membuatnya dari low level, bisa dilihat difile handler_test.go dan servemux_test.go
}

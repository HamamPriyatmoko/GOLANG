package belajar_golang_web

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {

	directory := http.Dir("./resources")
	// File server digunakan sebagai handler untuk static file server dengan ini kita tidak perlu men-load manual
	// FileServer sendiri merupakan handler jadi kita bisa tambahkan kedalam http.server ataupun ServeMux
	fileServer := http.FileServer(directory)

	mux := http.NewServeMux()

	// Dalam kenyataanya jika ada request masuk kedalam url static ini
	// maka yang dilakukan oleh method handle itu dia akan merefer ke
	// folder resources dan didalam resources dia akan mencari path static
	// sesuai dengan url yang ditambahkan tetapi karena kita hanya ingin ke path
	// "/resource/img.jpg" bukan "/resource/static/img.jpg" maka kita gunakan http.StripPrefix
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("Terjadi Masalah Pada Server")
		panic(err)
	}
}

// Digolang embed nama folder ikut menjadi nama resource nya, misal resource/index.js, jadi untuk
// mengaksesnya kita perlu gunakan url /static/resource/index.js
//
//go:embed resources
var resources embed.FS

func TestFileServerGolangEmbed(t *testing.T) {
	// Dengan menggunakan fs.Sub artinya kita masuk lebih dalam lagi ke sub directorinya
	// ibarat command seperti *cd resources*
	directory, _ := fs.Sub(resources, "resources")

	// http.FS berfungsi untuk mengkonversi resources bawaan golang embed menuju http filesystem
	fileServer := http.FileServer(http.FS(directory))

	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("Terjadi Masalah Pada Server")
		panic(err)
	}
}

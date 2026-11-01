package belajar_golang_web

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

// Target Objek: http.ServeFile menembak satu file spesifik (misal: nota.pdf),
// sedangkan http.FileServer membuka akses ke satu folder penuh (misal: folder /static).

// Sifat URL/Path: http.ServeFile bersifat statis (file yang dikirim selalu sama tidak peduli isi URL-nya),
// sedangkan http.FileServer bersifat dinamis (file yang dikirim otomatis mengikuti nama file
// yang diketik user di ujung URL).

// Tipe Kembalian: http.ServeFile langsung mengeksekusi response tanpa mengembalikan apa pun,
// sedangkan http.FileServer mengembalikan sebuah http.Handler siap pakai.

// Penanganan File Indeks: http.ServeFile tidak peduli pada file index.html,
// sedangkan http.FileServer akan otomatis mencari dan menampilkan index.html jika user hanya
//  mengakses rute foldernya saja.

// Cara Pemasangan: http.ServeFile biasanya dipanggil di dalam fungsi handler (http.HandleFunc),
// sedangkan http.FileServer dipasang sebagai handler rute utama (http.Handle).

// Jadi jika serve file kita
func ServeFile(writer http.ResponseWriter, request *http.Request) {

	if request.URL.Query().Get("name") != "" {
		var url string = "./resources/" + request.URL.Query().Get("name") + ".jpg"

		fmt.Println(url)
		http.ServeFile(writer, request, url)
	} else {
		http.ServeFile(writer, request, "./resources/notfound.html")
	}
}

func TestServeFileServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	server.ListenAndServe()
}

//go:embed resources/vario.jpg
var resourceOk string

//go:embed resources/notfound.html
var resourceNotFound string

func ServeFileEmbed(writer http.ResponseWriter, request *http.Request) {
	// fmt.Println(resourceOk + " " + resourceNotFound)
	if request.URL.Query().Get("name") != "" {
		fmt.Fprint(writer, resourceOk)
	} else {
		fmt.Fprint(writer, resourceNotFound)
	}
}

func TestServeFileEmbed(t *testing.T) {
	// fmt.Println(resourceOk + " " + resourceNotFound)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Println("There are an error in server")
	}

}

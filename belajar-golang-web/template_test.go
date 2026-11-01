package belajar_golang_web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SimpleHtml(writer http.ResponseWriter, request *http.Request) {
	// Yang titik didalam string template text nantinya akan disubtitusikan
	templateText := `<html><body>{{.}}<body><html>`
	// variable t adalah data templatenya
	// t, err := template.New("SIMPLE").Parse(templateText)

	// Dengan menggunakan must ini kita tidak perlu manual melakukan pengecekan errnya
	// New akan membuatkan template baru dan parse yang akan mengubah content file dari teks
	// menjadi html template
	t := template.Must(template.New("SIMPLE").Parse(templateText))
	// if err != nil {
	// 	fmt.Println(err)
	// 	panic(err)
	// }

	// Execute template inilah yang akan mengeksekusi templatenyaa dan akan mengirimkannmya ke
	// http ResponseWriter
	// Parameter ke-3 itu yang akan subtitusi dengan titik didalam templatenya
	t.ExecuteTemplate(writer, "SIMPLE", "HELLO HTML TEMPLATE")
}

func TestSimpleHtml(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/home", SimpleHtml)
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func SimpleHTMLFile(writer http.ResponseWriter, request *http.Request) {
	// ini akan memparsing file jadi template dan pastikan agar tipe filenya adalah .gohtml
	t, err := template.ParseFiles("./templates/simple.gohtml")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML Template")
}

func TestSimpleHTMLFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(SimpleHTMLFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TemplateDirectory(writer http.ResponseWriter, request *http.Request) {
	// Jika ingin parse dari directory pake ParseGlob
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))

	t.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML Template")
}

func TestTemplateDirectory(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDirectory(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

//go:embed templates/*.gohtml
var templates embed.FS

func TemplateEmbed(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML Template")
}

func TestTemplateEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateEmbed(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

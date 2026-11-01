package belajar_golang_web

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"testing"
)

// Template Caching ini adalah cara agar kita tidak perlu melakukan parsing terus-menerus ketika
// handlernya dieksekusi, disetiap kita menambahkan code parsing didalam handler artinya
// setiap handler dieksekusi maka dia akan melakukan parsing untuk templatenya dan ini menambah
// waktu eksekusi yang mana akan menjadi masalah dikemudian hari

//go:embed templates/*.gohtml
var templates2 embed.FS

// Jika parsingnya ditaruh diglobal atau diluar function artinya dia akan dieksekusi sekali
// ketika aplikasi berjalan
var myTemplates = template.Must(template.ParseFS(templates2, "templates/*.gohtml"))

func TemplateCaching(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "simple.gohtml", "Hello Template Caching")
}

func TestTemplateCaching(t *testing.T) {
	// request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	// recorder := httptest.NewRecorder()

	// TemplateCaching(recorder, request)

	// body, _ := io.ReadAll(recorder.Result().Body)

	// fmt.Println(string(body))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateCaching),
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}

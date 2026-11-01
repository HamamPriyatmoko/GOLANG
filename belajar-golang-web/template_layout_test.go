package belajar_golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateLayout(writer http.ResponseWriter, request *http.Request) {
	// ParseFiles mengubah file template menjadi suatu Tamplate dan penamaan templatenya akan sesuai
	// dengan nama filenya dan jika terdapat banyak file maka nama templatenya akan mengikuti
	// file terakhirnya
	t := template.Must(template.ParseFiles(
		"./templates/header.gohtml",
		"./templates/footer.gohtml",
		"./templates/layout.gohtml",
	))
	t.ExecuteTemplate(writer, "layout", map[string]interface{}{
		"Title": "Template Layout",
		"Name":  "Hamam",
	})
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
	// server := http.Server{
	// 	Addr:    "localhost:8080",
	// 	Handler: http.HandlerFunc(TemplateLayout),
	// }

	// if err := server.ListenAndServe(); err != nil {
	// 	fmt.Println(err)
	// 	panic(err)
	// }
}

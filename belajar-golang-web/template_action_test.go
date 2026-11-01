package belajar_golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Operator perbandingan pada action template go
// eq artinya arg1 == arg2
// ne artinya arg1 != arg2
// lt artinya arg1 < arg2
// le artinya arg1 <= arg2
// gt artinya arg1 > arg2
// ge artinya arg1 >= arg2

type Address struct {
	City   string
	Region string
}

type Page2 struct {
	Title      string
	Name       string
	Address    Address
	Action     string
	FinalValue int
}

func TemplateAction(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/action.gohtml"))
	t.ExecuteTemplate(writer, "action.gohtml", Page2{
		Title: "Template Data Struct",
		Name:  "Hamam",
		Address: Address{
			City: "Tegal",
		},
	})
}

func TestTemplateAction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAction(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

func TemplateActionRange(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))
	t.ExecuteTemplate(writer, "range.gohtml", map[string]interface{}{
		"Title": "Range",
		"Hobbies": []string{
			"Gaming", "Reading", "Coding",
		},
	})
}

func TestTemplateActionRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionRange(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

func TemplateActionOperator(writer http.ResponseWriter, request *http.Request) {
	// ParseFiles mengubah file template menjadi suatu Tamplate dan penamaan templatenya akan sesuai
	// dengan nama filenya
	t := template.Must(template.ParseFiles("./templates/perbandingan.gohtml"))
	t.ExecuteTemplate(writer, "perbandingan.gohtml", Page2{
		Title:      "Template Data Struct",
		Name:       "Hamam",
		FinalValue: 80,
	})
}

func TestTemplateActionOperator(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionOperator(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

func TemplateActionWith(writer http.ResponseWriter, request *http.Request) {
	// ParseFiles mengubah file template menjadi suatu Tamplate dan penamaan templatenya akan sesuai
	// dengan nama filenya
	t := template.Must(template.ParseFiles("./templates/with.gohtml"))
	t.ExecuteTemplate(writer, "with.gohtml", Page2{
		Title: "Template Data Struct",
		Name:  "Hamam",
		Address: Address{
			Region: "Indonesia",
			City:   "Tegal",
		},
	})
}

func TestTemplateActionWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionWith(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

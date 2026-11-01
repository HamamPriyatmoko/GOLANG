package belajar_golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MyPage struct {
	Name string
}

func (myPage *MyPage) SayHello(name string) string {
	return "Hello " + name + ", My Name is " + myPage.Name
}

func TemplateFunction(writer http.ResponseWriter, request *http.Request) {
	// Memanggil function sayhello dan mengirimkan parameter Budi
	// Konsep:
	// Function yang dipanggil adalah function yang ada didalam struct mypage
	t := template.Must(template.New("FUNCTION").Parse(`{{.SayHello "Budi"}}`))

	// Karena functionnya berada dipointer MyPage jadi kita harus manggilnya juga struct pointernya
	t.ExecuteTemplate(writer, "FUNCTION", &MyPage{
		Name: "Hamam",
	})
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

func TemplateFunctionGlobal(writer http.ResponseWriter, request *http.Request) {
	// Untuk mengakses func global tidak perlu ditambahkan titik "."
	// karena jika menggunakan titik artinya dia mengaksesk template data
	t := template.Must(template.New("FUNCTION").Parse(`{{len .Name}}`))

	t.ExecuteTemplate(writer, "FUNCTION", &MyPage{
		Name: "Hamam",
	})
}

func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobal(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

func TemplateFunctionCreateGlobal(writer http.ResponseWriter, request *http.Request) {

	t := template.New("FUNCTION")
	// Untuk membuat function global sendiri gunakan Funcs
	// Dan function global perlu diregistrasi dulu sebelum diparse
	t = t.Funcs(map[string]interface{}{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	t = template.Must(t.Parse(`{{ upper .Name}}`))

	t.ExecuteTemplate(writer, "FUNCTION", &MyPage{
		Name: "Hamam",
	})
}

func TestTemplateFunctionCreateGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionCreateGlobal(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

// Function pipelines adalah haisl dari function sebelumnya bisa langsung dikirim
// ke function berikutnya

func TemplateFunctionPipelines(writer http.ResponseWriter, request *http.Request) {

	t := template.New("PIPELINES")
	// Untuk membuat function global sendiri gunakan Funcs
	// Dan function global perlu diregistrasi dulu sebelum diparse
	t = t.Funcs(map[string]interface{}{
		"sayHello": func(name string) string {
			return "Hello " + name
		},
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	t = template.Must(t.Parse(`{{ sayHello .Name | upper}}`))

	t.ExecuteTemplate(writer, "PIPELINES", &MyPage{
		Name: "Hamam",
	})
}

func TestTemplateFunctionPipelines(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionPipelines(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

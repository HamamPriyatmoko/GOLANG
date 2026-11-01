package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

// Disclaimer:
// Query Parameter itu case sensitive

func SayHello(writer http.ResponseWriter, request *http.Request) {
	// Jika pake get value dari parameter yang sama yang didapet cuman 1 value aja
	name := request.URL.Query().Get("name")
	// fmt.Fprint(writer, name)
	fmt.Fprintf(writer, "Hello %s, good morning", name)
	// fmt.Println(name)
}

func MultipleParameter(writer http.ResponseWriter, request *http.Request) {
	// Jika pake get value dari parameter yang sama yang didapet cuman 1 value aja
	name := request.URL.Query().Get("name")
	age := request.URL.Query().Get("age")
	intAge, _ := strconv.Atoi(age)
	// fmt.Fprint(writer, name)
	fmt.Fprintf(writer, "Hello %s, good morning i am %d years old", name, intAge)
	// fmt.Fprintf(writer, "i am %d years old", intAge)
	// fmt.Println(name)
}

func MultipleValueQuery(writer http.ResponseWriter, request *http.Request) {
	// Query akan mengembalikan tipe data Value dimana tipe ini merupakan map[string][]string
	var query url.Values = request.URL.Query()
	name := query["name"]
	fmt.Fprintln(writer, strings.Join(name, ","))
	fmt.Println(name)
	data := reflect.ValueOf(query)
	fmt.Println(data)
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Hamam", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	// Body akan membalikan io.ReadCloser
	body, _ := io.ReadAll(response.Body)
	// fmt.Println(response)
	// fmt.Println(body)
	fmt.Println(string(body))
}

func TestMultipleQuery(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Hamam&age=15", nil)
	recorder := httptest.NewRecorder()

	MultipleParameter(recorder, request)

	response := recorder.Result()
	// Body akan membalikan io.ReadCloser
	body, _ := io.ReadAll(response.Body)
	// fmt.Println(response)
	// fmt.Println(body)
	fmt.Println(string(body))
}

func TestMultipleValueQuery(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Hamam&name=Silvya", nil)
	recorder := httptest.NewRecorder()

	MultipleValueQuery(recorder, request)

	response := recorder.Result()
	// Body akan membalikan io.ReadCloser
	body, _ := io.ReadAll(response.Body)
	// fmt.Println(response)
	// fmt.Println(body)
	fmt.Println(string(body))
}

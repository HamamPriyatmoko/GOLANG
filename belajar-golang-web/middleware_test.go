package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

type logMiddleware struct {
	Handler http.Handler
}

func (middleware *logMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// Disinilah kita bisa menambahkan logic misal login sebelum akses webnya
	fmt.Printf("Before handler execute \n")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Printf("After handler execute \n")
}

type ErrorHandler struct {
	Handler http.Handler
}

func (errorHandler *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Terjadi Error")
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Error: %s", err)
		}
	}()
	errorHandler.Handler.ServeHTTP(writer, request)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler Executed")
		fmt.Fprint(writer, "Hello Middleware")
	})
	// Mau buat handler sebanyak apapun mereka akan tetap eksekusi middlewarenya karena kita
	// sudah menangani mux akan dipassing ke logmiddleware
	mux.HandleFunc("/foo", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Foo Executed")
		fmt.Fprint(writer, "Hello Foo")
	})

	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Foo Executed")
		panic("ups")
	})

	// logMiddleware := new(logMiddleware)
	// logMiddleware.Handler = mux

	logMiddleware := &logMiddleware{
		Handler: mux,
	}

	errorHandler := &ErrorHandler{
		Handler: logMiddleware,
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: errorHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

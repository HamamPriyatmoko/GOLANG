package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

// Named parameter contoh:
// /products/:id

// Catch All Parameter untuk menangkap semua parameter contoh:
// /products/*filepath

func TestParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id := p.ByName("id")
		text := "Product " + id
		fmt.Fprint(w, text)
	})

	// server := http.Server{
	// 	Addr:    "localhost:8888",
	// 	Handler: router,
	// }

	// server.ListenAndServe()

	request := httptest.NewRequest("GET", "http://localhost:8888/products/Hamam", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Product Hamam", string(body))
}

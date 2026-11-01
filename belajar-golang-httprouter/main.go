package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// Pengganti ServeMux
	// Saat menggunakanm servemux ketika kita ingin membuat route biasanya menggunakan
	// http.Handler berbeda dengan Router kita tidak menggunakan http handler melainkan
	// httprouter.Handle akan tetapi handle ini memiliki 3 parameter
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello Http Router")
	})

	server := http.Server{
		Handler: router,
		Addr:    "localhost:8888",
	}

	server.ListenAndServe()
}

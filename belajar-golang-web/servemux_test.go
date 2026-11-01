package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

type ApiCar struct {
	Nama string
	Umur int64
}

func (a ApiCar) ServeHTTP(write http.ResponseWriter, read *http.Request) {
	fmt.Fprint(write, "This is response from api profile, you are ", a.Nama, "your age ", a.Umur, "years old")
}

func TestServeMux(t *testing.T) {
	// ServeMux adalah alternative lain implementasi HANDLER, karena ServeMux memiliki kelebihan lain
	// yaitu bisa membuat banyak sekali route/endpoint url sedangkan HandleFunc itu tidak
	// ServeMux bisa dibilang sebagai Router jika dibahasa pemrograman seperti javscript dll
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hallo User")
	})

	mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "This is an api")
	})

	// Kalo kita tambahkan slash setelah pattern urlnya maka apapun yang berkaitan dengan images
	// mau itu urlnya /images/car, /images/core jika mereka tidak dibuatkan handlenya sendiri
	// mereka akan masuk kedalama handle ini
	mux.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "This is images urls and yup you come back again")
	})

	// Namun jika terdapat URL Pattern yang lebih panjang, maka akan diprioritaskan yang lebih
	// panjang, misal jika terdapat URL /images/ dan /images/thumbnails/, maka jika mengakses
	// /images/thumbnails/ akan mengakses /images/thumbnails/, bukan /images

	// Jadi akan diprioritaskan yang panjang duluan
	mux.HandleFunc("/images/core", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "This is images core urls")
	})

	var hamam ApiCar = ApiCar{
		Nama: "Hamam",
		Umur: 22,
	}

	mux.Handle("/handler", hamam)

	var Auth http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "You are succes login")
	}

	mux.Handle("/api/auth", Auth)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func TestRequest(t *testing.T) {
	// Request merupakan sebuah struct
	// 	Request adalah struct yang merepresentasikan HTTP Request yang dikirim oleh Web Browser
	// Semua informasi request yang dikirim bisa kita dapatkan di Request
	// Seperti, URL, http method, http header, http body, dan lain-lain

	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, r.Method)
		// fmt.Fprint(w, r.Header)
		// fmt.Fprint(w, r.RequestURI)
		fmt.Println(r.RequestURI)
		// fmt.Fprint(w, r.URL)
		fmt.Fprint(w, r.RemoteAddr)
	}

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: handler,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}

// NewRecorder() akan mengembalikan sebuah recorder yang mengimplementasikan ResponseWriter
// ResponWriter/ResponseRecorder merupakan struct bantuan yang merekam HTTP RESPONSE dari hasil
// testing yang kita lakukan, bisa nangkep response body, status dll
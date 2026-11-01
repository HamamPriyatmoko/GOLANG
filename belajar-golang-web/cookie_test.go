package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func SetCookie(write http.ResponseWriter, req *http.Request) {
	// new() berfungsi untuk membuat pointer
	cookie := new(http.Cookie)
	cookie.Name = "x-api-key"
	cookie.Value = req.URL.Query().Get("name")
	cookie.Path = "/"

	// Set cookie ke response writernya
	http.SetCookie(write, cookie)
	fmt.Fprint(write, "Success create cookie")
}

func GetCookie(write http.ResponseWriter, req *http.Request) {
	// Mengambil cookie
	cookie, err := req.Cookie("x-api-key")
	cookie.Expires = time.Date(2026, 8, 15, 15, 44, 30, 0, time.UTC)
	if err != nil {
		// Untuk membuat status codenya
		write.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(write, "No Cookie")
	} else {
		// Browser hanya akan mengembalikan cookienya dengan dua data yaitu name dan value
		// selain name dan value itu akan dipotong otomatis oleh browser karna semua data itu akan 
		// ditangani oleh browser itu sendiri

		// Mengambil data cookienya/ value cookienya
		fmt.Fprintf(write, "Hallo %s", cookie.Value)
		// fmt.Fprintf(write, "%s", cookie.Expires.String())
		fmt.Println(cookie.Expires)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/get-cookie", GetCookie)
	mux.HandleFunc("/set-cookie", SetCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?name=Hamam", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)
	cookies := recorder.Result().Cookies()

	for _, cookie := range cookies {

		fmt.Printf("Cookie %s:%s \n", cookie.Name, cookie.Value)

	}
}

func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	cookie := new(http.Cookie)
	cookie.Name = "x-api-key"
	cookie.Value = "Ini Cookie"
	cookie.Expires = time.Date(2026, time.August, 15, 15, 55, 15, 0, time.Local)
	request.AddCookie(cookie)

	// waktu := time.Date(2026, 07, 15, 15, 55, 15, 0, time.UTC)

	// fmt.Println(waktu.String())

	recorder := httptest.NewRecorder()
	GetCookie(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

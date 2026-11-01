package belajar_golang_web

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func UploadForm(writer http.ResponseWriter, request *http.Request) {
	// Ini untuk merender file htmlnya
	myTemplates.ExecuteTemplate(writer, "upload.form.gohtml", nil)
}

func Upload(writer http.ResponseWriter, request *http.Request) {
	// ParseMulti ini untuk setting ukuran file yang bisa diterima berapa
	// untuk defaultnya 32 MB
	// request.ParseMultipartForm(32 << 20)
	// FormFile Untuk mengambil file
	// FormFile ini akan melakukan parsing terlebih dahulu
	// FormFile ini mengembalikan 3 value yaitu file, fileheader, dan err
	file, fileHeader, err := request.FormFile("file")
	if err != nil {
		panic(err)
	}
	// Untuk membuat folder yang akan kita gunakan sebagai penyimpanan untuk file
	// Cara membuatnya adalah gunakan package os method Create
	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}
	// Contoh error jika kita tidak memiliki permission
	// Untuk menyimpan datafilenya ke folder gunakan package io function Copy
	// File akan disimpan difolder resources
	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}

	// Untuk mengambil data yang bukan file
	name := request.PostFormValue("name")
	myTemplates.ExecuteTemplate(writer, "upload.success.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}

func TestUploadForm(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	// http.StripPrefix akan menghapus prefix diawal yaitu "static"
	// http.Dir akan mengubah string biasa menjadi string yang sudah dicustom type menjadi type FileSystem
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/img.jpg
var uploadFileTest []byte

func TestUploadFile(t *testing.T) {
	// new ini akan mengembalikan object pointer
	body := new(bytes.Buffer)

	// Karena body masih kosong dan kita ingin body ini berisi data berformat multipart
	// maka kita perlu isi body dengan membuat format multipart terlebih dahulu
	writer := multipart.NewWriter(body)
	// Setelah kita membuat format multipart tambahakan datanya
	writer.WriteField("name", "Hamam Priyatmoko")
	// Disini file masih belum ada datanya/masih kosong
	file, _ := writer.CreateFormFile("file", "CONTOHUPLOAD.png")
	// Disinilah kita menambahkan data kedalam filenya
	file.Write(uploadFileTest)
	// Close untuk memastikan tidak ada memory yang ngegantung
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/upload", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	Upload(recorder, request)

	bodyResponse, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyResponse))
}

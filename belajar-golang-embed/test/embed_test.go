package test

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

// Package embed adalah fitur baru untuk mempermudah membaca isi file
// pada saat compile time secara otomatis dimasukkan isi file nya dalam variable
// cara menggunakan embed adalah pakai commentar
// Selajutnya kita bisa tambahkan komentar *//go:embed* diikuti dengan nama file nya, diatas
// variable yang kita tuju

//go:embed version.txt
var version string

// Variable yang dituju tersebut nanti secara otomatis akan berisi konten file yang
// kita inginkan secara otomatis ketika kode golang di compile
// Variable yang dituju tidak bisa disimpan di dalam function

func TestString(t *testing.T) {
	fmt.Println(version)
}

//go:embed images.jpg
var images []byte

func TestByte(t *testing.T) {
	err := os.WriteFile("image_new.jpg", images, fs.ModePerm)
	if err != nil {
		panic(err)
	}
	fmt.Println(images)
}

//go:embed belajar-golang-embed/Files/a.txt
//go:embed belajar-golang-embed/Files/b.txt
//go:embed belajar-golang-embed/Files/c.txt
var files embed.FS

func TestMultipleFiles(t *testing.T) {
	a, _ := files.ReadFile("Files/a.txt")
	fmt.Println(string(a))

	b, _ := files.ReadFile("Files/b.txt")
	fmt.Println(string(b))

	c, _ := files.ReadFile("Files/c.txt")
	fmt.Println(string(c))
}

//go:embed Files/*.txt
var path embed.FS

func TestPathMatcher(t *testing.T) {
	dirEntries, _ := path.ReadDir("Files")

	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, _ := path.ReadFile("Files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}

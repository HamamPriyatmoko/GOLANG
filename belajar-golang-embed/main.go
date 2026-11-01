package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
)

//go:embed version.txt
var version string

//go:embed images.jpg
var logo []byte

//go:embed Files/*.txt
var path embed.FS

// Perlu diketahui, bahwa hasil embed yang dilakukan oleh package embed adalah
// permanent dan data file yang dibaca disimpan dalam binary file golang nya
// Artinya bukan dilakukan secara realtime membaca file yang ada diluar
// Hal ini menjadikan jika binary file golang sudah di compile, kita tidak butuh lagi file external nya,
// dan bahkan jika diubah file external nya, isi variable nya tidak akan berubah lagi

// Jadi jika kita sudah lakukan build dan ketika kita rubah isi file, hasil execute atau hasil contentnya
// akan seperti content diawal sebelum dibuild walaupun kita merubah isi content filenya setelah build

func main() {

	fmt.Println(version)

	err := os.WriteFile("image_new.jpg", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	dirEntries, _ := path.ReadDir("Files")

	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, _ := path.ReadFile("Files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}

package main

import (
	"fmt"
	"os"
)

// io.Reader adalah sebuah kontrak atau interface yang diberikan oleh Go
// Reader berfungsi untuk membaca data dari mana saja
// Reader biasanya digunakan ketika untuk membaca data dari
// input User seperti body atau argument terminal atau Membaca isi file

// io.Writer adalah sebuah kontrak atau interface yang diberikan oleh Go
// Writer berfungsi untuk menulis data ke mana saja

func main() {
    // os.Stdout adalah Writer
    fmt.Fprintln(os.Stdout, "Ini menulis ke layar")

    // File juga adalah Writer, file menggunakan kontrak Writer yang diimplementasikan oleh
	// *os.File jadi *os.File ini yang akan mengimplementasikan Writer oleh karena itu
	// fmt.Fprintln akan secara otomatis menjalankan method Write yang udah dibuat oleh *os.File
	// dan terjadilah fungsi penulisan ke file test.txt
    file, _ := os.Create("test.txt")
    fmt.Fprintln(file, "Ini menulis ke file")
}

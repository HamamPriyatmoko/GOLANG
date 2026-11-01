package main

import (
	"bufio"
	"os"
)

func main() {
	// Parameternya menerima io.Writer/ Tujuan dimana teks itu mau ditulis dan os.Stdout bertipe io.Writer 
	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("Hello World\n")
	_, _ = writer.WriteString("Selamat Belajar\n")
	// Saat Menulis Data ditampung di RAM dulu. Begitu penuh (atau saat kamu panggil Flush()), 
	// barulah data dikirim sekaligus ke file.
	writer.Flush()
}
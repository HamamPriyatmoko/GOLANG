package main

import (
	"encoding/csv"
	"os"
)

func main() {
	// os.Stdout juga dibaca atau dikenal sebagai File
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"Hamam","Priyatmoko"})
	_ = writer.Write([]string{"Silvya","Amanda"})
	_ = writer.Write([]string{"Iqbal","Ramadan"})

	// Untuk mengirim semua perubahannya sekaligus ke Terminal/Stdout tp jika menggunakan File Csv maka
	// akan dikirim ke file csv yang sudah dibuat
	writer.Flush()
}
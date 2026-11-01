package belajar_golang_database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Cara kerja database pooling adalah semisal aplikasi mendapatkan traffic tinggi 100 request masuk
// maka kita akan membuat 100 connection ke-database, jika tiba tiba ada 1000 request maka akan
// membuat 1000 connection kedatabase dan isi sangat bahaya karena database punya batasan maksimal
// koneksinya dan biasanya jumlahnya kecil dan untuk mysql sendiri sekitar 150 connection maksimalnya

// Adapun minimal connection untuk jaga jaga semisal tiba-tiba traffic tinggi, karena jika
// tidak diberi minimal maka connection akan membuat ulang connectionnya yang mana memakan waktu lebih lama
// karena jika kita menetapkan minimal database itu akan menyiapkan langsung diawal
func GetConnection() *sql.DB {
	// Driver nya akan otomatis conversi dari date ke string
	db, err := sql.Open("mysql", "root:Hamam@3415@tcp(localhost:3306)/golang_database?parseTime=true")
	if err != nil {
		panic(err)
	}

	// Ini adalah minimal koneksi yang akan dibuat ketika aplikasi pertama kali berjalan
	db.SetMaxIdleConns(10)

	// Ini adalah maximal connection kedatabase yang boleh dibuat
	db.SetMaxOpenConns(100)

	// Pengaturan berapa lama koneksi yang sudah tidak digunakan akan dihapus
	// Masa Bengong, maksudnya adalah ketika ada 20 koneksi telah dibuat lalu dan selesai digunakan
	// dia akan menunggu dalam waktu yang kita tentukan lalu jika sudah melewati waktunya dia akan
	// menghapus koneksinya
	db.SetConnMaxIdleTime(5 * time.Second)

	// Ini merupakan berapa lama koneksi boleh digunakan, misal kita punya 20 koneksi lalu kita set
	// waktunya satu jam maka koneksi itu setelah satu jam akan diclose tp jika ada koneksi minimum
	// maka akan dibuat ulang koneksinya dan akan tersedia lagi 10 koneksi sesuai dengan yang kita set up
	db.SetConnMaxLifetime(20 * time.Second)

	return db
}

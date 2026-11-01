package belajar_golang_database

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	// Karena kita hanya butuh initnya saja dan tidak butuh function lainnya bisa gunakan blank identifier
	// atau (_) agar golang tidak menjadikanya error
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	// sql.Open akan mengembalikan struct *sql.DB dan fungsinya untuk membuat koneksi baru
	// sql.DB sebetulnya bukan koneksi ke database, melainkan sebuah pool ke database
	db, err := sql.Open("mysql", "root:Hamam@3415@tcp(localhost:3306)/golang_database")
	if err != nil {
		panic(err)
	}

	// db.Query("SELECT * FROM ")

	// Jika objct sql.DB sudah selesai digunakan disarankan untuk menutupnya menggunakan ini
	// karena biar tidak terjadi connection leak karena biasanya databse punya maksimal open connection
	// jika tidak ditutup maka akan kepenuhan koneksinya didatabse dan kita perlu kill connectionnya secara
	// manual didatabsenya
	defer db.Close()
}

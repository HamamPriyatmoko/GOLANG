package belajar_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

// Paralel: memecahkan suatu masalah dengan cara membaginya menjadi yang terkecil kecil dan
// dijalankan secara bersamaan pada waktu yang bersamaan

// Process adalah sebuah eksekusi program
// Process mengkonsumsi memory besar
// Process saling terisolasi dengan process lain
// Process lama untuk dijalankan dihentikan
// Process ibarat kita menjalankan sebuah aplikasi, process(membuka aplikasi) microsoft office, process(membuka aplikasi) google chrome, process(membuka aplikasi) zoom.
// Secara memory lokasi memory dari proccess itu berbeda beda

// Thread adalah segmen dari process
// Thread menggunakan memory kecil
// Thread bisa saling berhubungan jika dalam process yang sama
// Thread cepat untuk dijalankan dan dihentikan
// Thread ibarat Tab yang ada digoogle chrome, dan setiap thread itu menjalankan tab berbeda beda

// concurrency adalah menjalankan beberapa pekerjaan secara bergantian

// CPU-Bound adalah sebuah kondisi di mana kecepatan penyelesaian suatu tugas (task) sangat
// bergantung pada kecepatan prosesor (CPU), bukan pada kecepatan memori, jaringan, atau hard disk

// I/O-Bound adalah dimana biasanya algoritma atau aplikasinya sangat bergantung pada
// kecepatan input dan output devices yang digunakan, contoh aplikasi seperti membaca data
// dari file, database, dll

// Goroutine adalah sebuah thread ringan atau goroutine running
// didalam thread yang dikelola oleh Go Runtimenya

// Goroutine dijalankan oleh Go Scheduler dalam thread

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	// berjlan asynchronus artinya tidak akan ditunggu sampai function tersebut selesai,
	// goroutin sama aja kaya async await dijavascript, tapi goroutine punya sifat
	// jika semua proses eksekusi function sudah selesai maka semua goroutinenya akan otomatis dikill/dimatikan
	go RunHelloWorld() // Mirip seperti await didalam javascript
	// Tips:
	// goroutine akan tidak cocok jika sebuah functionya itu mengembalikan sebuah value
	// karena tidak akan bisa ditangkap return valuenya

	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}

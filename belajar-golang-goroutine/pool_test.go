package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Sama kaya connection pool didatabse/prisma tp disini kita bisa membuat object pool itu sendiri
func TestPool(t *testing.T) {
	var pool sync.Pool = sync.Pool{
		// Memberikan nilai default jika semua datanya telah habis digunakan dan
		// jika masih ada yang perlu menggunakan data dari pool padahal semua datanya sudah habis
		// maka attribut "New" ini yang akan memberikan nilainya
		New: func() any {
			return "New"
		},
	}

	// Method ini untuk menambahkan/meletakan nilai kedalam pool
	pool.Put("Hamam")
	pool.Put("Priyatmoko")
	pool.Put("Silvya")
	pool.Put("Amanda")

	for i := 0; i < 10; i++ {
		go func() {
			// Get untuk mengambil datanya didalam Object pool
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			// Meletakkan kembali datanya kedalam object pool
			pool.Put(data)
		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("Selesai")
}

package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup, index *int) {
	// Ini akan mengurangi proses/goroutine jika semua goroutine sudah selesai dieksekusi
	// akan melakukan decrement pada proses/goroutine yang perlu ditunggu
	defer group.Done()

	// Ini menambahkan proses/goroutine yang perlu ditunggu sampai proses/goroutinenya
	// selesai dieksekusi
	group.Add(1)

	fmt.Println("Goroutine Ke-", *index)
	time.Sleep(1 * time.Second)
}

// WaitGroup adalah fitur yang bisa digunakan untuk menunggu sebuah proses selesai dilakukan
// misal kita ingin menjalankan beberapa proses menggunakan goroutine, tapi kita ingin
// semua proses selesai terlebih dahulu sebelum aplikasi kita selesai Kasus seperti ini
// bisa menggunakan WaitGroup
func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronous(group, &i)
	}

	// Ini method untuk menunggu proses/goroutinenya selesai
	group.Wait()
	fmt.Println("Selesai")
}

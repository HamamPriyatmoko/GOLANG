package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// sync.Cond adalah menjadi "Ruang Tunggu" (Waiting Room) bagi sekumpulan Goroutine
// yang sedang menunggu suatu kondisi/keadaan tertentu terpenuhi.

// Condition ketika butuh locking tapi pengen jalan klo semisal kita sudah perintahkan untuk jalan

// Akan mengembalikan object pointer Cond
var cond = sync.NewCond(&sync.Mutex{})
var group = &sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)
	cond.L.Lock()
	// Wait artinya akan menunggu atau membuat goroutine tertidur sampai ada yang membangunkannya
	// dengan memancarkan atau triger menggunakan cond.Signal/cond.BroadCast
	fmt.Println("waiting ", value)
	cond.Wait()
	// Didalam wait sendiri siklusnya ketika dieksekusi dia akan Unlock, lalu menunggu, klo diberi Signal
	// maka dia akan Lock lalu eksekusi perintah setelah Wait dan Unlock lagi
	fmt.Println("Done", value)
	cond.L.Unlock()
}

func TestCond(t *testing.T) {

	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		time.Sleep(1 * time.Second)
	// 		// Signal akan membangungkan goroutine satu persatu
	// 		cond.Signal()
	// 	}
	// }()

	go func() {
		time.Sleep(1 * time.Second)
		// Broadcast akan membangungkan semua goroutine
		cond.Broadcast()
		// Walaupun semua goroutine akan bangun serentak hanya ada 1 goroutine yang bisa
		// melakukan Lock
	}()
	group.Wait()
}

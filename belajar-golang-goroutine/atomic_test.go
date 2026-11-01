package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

// Atomic merupakan package yang digunakan untuk menggunakan data primitive secara
// aman pada proses concurrent

// Atomic mempermudah kita dalam melakukan Lock dan Unlock tp hanya untuk data primitive saja

// Karena go menganut Zero value, maka parameter t itu sudah otomatis terdeklarasi dengan object T
func TestAtomic(t *testing.T) {
	group := sync.WaitGroup{}
	var x int64 = 0
	for i := 0; i < 1000; i++ {
		group.Add(1)
		go func() {
			// Kenapa Done didalam goroutine karena jika diluar fungsi goroutine ini
			// maka bisa jadi akan dieksekusi duluan karena go itu bersifat lama atau tidak menunggu
			// jadi go akan eksekusi proses lain semisal goroutine belum selesai dieksekusi
			defer group.Done()
			for j := 0; j < 100; j++ {
				// Atomic akan otomatis melakukan lock unlock pada conccurent goroutine
				atomic.AddInt64(&x, int64(1))
				// Klo datanya berupa struct menggunakan mutex
			}
		}()
	}
	group.Wait()
	// time.Sleep(5 * time.Second)
	fmt.Println("Counter =", x)

}

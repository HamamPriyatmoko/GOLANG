package belajar_golang_goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

// Sudah dijelaskna bahwa Goroutine dijalankan didalam thread,
// untuk mengetahui berapa jumlah thread bisa gunakan "GOMAXPROCS", sebuah function package runtime
// yang bisa kita gunakan untuk mengubah jumlah thread atau mengambil jumlah thread
// Secara default, jumlah thread di Go-Lang itu sebanyak jumlah CPU di komputer kita.
// Kita juga bisa melihat berapa jumlah CPU kita dengan menggunakan function runtime.NumCpu()
func TestGetGomaxprocs(t *testing.T) {
	// Untuk mencari tahu berapa total Logical CPUs
	// ini adalah perkalian dari jumlah core dengan dikali 2
	totalCpu := runtime.NumCPU()
	fmt.Println("CPU:", totalCpu)

	// Mencari tahu total thread dan jika ingin merubahnya angka diparameternya harus lebih dari 1
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Thread:", totalThread)

	// Mencari tahu total Goroutine
	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Goroutine:", totalGoroutine)
	// Kenapa digoroutine ini totalnya 2 padahal kita sedang tidak mendeklarasikan goroutinenya
	// karena untuk menjalankan testing itu menggunakan goroutine dan satu lagi untuk menjalankan
	// "Garbage Collection"
}

func TestChangeThreadNumber(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		group.Add(1)
		go func() {
			time.Sleep(2 * time.Second)
			group.Done()
		}()
	}

	defer group.Wait()

	totalCpu := runtime.NumCPU()
	fmt.Println("CPU:", totalCpu)

	// Meningkatkan threadnya artinya scaling secara vertical
	// Mencari tahu total thread dan mengubahnya
	// Untuk best practice yang disarankan golang adalah sebanyak jumlah cpu corenya
	runtime.GOMAXPROCS(20)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Thread:", totalThread)

	// Mencari tahu total Goroutine
	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Goroutine:", totalGoroutine)

}

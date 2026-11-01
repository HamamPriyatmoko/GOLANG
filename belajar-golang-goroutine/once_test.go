package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOnce() {
	counter++
}

// Once adalah fitur di Go-Lang yang bisa kita gunakan untuk memastikan bahwa
// sebuah function di eksekusi hanya sekali Jadi berapa banyak pun goroutine yang mengakses,
// bisa dipastikan bahwa goroutine yang pertama yang bisa mengeksekusi function nya
func TestOnce(t *testing.T) {
	// DiGo memiliki konsep Zero value Ready-To-Use jadi tanpa perlu deklarasi langsung
	// jadi bisa langsung dipakai hanya dengan menambahkan tipe datanya
	// Contoh:
	// var angka int, didalam variable angka ini akan langsung bernilai 0 walaupun tidak dideklarasikan langsung
	// dan beberapa fitur package dari Go ada yang bisa langsung begini ada juga yang tidak jika semisal
	// dia membutuhkan parameter untuk membuat objectnya
	var group sync.WaitGroup // berarti WaitGroup ini didalamnya sudah menggunakan nilai default 0 atau apapun itu
	var once sync.Once       // ini juga sama Once sudah menggunakan nilai defaultnya atau Zero value

	// Mau cara yang diatas atau cara yang dibawah sama aja mereka memberi zero value, kecuali
	// mau menambahkan value defaultnya diawal
	// group := sync.WaitGroup{}
	// once := sync.Once{} 

	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			// Function yang boleh dimasukan adalah hanya function yang tidak punya parameter
			once.Do(OnlyOnce)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Jumlah counter:", counter)
}

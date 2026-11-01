package belajar_golang_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

// Context merupakan sebuah data yang membawa value, sinyal cancel, sinyal timeout dan sinyal deadline
// Context biasanya dibuat per request (misal setiap ada request masuk ke server web melalui http request)
// Context digunakan untuk mempermudah kita meneruskan value, dan sinyal antar proses

// Context di Golang biasa digunakan untuk mengirim data request atau sinyal ke proses lain
// Dengan menggunakan context, ketika kita ingin membatalkan semua proses, kita cukup
// mengirim sinyal ke context, maka secara otomatis semua proses akan dibatalkan
// Hampir semua bagian di Golang memanfaatkan context, seperti database, http server, http client, dan lain-lain

func TestContext(t *testing.T) {
	// Karena Context adalah sebuah interface, untuk membuat context kita butuh
	// sebuah struct yang sesuai dengan kontrak interface Context Namun kita tidak perlu
	// membuatnya secara manual Di Golang package context terdapat function yang bisa
	// kita gunakan untuk membuat Context

	// Background = Membuat context kosong. Tidak pernah dibatalkan, tidak pernah timeout, dan tidak memiliki
	// value apapun. Biasanya digunakan di main function atau dalam test, atau dalam awal proses request terjadi.
	background := context.Background()
	fmt.Println(background)

	// TODO = Membuat context kosong seperti Background(), namun biasanya menggunakan ini
	// ketika belum jelas context apa yang ingin digunakan
	todo := context.TODO()
	fmt.Println(todo)
}

// Context menganut konsep parent dan child Artinya, saat kita membuat context,
// kita bisa membuat child context dari context yang sudah ada
// Parent context bisa memiliki banyak child, namun child hanya bisa memiliki satu parent context

// Jika kita melakukan cancel terhadap parentnya maka childnya akan ikut tercancel
// Begitu juga ketika menyisipkan data ke parentnya maka childnya akan bisa mendapatkan data tersebut
// tapi jika menyisipkan data hanya ke childnya maka parent tidak bisa mengakses datanya

// Context merupakan object immutable
// Ketika kita menambahkan value ke dalam context, atau menambahkan pengaturan timeout dan yang
// lainnya, secara otomatis akan membentuk child context baru, bukan merubah context tersebut

// Data Pair adalah data yang memiliki struktur (key -> value)

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	// Untuk membuat child atau menambahkan value ke context, kita bisa menggunakan function
	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	// Untuk mengambil value gunakan method Value
	// Flow pengambilan value dari context: akan cari value dari context itu sendiri, klo tidak ada keatasnya lagi(parent)
	fmt.Println("Value:", contextF.Value("f"))                // mendapatkan value
	fmt.Println("Value Parentnya:", contextF.Value("c"))      // mendapatkan value karena c adalah data dari parentnya
	fmt.Println("Value bukan Parentnya", contextF.Value("b")) // tidak bisa mendapatkan value yang bukan parentnya
	fmt.Println("Ambil nilai childnya:", contextA.Value("b")) // tidak bisa mendapatkan data childnya

	fmt.Println(contextD)
	fmt.Println(contextE)
}

// Simulasi goroutine leak
// func CreateCounter() chan int {
// 	destination := make(chan int)
// 	go func() {
// 		defer close(destination)
// 		counter := 1
// 		// Karena ini adalah infinite for walaupun ada yang ngeconsume atau tidak sichannelnya
// 		// goroutine tetap akan menjalankan infinite loopnya
// 		for {
// 			destination <- counter
// 			counter++
// 		}
// 	}()
// 	return destination
// }

// CreateCounter yang sudah mengimplementasikan context untuk menghindari goroutine leak
// func CreateCounter(ctx context.Context) chan int {
// 	destination := make(chan int)

// 	go func() {
// 		defer close(destination)
// 		counter := 1
// 		// Karena ini adalah infinite for walaupun ada yang ngeconsume atau tidak sichannelnya
// 		// goroutine tetap akan menjalankan infinite loopnya
// 		for {
// 			select {
// 			// Kita melakukan listen dari Done ini, Done akan me-return sebuah channel
// 			// jika Done mengembalikan channel artinya sudah dicancel tetapi
// 			// jika tidak artinya belum dilakukan cancel
// 			case <-ctx.Done():
// 				return
// 			default:
// 				destination <- counter
// 				counter++
// 			}
// 		}
// 	}()
// 	return destination
// }

// Simulasi query lambat dan cara penanganannya
func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		// Karena ini adalah infinite for walaupun ada yang ngeconsume atau tidak sichannelnya
		// goroutine tetap akan menjalankan infinite loopnya
		for {
			select {
			// Kita melakukan listen dari Done ini, Done akan me-return sebuah channel
			// jika Done mengembalikan channel artinya sudah dicancel tetapi
			// jika tidak artinya belum dilakukan cancel
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second) // Simulasi slow
			}
		}
	}()
	return destination
}

// Simulasi goroutine leak
// func TestContextWithCancel(t *testing.T) {
// 	fmt.Println("Total Goroutine:", runtime.NumGoroutine())

// 	destination := CreateCounter()
// 	for n := range destination {
// 		fmt.Println("Counter", n)

// 		if n == 10 {
// 			break
// 		}
// 	}

// 	fmt.Println("Total Goroutine:", runtime.NumGoroutine())
// }

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())

	parent := context.Background()

	ctx, cancel := context.WithCancel(parent) // Ini akan membuat context child dan function cancel

	destination := CreateCounter(ctx)

	fmt.Println("Total Goroutine:", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("Counter", n)

		if n == 10 {
			break
		}
	}

	cancel() // mengirim sinyal cancel ke konteks
	// Jangan lupa tambahkan waitgroup atau sleep karena bisa jadi goroutinenya belum
	// sempet ke cancel lalu eksekusi perintah dibawahnya
	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine:", runtime.NumGoroutine())
}

// Kita bisa menambahkan sinyal cancel secara otomatis dengan menggunakan timeout
// Contoh kasus: Timeout ini bisa kita gunakan misal melakukan query untuk ambil data ke database dan ketika
// dalam waktu yang sudah kita tentukan belum mendapatkan data maka context akan auto cancel
func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())

	parent := context.Background()

	ctx, cancel := context.WithTimeout(parent, 5*time.Second) // Ini akan membuat context child dan function cancel
	// Hati hati disini, walaupun WithTimeout sudah melakukan cancel otomatis terkadang proses bisa berjalan lebih cepat
	// dari yang kita kira, karena WithTimeout melakukan timeout 5 detik ketika proses query selesai <5 detik
	// maka jika kita tidak melakukan cancel manual maka masih akan berjalan prosesnya dibackground, jadi alangkah baiknya
	// menghemat resource/menghindari goroutine leak dengan melakukan cancel manual saat proses sudah selesai, walupun
	// query ternyata sudah >5 detik dia akan tetap dicancel secara otomatis sesuai timeout yang kita atur
	defer cancel()

	destination := CreateCounter(ctx)

	fmt.Println("Total Goroutine:", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("Counter", n)
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine:", runtime.NumGoroutine())
}

// Pengaturan deadline sedikit berbeda dengan timeout, jika timeout kita beri waktu dari sekarang,
// kalo deadline ditentukan kapan waktu timeout nya, misal jam 12 siang hari ini
func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Total Goroutine:", runtime.NumGoroutine())

	parent := context.Background()

	// Kapan waktu yang tepat untuk dicancel, ini menggunakan keadaan jam yang sebenarnya dan akan
	// dieksekusi mengikuti waktu yang ada diserver
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	destination := CreateCounter(ctx)

	fmt.Println("Total Goroutine:", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("Counter", n)
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine:", runtime.NumGoroutine())
}

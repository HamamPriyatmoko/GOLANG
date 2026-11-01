package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Timer adalah representasi satu kejadian Ketika waktu timer sudah expire,
// maka event akan dikirim ke dalam channel

func TestTimer(t *testing.T) {
	// Kembaliannya adalah object Timer
	timer := time.NewTimer(10 * time.Second)
	fmt.Println(time.Now())

	fmt.Println("Waiting timer and this is still execution")

	// C adalah method Channel
	time := <-timer.C
	fmt.Println(time)
	fmt.Println("Apakah Ini Masih Dieksekusi?") // Ini akan dieksekusi setelah channel return valuenya
}

func TestTimerSecond(t *testing.T) {
	// Kembaliannya adalah object Timer
	timer := time.NewTimer(3 * time.Second)

	// Hanya akan mengembalikan channel sekali
	for time := range timer.C {
		fmt.Println("Ini Timer:", time)
	}
}

// Kadang kita hanya butuh channel nya saja, tidak membutuhkan data Timer nya
// Untuk melakukan hal itu kita bisa menggunakan function time.After(duration)
// Ini akan mengembalikan channelnya langsung bukan objectnya
func TestAfter(t *testing.T) {
	channel := time.After(4 * time.Second)
	fmt.Println("Ini akan tetap dieksekusi")

	tick := <-channel
	fmt.Println(tick)
	fmt.Println("Ini baru dieksekusi setelah channel return value")
}

// Kadang ada kebutuhan kita ingin menjalankan sebuah function dengan delay waktu tertentu
// Kita bisa memanfaatkan Timer dengan menggunakan function time.AfterFunc()
// Kita tidak perlu lagi menggunakan channel nya, cukup kirim kan function yang
// akan dipanggil ketika Timer mengirim kejadiannya
func TestAfterFunc(t *testing.T) {
	// Ini runningnya pake Goroutine
	group := sync.WaitGroup{}
	group.Add(1)
	time.AfterFunc(5*time.Second, func() {
		fmt.Println("AfterFunc Selesai Dirunning")
		fmt.Println(time.Now())
		group.Done()
	})
	fmt.Println("Setelah Run AfterFunc:", time.Now())
	group.Wait()
}

package belajar_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

// Ticker adalah representasi kejadian yang berulang
// Ketika waktu ticker sudah expire, maka event akan dikirim ke dalam channel
// Untuk membuat ticker, kita bisa menggunakan time.NewTicker(duration)
// Untuk menghentikan ticker, kita bisa menggunakan Ticker.Stop()
func TestTicker(t *testing.T) {
	ticker := time.NewTicker(3 * time.Second)

	go func(){
	  time.Sleep(5 * time.Second)
	  ticker.Stop()
	  
	}()

	// Ini akan berjalan terus menerus
	for tick := range ticker.C {
		fmt.Println("Tick selesai:", tick)
	}

	// tick := <-ticker.C
	// fmt.Println(tick)
}

func TestTick(t *testing.T) {
	channel := time.Tick(2 * time.Second)

	for time := range channel {
		fmt.Println(time)
	}
}

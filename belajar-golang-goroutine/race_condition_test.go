package belajar_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	// Ini demo test kondisi terjadinya race condition
	x := 0
	for i := 0; i < 1000; i++ {
		// Ada kemungkinan gouroutine dieksekusi secara bebarengan/paralel
		go func() {
			for j := 0; j < 100; j++ {
				x += 1
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter =", x)

	/** Hal ini bisa ditangani menggunakan Mutex, bisa diliat difile mutex_test.go dan juga atomic_test.go*/
}

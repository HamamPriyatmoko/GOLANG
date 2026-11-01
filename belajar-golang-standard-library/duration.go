package main

import (
	"fmt"
	"time"
)

func main() {
	// Artinya adalah 100 detik
	var duration1 time.Duration = 100 * time.Second
	duration2 := 10 * time.Millisecond
	duration3 := duration1 - duration2

	fmt.Println(duration3)
	fmt.Printf("duration: %d\n", duration3)
}
package main

import "fmt"

func logging() {
	fmt.Println("Function logging baru dijalankan")
}

func runApplication(){
	fmt.Println("First Run")
	fmt.Println("Second Run")
	// Menjadwalkan function untuk dieksekusi setelah function lainnya selesai dieksekusi
	// Walaupun code sebulumnya error function defer logging akan tetap dieksekusi
	defer logging() 
	fmt.Println("Third Run")
	fmt.Println("Run Application")
}

func main() {
	runApplication()
}
package main

// Import berfungsi untuk mengakses code difile/package lain
import (
	"belajar-golang-dasar/helper" // Untuk akses code/package lain pastikan diawali dengan "(nama module)/(nama package)"
	"fmt"
)

func main()()  {
	sayHello := helper.SayHello("Hamam")
	fmt.Println(sayHello)

	fmt.Println(helper.version)
	fmt.Println(helper.sayGoodBye("Hamam"))
	fmt.Println(helper.Application)
}
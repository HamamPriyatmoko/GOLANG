package main

import "fmt"

func main() {
	name := "Hamam"

	if name == "Hamam" {
		fmt.Println("Hallo Hamam")
	} else if name == "Rian" {
		fmt.Println("Hallo Rian")
	} else {
		fmt.Println("Hi, Boleh kenalan?")
	}

	// if(name == "Obiwan"){
	// 	fmt.Println("Anjay")
	// }

	if length := len(name); length > 5 { // Mendukung short statement sebelum kondisi
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
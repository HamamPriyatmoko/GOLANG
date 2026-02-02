package main

import "fmt"

func main() {
	name := "Hamam"

	// switch name {
	// case "Hamam":
	// 	fmt.Println("Hallo Hamam")
	// case "Joko":
	// 	fmt.Println("Hallo Joko")
	// default:
	// 	fmt.Println("Halo, Boleh kenalan?")
	// }

	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("Nama kamu terlalu panjang")
	// case false:
	// 	fmt.Println("Nama kamu sudah sesuai aturan")
	// }

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
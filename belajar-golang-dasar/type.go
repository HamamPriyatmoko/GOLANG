package main

import "fmt"

func main () {

	type NoKTP string // Type digunakan untuk membuat tipe data lain

	var ktpHamam NoKTP = "1111111111111"

	fmt.Println(ktpHamam)
	fmt.Println(NoKTP("22222222222")) // Langsung punya method untuk melakukan convertion
}
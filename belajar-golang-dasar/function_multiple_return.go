package main

import "fmt"

func getPerson(name string, age int) (string, int) {
	return name, age // Setelah koma adalah return keduanya
}

func main() {
	name, umur := getPerson("Hamam", 15) // Multiple return wajib valuenya ditangkap semua

	firstName, _ := getPerson("Hamam", 15)

	fmt.Println(name, umur)
	fmt.Println(firstName)
}
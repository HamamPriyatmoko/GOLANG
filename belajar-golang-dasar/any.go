package main

import "fmt"

// interface{} merupakan tipe data any, yaitu tipe data yang bisa mengembalikan tipe data apapun
func Ups() interface{} {
	// return 1
	// return true
	return "Ups"
}

// func Ups() any {
// 	return 1
// }

func main() {
	var kosong interface{} = Ups()
	fmt.Println(kosong)
}
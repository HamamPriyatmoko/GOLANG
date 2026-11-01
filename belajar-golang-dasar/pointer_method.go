package main

import "fmt"

// Didalam method defaultnya hanya menduplikasi data bukan melakukan pointer

type Man struct {
	Name string
}

// Jika seperti ini artinya parameter man adalah data hasil duplicat/copy bukan reference
// func (man Man) Married() {
// 	man.Name = "Mr. " + man.Name
// }

// Jika seperti ini artiny tipe/objectnya berbentuk pointer (Disarankan pake ini)
// artinya Married akan disimpan di Pointer Man jadi tidak akan memakan memori
func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	man := Man{"Hamam"}
	man.Married()
	fmt.Println(man.Name)
}
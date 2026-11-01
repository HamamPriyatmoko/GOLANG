package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// Membuat pointer kosong
	// var alamat1 *Address = &Address{}
	// dengan menambahkan * pada tipe data artinya adalah tipenya (pointer Address)
	// Jadi kalo *Address tipenya adalah pointer Address bukan Address
	var alamat1 *Address = new(Address) //new digunakan untuk membuat pointer
	var alamat2 *Address = alamat1

	alamat2.Province = "Yogyakarta"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}
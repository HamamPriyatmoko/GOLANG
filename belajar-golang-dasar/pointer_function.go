package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// Artinya adalah parameter ini hanya menerima tipe pointer address
// jika asterik diikuti dengan Tipe data maka disebutnya Type Literal
// artinya adalah pointer bertipe Address
func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {

	address := Address{}
	// & artinya adalah memberikan alamat memori atau disebut pointer dari variable address
	// atau bisa dibilang & (Address of)
	ChangeCountryToIndonesia(&address)

	fmt.Println(address)
}
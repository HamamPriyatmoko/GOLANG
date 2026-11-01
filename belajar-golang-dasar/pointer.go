package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// Passing by Value atau address2 hanya menyimpan duplicate value dari address1 bukan sebagai reference
	address1 := Address{"Bantul", "Yogyakarta", "Indonesia"}
	address2 := address1 // Nge-Copy value

	// *Address artinya hanya boleh diisi data alamat memori dari sebuah struct address 
	var address3 *Address = &address1 // & dan * memiliki arti yaitu pointer/memberikan data reference ke address1
	// Pastikan setelah keyword & mengacu ke variable yang ingin direference
	fmt.Println(address3)

	nama := "Hamam"
	alamatNama := &nama // Akan mengembalikan/memberikan alamat dari variable nama

	// Ini disebut Dereference yaitu aksi pergi ke alamat memori untuk mengambil atau
	// mengubah data yang ada disana 
	fmt.Println(*alamatNama) // Akan melihat isi dari sebuah alamat pointer

	address3.Province = "Jawa Barat"

	address2.City = "Bandung" // Tidak akan berdampak dengan address1

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
}
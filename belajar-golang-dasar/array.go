package main

import "fmt"

func main() {

	// 1. Deklarasi dulu baru ditambahkan datanya
	var names [3]string // Sebelum membuat array perlu ditentukan jumlah datanya

	names[0] = "Hamam"
	names[1] = "Priyat"
	names[2] = "moko"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// 2. Membuat array langsung saat dideklarasikan

	var nilai1 = [3]int{ 70, 80, 90}

	fmt.Println(nilai1)

	// 3. Adapun jika ingin membuat array tanpa ditentukan berapa panjangnya di awal dan adapula Function Array 

	var nilai = [...]int{
		80,
		90,
		78,
		66,
	}


	fmt.Println(len(nilai))
	fmt.Println(nilai[0])
}
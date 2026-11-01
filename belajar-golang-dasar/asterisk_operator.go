package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// Passing by Value atau address2 hanya menyimpan duplicate value dari address1 bukan sebagai reference
	var address1 Address = Address{"Bantul", "Yogyakarta", "Indonesia"}
	var address2 *Address = &address1 // & dan * memiliki arti yaitu pointer / memberikan reference ke address1
	// Arti dari tipe *Address adalah varible bertipe pointer Address
	address2.City = "Bandung"
	fmt.Println(address1)
	fmt.Println(address2)

	// Karena address2 adalah pointer dan kita ingin mengubah datanya ke data baru maka gunakan &(Tambahkan Tipe data)
	// Sehingga sekarang akan mengacu ke address yang lain dan baru	
	// address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	// fmt.Println(address2)

	// Akan mengubah value address2 dan juga siapapun yang mengacu ke valuenya
	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
}
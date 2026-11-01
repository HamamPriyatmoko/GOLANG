package main

import "fmt"

// type Human interface{
// 	GetName(string) string
// }

// Type Assertions adalah kemampuan merubah tipe data interface kosong ke tipe data lainnya
// artinya adalah kita bisa return data apapun terserah kita jika kita 
// deklarasikan type returnya interface{} (interface kosong)
func random() interface{} {
	return "OK"
}

func angka() int {
	return 10
}

func main()  {
	var result any = random()
	fmt.Println(result)
	var jumlah any = angka()
	// value untuk menyimpan nilainya
	// ok untuk menyimpan informasi jika gagal dilakukan konversi dengan tipe data yang dikembalikan boolean
	value, ok := jumlah.(int16) 
	fmt.Println(value, ok)

	// var resultString string = result.(string)
	// fmt.Println(resultString)

	// var resultInt int = result.(int)
	// fmt.Println(resultInt)

	// (type) artinya mengembalikan/mengkonversi nilai menjadi tipe data
	switch value := result.(type) {
	case string:
		value = result.(string)
		fmt.Println("String")
	case int:
		value = result.(int)
		fmt.Println("Integer")
	default:
		fmt.Println("Tidak diketahui", value)
	}
}
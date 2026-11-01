package main

import "fmt"

// Nil hanya bisa digunakan dibeberapa tipe data yaitu:
// interface, map, function, slice, pointer, dan channel

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	var MAP map[string]string = NewMap("Hamam")
	dataKosong := NewMap("")

	if dataKosong == nil {
		fmt.Println("Data Masih Kosong")
	} else {
		fmt.Println("Memiliki Data")
	}
	fmt.Println(MAP)
	fmt.Println(dataKosong["name"])
}
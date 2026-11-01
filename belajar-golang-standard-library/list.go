package main

import (
	"container/list"
	"fmt"
)

func main(){
	// Data kembaliannya bertipe pointer List
	var data *list.List = list.New()
	data.PushBack("Hamam")
	data.PushBack(15)
	data.PushBack("Priyatmoko")

	// Mengambil data list dari paling belakang/terakhir
	// var head *list.Element = data.Back()

	// Mengambil data list dari paling depan/pertama
	var head *list.Element = data.Front()

	// Next mengembalikan list element selanjutnya
	next := head.Next()
	fmt.Println(next.Value) // Value untuk mengambil nilainya

	// nilai "apakek" akan jadi nilai paling pertama karena menambahkannya ke paling depan
	data.PushFront("apakek")

	// Deklarasi datanya dulu, lalu cek apakah nil/tidak, lalu deklarasi lagi e-nya jadi nilai selanjutnya
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
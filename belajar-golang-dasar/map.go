package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{
	// 	"name": "Hamam",
	// 	"age":  "19",
	// }
	// person["name"] = "Hamam"
	// 	// person["age"] = "20"

	// fmt.Println(person)

	// var nilai = map[string]int{
	// 	"A":  100,
	// 	"B+": 90,
	// 	"B":  80,
	// 	"C":  70,
	// 	"D":  60,
	// }

	// fmt.Println(nilai["A"])

	// person := map[int]string{
	// 	1: "Jelek",
	// 	2: "Agak Jelek",
	// 	3: "Bagus",
	// 	4: "Bagus Banget",
	// }

	// fmt.Println(person[1])

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Hamam"
	book["ups"] = "Salah"

	fmt.Println(book)

	delete(book, "ups")
	fmt.Println(book)
}
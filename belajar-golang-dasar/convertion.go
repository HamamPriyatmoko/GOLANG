package main

import "fmt"

func main() {
	// var name = "Hamam"
	// var size string
	// size = "2XL"
	// var age string = "15"

	// var ageInt int16 = int16(age) // you cannot do this

	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	
	var nilai16 int16 = int16(nilai64)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)
	// fmt.Println(ageInt)

	var nama string  = "hamam"
	var h = nama[0] // Type byte
	var hString string = string(h)

	fmt.Println(h)
	fmt.Println(hString)
}
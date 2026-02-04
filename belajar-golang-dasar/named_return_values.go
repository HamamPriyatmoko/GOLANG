package main

import "fmt"

// Jika type data hanya dideklarasikan diakhir type return maka semua variable akan bertipe tersebut
func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Hamam"
	middleName = "Priyat"
	lastName = "moko"

	return firstName, middleName, lastName
}

func main() {

	first, second, third := getCompleteName()

	fmt.Println(first + " " + second + third)
}
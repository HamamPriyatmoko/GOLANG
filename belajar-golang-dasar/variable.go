package main

import "fmt"

func main (){
	// this is how to make a variable, and there is 2 types of how we create variable

	// 1. First with var

	var name = "Hamam" // must have a value for identified what kind of data from this variable in go
	var age int16 = 18
	var size string // and you can actually do this if you dont want to input value immedietly
	size = "2XL"
	// name = 28  // dont do this because go not allow you to change kind of data that have declare

	fmt.Println(size)
	var (
		lastName string = "Priyatmoko"
		yourName = "Aza"
	)

	fmt.Println(lastName)
	fmt.Println(yourName)
	fmt.Println(name)
	fmt.Println(age)

	// 2. Second whitout var but with :=

	// this is called immedietly declare

	word := "anjay"
	// you can do this for change it value
	word = "apakek"

	fmt.Println(word)

	// 3. Third use const for making constant variable if you won't change the value

	// const firstName string = "Hamam"
	// const middleName = "Priyatmoko"

	const (
		firstName string = "Hamam"
		middleName = "Priyatmoko"
	)

	fmt.Println(firstName)
	fmt.Println(middleName)
	
}
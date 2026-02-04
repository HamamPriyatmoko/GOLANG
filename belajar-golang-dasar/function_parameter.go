package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hallo", firstName, lastName)
}

func main() {
	sayHelloTo("Hamam", "Priyatmoko")
}
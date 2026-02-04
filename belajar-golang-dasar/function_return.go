package main

import "fmt"

func getHello(name string) string {
	return "Hello " + name
}

func main() {
	var result string = getHello("Hamam")

	fmt.Println(result)
}

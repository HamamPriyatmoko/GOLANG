package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Maaf Anda Diblacklist", name)
	} else {
		fmt.Println("Selamat Datang", name)
	}
}

func main(){
	blacklist := func(name string)bool{
		return name == "anjing"
	}
	registerUser("Hamam", blacklist)
	registerUser("anjing", blacklist)
}
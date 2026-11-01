package helper // Mengikuti nama folder
import "fmt"

// Access modifier didalam go adalah menggunakan huruf besar agar bisa diakses keluar package

var version = "1.0.0"      // Tidak akan bisa diakses diluar package karna pake huruf kecil awalanya
var Application = "golang" // Bisa diakses diluar package karna mengugnakan huruf besar awalanya

func sayGoodBye(name string) string {
	return "Good Bye " + name
}

func SayHello(name string) string {
	return "Hello " + name
}

func Contoh(){
	sayGoodBye("Hamam")
	fmt.Println(version)
}
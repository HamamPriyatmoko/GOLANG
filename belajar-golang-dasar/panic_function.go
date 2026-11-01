package main

import "fmt"

func endApp() {
	fmt.Println("End App")
	message := recover() // Recover berfungsi untuk menangkap pesan error panic, dan mengentikan proses panic
	fmt.Println("Terjadi panic", message)
}

func runApp(Error bool){
	defer endApp() // Defer akan tetap dipanggil walau program terjadi panic
	if Error {
		panic("Ups Error")
	}
	/** Ini tidak akan berjalan karena panic mengentikan program 
	didalam function tersebut dan sebetulnya mengentikan satu 
	aplikasi tetapi karena ada recover dia akan mencegahnya menghentikan seluruh program*/
	fmt.Println("Mokoy") 
	// message := recover() //Cara yang salah, recover biasanya dimasukkan kedalam function yang mau didefer
	// fmt.Println("Terjadi panic", message)
}

func main (){
	// Ini seperti Try ..catch dinodejs dengan menggunakan defer, panic, recover
	runApp(true)
	fmt.Println("Hamam Priyatmoko")
}

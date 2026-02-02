package main

// Kalo pengen tahu cara-cara membuat slice -> operation/slice.png
// Kalo pengen tahu konsep slice -> operation/konsepSlice.png

import "fmt"

func main (){

	names := [...]string {
		"Eko",
		"Hamam",
		"Joko",
		"Buddy",
		"Bagas",
		"Aji",
	}

	slice := names[4:6] // length 2, capacity 5, pointer 4
	fmt.Println(slice)

	slice2 := names[:3] // Mengambil dari index 0 sampe ke index 2
	fmt.Println(slice2)

	slice3 := names[1:]
	fmt.Println(slice3)

	var slice4 []string = names[:] // Kalo slice tidak usah ditentukan berapa datanya
	fmt.Println(slice4)


	// Function pada slice

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daySlice1 := days[5:]
	fmt.Println("Ini adalah daySlice1", daySlice1)
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru") // Membuat slice baru jika menambahkan data melebihi capacitasnya
	daySlice3 :=  append(days[3:], "Libur Baru")
	fmt.Println("Ini adalah daySlice3", daySlice3)
	// daySlice3[]
	fmt.Println("Ini adalah daySlice2", daySlice2)
	daySlice2[0] = "Ups"
	fmt.Println(daySlice2)
	fmt.Println(days)


	var newSlice []string = make([]string, 2, 5)
	// var newSlice []string = make([]string, SIZE, CAPACITY)
	newSlice[0] = "Hamam"
	newSlice[1] = "Hamam"
	// newSlice[2] = "Hamam" // error, harusnya menggunakan append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Hamam")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2)) // Mengembalikan Capacity dari Slice newSlice2 

	newSlice2[0] = "Budi" // Karena kapasitasnya 5 dan baru kepake 2 maka newSlice2 ini tidak membuat array baru
	fmt.Println(newSlice2)
	fmt.Println(newSlice) 

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice) // Meng-Copy dari slice fromSlice ke toSlice
	fmt.Println(toSlice)
	fmt.Println(fromSlice)

	iniArray := [...]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5} // Perhatikan titik tiga di array, yang membedakan adalah titik tiganya

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}

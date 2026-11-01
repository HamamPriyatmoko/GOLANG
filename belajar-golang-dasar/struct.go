package main

import "fmt"

/**
Struct adalah template data yang menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
Struct biasanya representasi data dalam program aplikasi yang kita buat
Data distruct disimpan dalam field
*/

type Customer struct {
	Name, Address string
	Age           int
}

// Ini adalah cara membuat method pada object/dari struct
// Buat paramter structnya sebelum nama function
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {
	var hamam Customer
	hamam.Name = "Hamam Priyatmoko"
	hamam.Address = "Indonesia"
	hamam.Age = 35
	// hamam.Telepone = "2242353443" // Ini akan error karena ditamplatenya tidak menerima atribute Telepone
	fmt.Println(hamam)
	fmt.Println(hamam.Name)

	// Ini adalah cara lain membuat object dari struct
	silvya := Customer{
		Name:    "Silvya Amanda",
		Address: "Indonesia",
		Age:     23,
	}
	fmt.Println(silvya)

	bambang := Customer{"Bambang Doremi", "Jakarta", 30}
	fmt.Println(bambang)

	budi := Customer{"Budi Doremi", "Belanda", 15}
	budi.sayHello("Hamam")

	var joko = struct {
		Name string
		Age  string
	}{
		Name: "Joko Anwar",
		Age:  "80",
	}
}

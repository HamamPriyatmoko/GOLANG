package main

import "fmt"

// Interface adalah kontrak jadi jika ingin mengimplementasikan ini harus memenuhi kontraknya
type HasName interface {
	GetName() string
}

func sayHello(hasName HasName) {
	fmt.Println("Hallo", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func main() {

	// Karena person sudah mendeklarasikan method GetName sebelumnya maka person sudah 
	// mengimplementasikan interface, jadi klo struct memiliki method yang memenuhi syarat kontrak dari interface
	// maka struct tersebut akan bisa disebut menjadi jenis/bagian dari interfacenya
	budi := Person{Name: "Budi Doremi"}
	sayHello(budi)
}
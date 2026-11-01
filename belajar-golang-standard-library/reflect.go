package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	// Fitur struct tag gunakan backtick, namaTag:"ValueTag" untuk value wajib gunakan double quote
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
	Email string `required:"true" max:"10"`
}

func readField(value any) {
	// reflect untuk membaca tipe structur data pada object
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println(valueType)
	
	fmt.Println("Type Name:", valueType.Name())

	// NumField untuk mengambil jumlah Field yang ada didalam struct
	for i := 0; i < valueType.NumField(); i++ {
		var valueField reflect.StructField = valueType.Field(i)
		fmt.Println("Property Name:",valueField.Name, "Type:", valueField.Type, "Nilai Tag Required:", valueField.Tag.Get("required"), "Nilai Tag Max:", valueField.Tag.Get("max"))
	}
}

// Interface kosong (interface{}) akan mengembalikan tipe data apapun sama seperti any
func IsValid(value any)(result interface{} , ok bool) {
	// Mengambil Type Objectnya/Type Structnya
	ok = true
	result = value
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		// Ambil Field Dari Type Objectnya
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			// ValueOf untuk mengambil isi dari si Objectnya bukan Type-Type Field didalam Objectnya/Type objectnya
			// Field disini mengambil field dr object by index, misal field pertama maka yang diambil adalah Name
			// Interface mengembalikan nilai dari siFieldnya
			data := reflect.ValueOf(value).Field(i).Interface()
			ok = data != "" 
			if !ok {
				result = "Atribute " + f.Name + ": " + "Tidak Boleh Kosong"
				return result, ok
			}
		}
	}
	return result, ok
}

func main() {
	readField(Sample{"Hamam"})
	readField(Person{"Hamam", "Jakarta", "hamam@gmail.com"})

	var Age int = 10

	fmt.Println("Tipe dari Variable Age:",reflect.TypeOf(Age))
	// fmt.Println(Age.NumField())

	person := Person{
		Name: "Hamam",
		Address: "indonesia",
		Email: "fads",
	}

	result, ok := IsValid(person)
	if !ok {
		fmt.Println("Error:", result)
	} else {
		fmt.Println(result.(Person))
	}

	fmt.Println(IsValid(person))
	// fmt.Println(result)
}
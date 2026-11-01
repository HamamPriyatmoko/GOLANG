package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError = errors.New("not found error")
)

func GetById(id string) error  {
	if id == ""{
		return ValidationError
	} 
	
	if id != "hamam" {
		return NotFoundError
	}

	return nil
}

func main() {
	err := GetById("hamam")

	if err != nil {
		// Fungsi errors.Is adalah untuk mencari tahu jenis tipe errornya apa
		if errors.Is(err, ValidationError){
			fmt.Println("validation error")
		} else if errors.Is(err, NotFoundError){
			fmt.Println("not found error")
		} else {
			fmt.Println("Unknown Error")
		}
	}
}
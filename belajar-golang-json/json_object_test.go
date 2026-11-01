package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

// Struct merupakan representative sebuah object
type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Height     int
	Hobbies    []string
	Addresses  []Address
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Ali",
		MiddleName: "Akbar",
		LastName:   "Hamid",
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

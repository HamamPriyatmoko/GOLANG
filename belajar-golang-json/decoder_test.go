package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamDecoder(t *testing.T) {
	// os.File merupakan jenis reader/streaming
	reader, _ := os.Open("Customer.json")
	// Menggunakan NewDecoder karena data JSON yang ingin kita decode berupa data streaming/io.Reader
	// yang dimana datanya tidak langsung semua disimpan dalam memori
	decoder := json.NewDecoder(reader)

	customer := &Customer{}
	decoder.Decode(customer)

	fmt.Println(customer)
}

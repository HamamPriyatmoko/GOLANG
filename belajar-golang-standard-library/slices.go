package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Hamam", "Silvya", "Paul"}
	values := []int{100, 20, 56, 89}

	// fmt.Println(slices.Chunk(names,10))
	// Mencari tahu apakah slice mengandung string "Paul"
	fmt.Println(slices.Contains(names, "Paul"))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(values))
	// Untuk mendapatkan index dari string yang kita cari didalam slice
	fmt.Println(slices.Index(names, "Hamam"))	
}
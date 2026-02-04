package main

import "fmt"

func main() {
	// for i := 0; i < 10; i++ {
	// 	if i%2 == 0 {
	// 		// fmt.Println("Angka Genap:", i)
	// 		continue
	// 	}

	// 	fmt.Println("Angka:", i)
	// }

	for i := 0; i < 10; i++ {
		if i == 5 {
			// fmt.Println("Angka Genap:", i)
			break
		}

		fmt.Println("Angka:", i)
	}
}
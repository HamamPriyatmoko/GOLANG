package main

import (
	"fmt"
	"math"
)

func main()  {
	// Membulatkan float keatas atau kebawah, sesuai dengan yang paling dekat
	fmt.Println(math.Round(45.5)) // Jika commanya sudah diangka 5 dia akan dibulatkan keatas
	fmt.Println(math.Floor(35.9)) // Angka berapapun dibelakang koma maka tetap dibulatkan kebawah
	fmt.Println(math.Ceil(25.4)) // Angka berapapun dibelakang koma maka tetap dibulatkan keatas
	fmt.Println(math.Max(14,23))
	fmt.Println(math.Min(14,23)) 
}
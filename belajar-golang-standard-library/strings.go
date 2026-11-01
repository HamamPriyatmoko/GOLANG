package main

import (
	"fmt"
	"strings"
)

func main()  {
	// Meminta 2 parameter, 1-string yang akan dicari, 2-string yang dicari 
	fmt.Println(strings.Contains("Hamam Priyatmoko", "H"))
	// Meminta 2 parameter, 1-string yang akan dipecah, 2-mau dipecah berdasarkan apa
	fmt.Println(strings.Split("Hamam/Priyatmoko", "/"))
	// Meminta 1 parameter, 1-string yang akan dikecilkan hurufnya
	fmt.Println(strings.ToLower("Hamam Priyatmoko"))
	// Meminta 1 parameter, 1-string yang akan dicapsclock hurufnya
	fmt.Println(strings.ToUpper("Hamam Priyatmoko"))
	// Meminta 2 parameter, 1-string yang akan ditrim, 2-trim berdasarkan apa
	fmt.Println(strings.Trim("/Hamam Priyatmoko/", "/"))
	// Meminta 3 parameter, 1-string yang akan direplace, 
	// 2-kata mana yang akan direplace, 3-direplace dgn kata apa
	fmt.Println(strings.ReplaceAll("Hamam Priyatmoko", "Hamam", "Ali"))
}
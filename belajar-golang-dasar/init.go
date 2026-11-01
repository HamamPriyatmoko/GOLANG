package main

import (
	"belajar-golang-dasar/database"
	_ "belajar-golang-dasar/internal" // Agar golang tidak komplain karna kita hanya ingin menjalankan initnya saja
	"fmt"
)

func main()  {
	fmt.Println(database.GetDatabase())
}
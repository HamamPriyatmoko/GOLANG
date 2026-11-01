package main

import (
	"fmt"
	"path"
)

func main() {
	// Untuk mendapatkan foldernya/directorinya
	fmt.Println(path.Dir("hello/world.go"))
	// Akan mengecek apakah pathnya merupakan dari jalur root(absolute path)
	fmt.Println(path.IsAbs("hello/world.go"))
	// Memotong path
	fmt.Println(path.Split("hello/world.go"))
	// Akan mengembalikan path yang paling akhir yaitu world.go
	fmt.Println(path.Base("hello/world.go"))
	// Akan mengembalikan extension dari path tujuan
	fmt.Println(path.Ext("hello/world.go"))
	// Akan menggabungkan path
	fmt.Println(path.Join("hello", "utils", "world.go"))
}
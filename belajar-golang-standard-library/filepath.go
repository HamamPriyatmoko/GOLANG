package main

import (
	"fmt"
	// filepath akan otomatis menangani penggunaan path yang 
	// berbeda beda seperti diwindos maupun linux
	"path/filepath"
)

func main(){
	// Untuk mendapatkan foldernya/directorinya
	fmt.Println(filepath.Dir("hello/world.go"))
	// Akan mengecek apakah pathnya merupakan dari jalur root(absolute path)
	fmt.Println(filepath.IsAbs("hello/world.go"))
	// Akan mengecek apakah dari root atau tidak
	fmt.Println(filepath.IsLocal("hello/world.go"))
	// Memotong path
	fmt.Println(filepath.Split("hello/world.go"))
	// Akan mengembalikan path yang paling akhir yaitu world.go
	fmt.Println(filepath.Base("hello/world.go"))
	// Akan mengembalikan extension dari path tujuan
	fmt.Println(filepath.Ext("hello/world.go"))
	// Akan menggabungkan path tp jika diwindos akan mengembalikan berupa backslash
	fmt.Println(filepath.Join("hello", "utils", "world.go"))
}
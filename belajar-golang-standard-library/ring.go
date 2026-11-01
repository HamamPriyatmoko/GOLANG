package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	// Saat Ring dibuat maka posisi data berada dielement pertama
	var data *ring.Ring = ring.New(5)
	fmt.Println(data.Len())
	// fmt.Println("Isi data", data.Value)

	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.FormatInt(int64(i), 2)
		// Jika data sudah berada dielement terakhir maka next selanjutnya akan mengembalikan element pertama
		data = data.Next()
	}

	// Ini mengambil/menyimpan nila kedalam element pertama didalam data
	// data.Value = "Value 1"

	// // Sekarang data berada dielement berikutnya yaitu elemen kedua
	// data = data.Next()
	// data.Value = "Value 2"

	// data = data.Next()
	// data.Value = "Value 3"

	// data = data.Next()
	// data.Value = "Value 4"

	// data = data.Next()
	// data.Value = "Value 5"

	// Ini melakukan iterasi otomatis kepada setiap element yang ada didalam data
	// Mirip seperti penggunaan array.map() pada javascript
	// value akan berisikan data element
	data.Do(func(value any) {
		// value = "Data 1"
		fmt.Println(value)
	})
}
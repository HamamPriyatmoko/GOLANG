package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	// Contoh reka CSV
	csvString := "muhammad,hamam,priyatmoko\n" + "budi,santoso,babon\n" + "ahmad,yani,dahlan\n"

	// Membuat reader
	reader := csv.NewReader(strings.NewReader(csvString))

	// Jika tanpa ada statement difor-nya maka itu infinity/tidak akan pernah berhenti
	for {
		// record adalah data hasil baca perbarisnya
		record, err := reader.Read()
		
		// EOF (End Of File) artinya sudah baris paling akhir difile
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}
}
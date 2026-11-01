package main

import (
	"fmt"
	"time"
)
func main()  {
	// Now akan mengembalika waktu sesuai dengan timezone laptop yang kita gunakan
	var now time.Time = time.Now()
	fmt.Println(now)
	// Waktu Local adalah waktu sesuai dengan laptop/komputer kita
	fmt.Println(now.Local())
	
	// Month bisa berupa int ataupun type Month itu sendiri seprti time.January
	var utc time.Time = time.Date(2009, 1, 15, 15, 8, 15, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	// Cara membuat format waktu diGo bukan dengan yyyy-MM-dd tapi seperti ini:
	formatter := "2006-01-02 15:04:05"

	// Contoh jika user input waktunya
	value := "2020-02-26 18:02:03"
	valuTime, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println(valuTime)
	}

	fmt.Println(valuTime.Year())
	fmt.Println(valuTime.Month())
	fmt.Println(valuTime.Day())
	fmt.Println(valuTime.Hour())
}
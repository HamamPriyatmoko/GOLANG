package main

// strconv singkatan dari string convertion, nah untuk melakukan
// konversi dari int ke string dan sebagainya gunakan "strconv"

import (
	"fmt"
	"strconv"
)

func main()  {
	var Number int = 10 
	// A/a = Ascii, i/I = int, jadi Atoi adalah Ascii(string) to Integer dan begitu juga sebaliknya
	s, err := strconv.Atoi("-14")
	fmt.Println(strconv.Itoa(Number))
	fmt.Println(s, err)

	// Konversi string ke boolean
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println("Error:", err.Error())
	}

	// Konversi string ke integer
	value, err := strconv.ParseInt("15", 10, 32)
	if err != nil{
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println(value)
	}

	// Konversi string to float
	angkaFloat, err := strconv.ParseFloat("344.50", 32)
	fmt.Println(angkaFloat)

	// Konversi string to unsigned number
	angkaUint, err := strconv.ParseUint("354464",10,32)
	fmt.Println(angkaUint)

	// Jika menggunakan function Format artinya dari tipe lain dikonversi ke string
	bolean := strconv.FormatBool(true)
	fmt.Println(bolean)

	// base 2 adalah bilangan biner
	biner := strconv.FormatInt(89, 2)
	fmt.Println(biner)

	// strconv.FormatFloat()
	// strconv.FormatUint()
}

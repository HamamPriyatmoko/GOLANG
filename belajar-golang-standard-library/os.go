package main

import (
	"fmt"
	"os" // Os untuk agar kita bisa berinteraksi dengan operating systemnya
)

func main()  {
	args := os.Args
	for _, value := range args {
		fmt.Println(value)
	}

	OS, OK := os.Hostname()
	if OK == nil {
		fmt.Println(OS, OK)
	} else {
		fmt.Printf("Ini errornya: %s", OK.Error())
	}
}
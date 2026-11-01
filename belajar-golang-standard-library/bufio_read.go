package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	input := strings.NewReader("this is long string\nwith new line\n")
	fmt.Println(input)
	// Bufio adalah sebuah package yang sudah jadi artinya siap pakai untuk read ataupun write
	reader := bufio.NewReader(input)
	for {
		// Readline untuk membaca perbaris, dan sebetulnya Readline mengimplementasikan Read
		// tp disini sudah disediakan oleh go agar kita tidak perlu menuliskan code low level atau dari dasar
		// jadi tinggal dipakai saja
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(line))
	}

	
}
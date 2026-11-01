package main

import (
	"bufio"
	"io"
	"os"
)

func createNewFile(name string, message string) error {
	// jika ingin menambahkan flag lebih pakai pipe |, lalu untuk permissionnya 
	// biasanya ditambahkan angka 0 didepannya
	file, err := os.OpenFile(name,os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	// Pakai defer agar Closenya setelah semua fungsi selesai di eksekusi
	defer file.Close()

	file.WriteString(message)
	return  nil
}

func readFile(name string) (string, error)  {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()
	
	reader := bufio.NewReader(file)
	var message string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		message += string(line) + "\n"
	}
	return message, nil
}

func addToFile(name string, message string) error {
	// Append untuk menambahkan tulisan jadi nanti bukan dioverwrite, tapi ditambahkan
	file, err := os.OpenFile(name,os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	// Pakai defer agar Closenya setelah semua fungsi selesai di eksekusi
	defer file.Close()

	// Jadi ini sekarang fungsinya akan menambahkan teks kedalam file bukan menimpa isi file dengan yang baru
	file.WriteString(message)
	return  nil
}

func main() {
	// createNewFile("cobaFileHamam.txt", "Hallo Nama Saya Hamam Priyatmoko")
	// file, err := readFile("cobaFileHamam.txt")
	// if err != nil {
	// 	err.Error()
	// }
	// fmt.Println(file)
	addToFile("cobaFileHamam.txt", "\nCoba menambahkan text baru")
}
package main

import (
	"errors"
	"flag"
	"fmt"
)

type Nilai struct{
	value string
}

func (v *Nilai) Set(message string) error {
	if message == ""{
		return errors.New("Tidak boleh kosong")
	}
	v.value = message
	return nil
}

func (v *Nilai) String() string {
	return v.value 
}

func main() {
	// Flag akan mengembalikan sebuah pointer
	host := flag.String("host", "localhost", "to inform what host should we running in")
	username := flag.String("username", "root", "your username to access") 
	port := flag.Int("port", 5173, "which port do you want to running in")
	enable := flag.Bool("enable", false, "do you want to enabled this feature?")

	// var nilai *Nilai // Jika seperti ini go akan otomatis ngasi default valuenya nil
	var nilai Nilai // Nah klo gini go akan langsung kasi defaultnya valuenya object Nilai
	flag.Var(&nilai, "alamat" , "addressing the app")
	// flag.Var butuh yang namanya parameter Value, Value adalah type interface jadi perlu parameter 
	// yang sama dengan kontrak Value nya 
	
	var database *string = flag.String("database", "postgresql", "database to use")
	// Wajib parse untuk memproses input dari argument
	flag.Parse()
	
	fmt.Println(nilai.value)
	fmt.Println(database)
	fmt.Printf("The host: %s, Username: %s, Port: %d, Feature: %t", *host, *username, *port, *enable)
}
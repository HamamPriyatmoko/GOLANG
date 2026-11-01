package main

import (
	"fmt"
	"sort"
)

// Untuk menggunakan package sort perlu implementasi kontrak dari interface sort itu sendiri
type User struct{
	Name string
	Age int
}
// UserSlice implements sort.Interface untuk type slice []User berdasarkan the Age field.
// Bisa bikin mau sort berdasarkan apapun tidak hanya age
type UserSlice []User

// Secara default slice sudah bertipe pointer
func (s UserSlice) Len() int{
	return len(s)
}

func (s UserSlice) Swap(i,j int){
	s[i], s[j] = s[j], s[i]
}

func (s UserSlice) Less(i, j int) bool  {
	return s[i].Age < s[j].Age
}

func main()  {
	user := []User{
		{"Bobon", 50},
		{"Hamam", 22},
		{"Eko", 34},
		{"Silvya", 22},
		{"Budy", 24},
		{"Jonatan", 35},
	}

	sort.Sort(UserSlice(user))

	fmt.Println(user)
}
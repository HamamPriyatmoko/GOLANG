package main

import (
	"regexp"
	"fmt"
)

func main(){
	// Membuat format regexnya
	var regex *regexp.Regexp = regexp.MustCompile(`h([a-z])m`)

	fmt.Println(regex.MatchString("ham"))
	fmt.Println(regex.MatchString("h5m"))
	fmt.Println(regex.MatchString("hfm"))

	// Akan mengembalikan slice
	fmt.Println(regex.FindAllString("ham Ham hom hrm dgd h3m", 4))
}
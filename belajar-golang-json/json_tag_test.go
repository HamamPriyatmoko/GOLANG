package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id:       "P001",
		Name:     "Apple Macbook Pro M1",
		ImageURL: "http://example.com/image.png",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))

}

// golang tidak akan peduli dengan huruf besar atau kecilnya json dia akan tau atribute mana yang dimaksud
// misal di json atribut ID pakai huruf besar, golang akan tau kalo itu adalah atribute Id, kecuali jika
// ada atribute yang pakai _ jika kita diawal udah setting nama attribute dijson itu pakai _ berarti atribute
// yang dimaksud adalah atribute yang menggunakan _ jika tidak maka golang tidak akan memahaminya, contoh:
// image_url tapi didalam json tertulisnya Imageurl nah hal ini didalam golang artinya atribute yang berbeda
// nah yang benar misal IMAGE_URL itu golang akan memahaminya sebagai atribute yang sama yang kita setting sebagai
// image_url
func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id":"P001","name":"Apple Macbook Pro M1","image_url":"http://example.com/image.png"}`
	bytes := []byte(jsonString)

	product := &Product{}
	json.Unmarshal(bytes, product)

	fmt.Println(product)
}

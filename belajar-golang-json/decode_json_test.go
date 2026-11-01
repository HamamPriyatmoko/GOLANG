package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJSON(t *testing.T) {
	JSONObject := `{"FirstName":"Hamam","LastName":"Priyatmoko","Age":15,"Height":178}`
	jsonObject := []byte(JSONObject)

	customer := &Customer{}
	// Gunakan unmarshal jika Data JSONnya sudah berupa []byte atau string (data yang sudah siap di RAM)
	json.Unmarshal(jsonObject, customer)

	fmt.Println(customer)
}

package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"Validation Error"}
	}

	if id != "hamam" {
		return &notFoundError{"data not found"}
	}

	return nil
}

func main() {
	
	err := SaveData("hamam", nil)
	fmt.Println(err)
	if err != nil {
		// Terjadi error
		// if validationErr, ok := err.(*validationError); ok {
		// 	fmt.Println(validationErr.Error())
		// } else if notFoundErr, ok := err.(*notFoundError); ok {
		// 	fmt.Println(notFoundErr.Error())
		// } else {
		// 	fmt.Println("Unknown Error:", err.Error())
		// }

		switch finalError := err.(type){
		case  *validationError:
			fmt.Println(finalError.Error())
		case *notFoundError:
			fmt.Println(finalError.Error())
		default:
			fmt.Println("Unknown Error:", finalError.Error())
		}
	} else {
		fmt.Println("Sukses")
	}
}
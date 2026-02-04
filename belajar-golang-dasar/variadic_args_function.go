package main

import "fmt"

func sumAll(numbers ...int) int {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	total := sumAll(18, 25, 90, 45, 66)
	fmt.Println(total)

	numbers := []int{25,33,15,10}
	total = sumAll(numbers...)
	fmt.Println(total)
}
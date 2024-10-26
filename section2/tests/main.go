package main

import "fmt"

func divideNumbers(x, y float64) float64 {
	var result float64

	result = x / y

	fmt.Println(result)
	return result

}

func main() {

	divideNumbers(10, 5)

}

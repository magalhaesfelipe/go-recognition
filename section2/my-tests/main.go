package main

import (
	"errors"
	"fmt"
)

func main() {

	result, myError := divideNumbers(163, 31)
	if myError != nil {
		fmt.Println(myError)
		return
	}

	fmt.Println("The result of your division is:", result)
}

func divideNumbers(x, y float32) (float32, error) {
	var result float32

	if (y == 0) {
		return result, errors.New("Cannot divide by 0")
	}

	result = x / y
	return result, nil
}


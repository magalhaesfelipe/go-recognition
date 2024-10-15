package main

import (
	"fmt"
)

// global variable
var middleName string = "Magalhaes"

func main() {
	fmt.Println("Pointer of the global middleName" ,&middleName, "value:", middleName)
	
	var x, y string = returnTwo("Lopes")
	
	fmt.Println("x:",x, "y:",y)

}

func returnTwo(lastName string) (string, string) {
	
	// variable with the same name of the global variable
	var middleName string = "Johnson"

	fmt.Println("Pointer of the middleName" ,&middleName, "value:", middleName)
	
	// 'middleName' is the middleName inside this function scope. Go gives preference to the closesest variable, if both have the same name
	return middleName, lastName
}

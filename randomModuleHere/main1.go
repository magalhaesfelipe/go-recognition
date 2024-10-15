package main

import (
	"fmt"
	"time"
	"randomModule/other"
)

func main() {
	var manga1 Manga = Manga{
		Name:             "Naruto",
		Author:           "Masashi",
		Chapters:         700,
		Released:         time.Now(),
		OriginalLanguage: "Japanese",
	}
	fmt.Println(manga1)

	printMyName("Sosuken Aizen")
	other.PrintSomeCountry("Paraguay")
}

// global variable
var MiddleName string = "Magalhaes"

func hei() {
	fmt.Println("Pointer of the global middleName", &MiddleName, "value:", MiddleName)

	var x, y string = returnTwo("Lopes")

	fmt.Println("x:", x, "y:", y)

}

func returnTwo(lastName string) (string, string) {

	// variable with the same name of the global variable
	var middleName string = "Johnson"

	fmt.Println("Pointer of the middleName", &middleName, "value:", middleName)

	// 'middleName' is the middleName inside this function scope. Go gives preference to the closesest variable, if both have the same name
	return middleName, lastName
}

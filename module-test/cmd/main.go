package main

import (
	"fmt"
	"time"
	"module-test/other"
	"module-test/mangas"

)

func main() {
	var manga1 mangas.Manga = mangas.Manga{
		Name:             "Naruto",
		Author:           "Masashi",
		Chapters:         700,
		Released:         time.Now(),
		OriginalLanguage: "Japanese",
	}
	fmt.Println(manga1)
	
	// Function from the 'mangas' package
	mangas.PrintMyManga("Naruto")
	
	// Function from the 'other' package 
	other.PrintSomeCountry("Paraguay")
	showGlobalVariable()
}

// Shows the global variable 
func showGlobalVariable() {
	fmt.Println("Pointer of the global variable 'MiddleName':", &mangas.MiddleName, "value:", mangas.MiddleName)

	var x, y string = returnTwo("Lopes")

	fmt.Println("x:", x, "y:", y)

}

// Returns two string values
func returnTwo(lastName string) (string, string) {

	// variable with the same name of the global variable
	var middleName string = "Johnson"

	fmt.Println("Pointer of the local variable 'middleName':", &middleName, "value:", middleName)

	// 'middleName' is the middleName inside this function scope. Go gives preference to the closesest variable, if both have the same name
	return middleName, lastName
}

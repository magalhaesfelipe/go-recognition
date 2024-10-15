package main

import "fmt"

// available anywhere inside the 'main' package/file
var omega float64 = .9 // package level variable

func main() {

	var name string
	name = "Felipe"
	fmt.Println("The world is yours", name)

	var firstNumber float64 = 30
	firstNumber = firstNumber + omega
	fmt.Println(firstNumber)

	z := saySomething("Part 1, ")
	fmt.Println("z variable value: ", z)
	
	str, num := third()
	
	fmt.Println("This is the string: ", str)
	fmt.Println("This is the int: ", num)

}

func saySomething(phrase1 string) float64 {
	
	var phrase2 string = "Part 2"
	var result string = phrase1 + phrase2
	fmt.Println(result)
	
	return omega
}


func third() (string, int) { // the function returns two values, of different types
	return "ffffffff", 2222222
}

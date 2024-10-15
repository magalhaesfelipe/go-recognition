package main

import "log"

func main() {

	var greeting string

	greeting = "Hello there"

	log.Println("My greeting is:", greeting)
	transform(&greeting) // pass a pointer pointing to the variable
	log.Println("My greeting is:", greeting)

	showMyPointer()
}

func transform(greeting *string) { // parameter is a reference/pointer to a value of type string
	log.Println("This is the pointer to the variable 'greeting':",greeting)
	
	newValue := "Hi there"
	*greeting = newValue
	
}

func showMyPointer() {
	h := 20
	log.Println("Pointer of 'h':", &h)

}

// Pointers change the value overtime
package goTypes

import "fmt"

// The zero value for a string is the empty string ""
// String in Go are immutable; you can reassign the value of a string variable, but you cannot change the value of the string that is assigned to it.

func RunesAndString() {
	var name = "Felipe"
	var secondName = "Julia"

	name = "Kalieu"
	secondName = secondName + " Zen"

	const firstInitial rune = 'F'
	const lastInitial float32 = 'O'

	fmt.Println(secondName, firstInitial, lastInitial)

	if name > secondName {
		fmt.Println("Name is greater than second name")
	} else {
		fmt.Println("Second name is greater than name")
	}
}

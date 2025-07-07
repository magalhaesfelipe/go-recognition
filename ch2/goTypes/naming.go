package goTypes

import "fmt"

/*
	When naming variables and constants in the package block, use more
	more descriptive names to clarify what the value represents, since the scope is wider
*/

// Go uses the case of the first letter in the name of a package-level declaration to determine if the item is accessible outside the package
var Movie string = "Harry Potter" // starts with uppercase, so it's be available outside the package

// This function name starts with Upper case, so it's available outside the package
func NamingVariablesAndConstants() {
	//variable name to never use
	_0 := 0_0
	_ρ := 20
	ã := 10
	__ := "'double quotes'"

	fmt.Println(_0)
	fmt.Println(_ρ)
	fmt.Println(ã)
	fmt.Println(__)

}

func Naming2() {

	/*
			These short names serve two purposes. The first is that they eliminate repetitive typing,
		 	keeping your code shorter. Second, they serve as a check on how complicated your code is.
	*/

	/*
		If you find it hard to keep track of
		your short-named variables, your block of
		code is likely doing too much
	*/

	f := 23.5 // 'f' for 'float'
	i := 800  // 'i' for 'integer'

	fmt.Println(f, i)
}

func PrintSomething() {
	otherFunction()
}

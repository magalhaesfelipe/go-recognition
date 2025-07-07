package goTypes

import "fmt"

// You should rarely declare variables outside of functions, in what's called the package block
// You should only declare variables in the package block that are effectively immutable
// Avoid declaring variables outside of functions because they complicate data flow analysis

// Ways to declare variables

// Using VAR
func DeclaringVariables() {

	// most verbose
	var a int = 10

	// without explicit type (type inference)
	var b = 20

	// with the zero value but keeping the type
	var c int

	// declaring multiple vars with var and the same type
	var d, e int = 30, 40

	// declare var of different types (with type inference)
	var h, j = 56.7, "Aijira"

	// declare all zero values of the same type
	var f, g int

	// Declaring multiple variable with a declaration list
	var (
		k    float64
		l        = 33
		m    int = 70
		n, o     = 10, "something"
		p, q string
	)

	fmt.Println(k, l, m, n, o, p, q)
	fmt.Println("whatever", a, b, c, d, e, f, g, h, j)
}

// Using :=
func DeclaringVariables2() {
	// ':=' is not legal outside of functions

	// Assigning a value to a variable
	x := "a"

	// Declaring multiple variables at once
	vone, vtwo := 1, 2
	vthree, vfour := "Washington", 1732

	// Assigning values to existing variables
	// It creates new a new variable 'vfive', and assigns new values to the variables 'vone' and 'vthree'
	// At least one new variable need to be on the lefthand side of the :=, any of the other variables can already exist and have new values assigned to them
	vone, vthree, vfive := 10, "George Washington", 1799

	fmt.Println(x, vone, vtwo, vthree, vfour, vfive)

	// Prefer this:
	var first byte = 40 // idiomatic

	// Over this:
	second := byte(40) // not idiomatic

	fmt.Println(first, second)

}

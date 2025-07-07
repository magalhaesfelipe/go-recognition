package goTypes

import "fmt"

const myNum int64 = 10

// This is a constant block, defining multiple constants
const (
	idKey   = "id"
	nameKey = "name"
)

const sum = 20 * 10
const whatever string = "whatever"

// Constants in Go are a way to give names to literals. There is NO way in Go to declare that a variable is immutable

// Constants can be typed or untyped. An untyped constant works exactly like a literal;
// it has no type of its own but does have a default type that is used when no other type can be inferred.

func UsingConst() {
	fmt.Println(myNum, idKey, nameKey, sum, whatever)

	a := 5
	b := 10

	// Go doesn't provide a way to specify that a value calculated at runtime is immutable.
	//const sum = a + b // This won't compile!

	fmt.Println(a, b)
}

func UsingConst2() {

	// Untyped constant
	const a = 1 // the literal is by default 'int'

	// Since the const is untyped but receives a number literal, we can assign it to different number types
	var b int = a
	var c float64 = a
	var d byte = a

	// Typed constant
	const constNumber int = 10 // Can only be assigned directly only to an int
	var y int = constNumber

	const constString string = "Name" // Can only be assigned directly only to an int
	var z string = constString

	// const x byte = constNumber // compile-time error

	fmt.Println(z, y)
	fmt.Println(b, c, d)
}

func UsingConst0() {
	const phrase = "hello"

	fmt.Println(myNum)
	fmt.Println(phrase)

	//myNum = myNum + 1 // It'll not compile
	//phrase = "Hi, how are you going"

	fmt.Println(myNum)
	fmt.Println(phrase)
	fmt.Println(whatever)
	fmt.Println(sum)
}

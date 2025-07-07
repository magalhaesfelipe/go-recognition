package goTypes

import "fmt"

func Two() {
	var a uint8 = 255
	var b byte = 255
	var c byte = 255

	var isEqual = a == b && b == c
	fmt.Println(isEqual)

	var d int32 = 20
	var e int = 50
	var f = 10 // Integer literals default is type 'int'
	var g int = 49

	fmt.Println(e == g)
	fmt.Println(d, e, f)

	var divisionResult = e / f
	fmt.Println("This is the division result: ", divisionResult)

}

func ReturnDivision(a int, b int) (int, int) {
	a /= 2
	b /= 2

	return a, b
}

func ReturnSum(x int, y int) (int, int) {
	x += 10
	y += 10

	return x, y
}

func SomeFloats() float64 {

	var x float64 = 0

	var infinity = x / 0

	fmt.Println(infinity)
	return 0

}

func MyNewFunction() {
	NamingVariablesAndConstants()
}

// This function starts with lower case, so it's available only in this package
func otherFunction() string {
	fmt.Println("Available only in this package")
	return "my string"
}

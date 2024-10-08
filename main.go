package main

import "fmt"

func main() {
	myArrays()
}

func x() {
	var v int = 20
	var w int = 30

	if v < w {
		v = v - 10
	} else {
		v = v + 10
	}

	var result int = v + w

	fmt.Println("RESULT: ", result)
}

func y() {
	var z float32 = 23.3

	for x := z; x >= 0.28; x = x - 1.0 {
		fmt.Println(x)
	}
}

func myArrays() {
	var one [5]int // Create an array with 5 values of type integers

	one[4] = 49

	fmt.Println(one)
}
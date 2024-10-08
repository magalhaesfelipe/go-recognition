package main

import "fmt"

func main() {
	mySlices()
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

func mySlices() {

	//s := make([]int, 5) // creates a slice with length 5 and capacity 5, initialized with zero values

	//t := make([]int, 5, 10) // slice with length 5 and capacity 10


// Slicing an Array
arr := [5]int{1, 80, 30, 400, 5}
slice := arr[1:4] // creates a slice from elements arr[1] to arr[3], i.e., [80, 30, 400]

fmt.Println(slice)

}
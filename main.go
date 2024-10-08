package main

import "fmt"

func main() {
	sliceThree()
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

	// t := make([]int, 5, 10) // slice with length 5 and capacity 10


// Slicing an Array
// arr := [5]int{1, 80, 30, 400, 5}
// slice := arr[1:4] // creates a slice from elements arr[1] to arr[3], i.e., [80, 30, 400]

// Using a Slice Literal
 //s := []int{1, 2, 3, 4, 5} // Creates a slice with elements 1 to 5 


g := []int{49, 7, 3}
g = append(g, 200, 10000)

sliceTwo := g[1:4]

destination := make([]int, len(sliceTwo)) // Create a slice with the same length(3)
copy(destination, sliceTwo)  // Copy elements from 'sliceTwo' to 'destination'

fmt.Println(sliceTwo)
fmt.Println(len(sliceTwo))
fmt.Println(cap(sliceTwo))


}

func sliceThree() {
	// Create a slice using a literal
	numbers := []int{1, 2, 3, 4, 5}
	
	// Get the length and capacity
	var sliceLength = len(numbers)
	var sliceCapacity = cap(numbers)

	fmt.Println("Length:", sliceLength)
	fmt.Println("Capacity:", sliceCapacity)

	// Append a new element to the slice
	numbers = append(numbers, 900)
	fmt.Println("After appending:", numbers)

	// Slice the slice
	subSlice := numbers[2:4]
	fmt.Println("Subslice:", subSlice) // Output are the values of the positions 2 and 3, before 4, so 3, 4

	// Copy slice
	copySlice := make([]int, len(numbers))
	copy(copySlice, numbers)
	fmt.Println("Copy of slice:", copySlice) 


}
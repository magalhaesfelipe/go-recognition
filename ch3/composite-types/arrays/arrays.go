package arrays

import "fmt"

func MyArrays() {

	// specify the size and type of the elements
	var c [5]int
	
	// specify the size and type and pass the values
	var d = [2]string{"Felipe", "Magalhaes"}
	
	// Specifying the values of only a few indices
	var array3 = [10]float64{0: 4.6, 4: 56.3, 5: 60, 30, 9: 2.433}
	fmt.Println(array3)
	
	
	// array with length 2
	var a = [...]int{4, 5} // '...' represents the number of elements in the array, defined by the amount of values you passed, in this case 2
	a = [2]int{40, 50}
	
	// array with length 0
	var b = [...]float64{} // Here the array length will be 0 and so you can't pass any values to it
	
	fmt.Println(a, b)
	
	
	fmt.Println(c, d)
	
	return d[:] // Return all the values inside the arrary with '[:]'
}

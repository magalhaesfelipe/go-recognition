package main

import "fmt"

func main() {
	fast()
}

// CONDITIONS
func conditions() {
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


// LOOPS
func forLoop() {
	var z float32 = 23.3

	for x := z; x >= 0.28; x = x - 1.0 {
		fmt.Println(x)
	}
}

func myArrays() {
	var first [5]int  // Create an array with 5 values of type integers
	first[0] = 49     // Add '49' to the index 0 

	var second = [3]int{355, 4433, 5000}

	fmt.Println(first)
	fmt.Println(second)
}

// SLICES
func mySlices() {
	sliceA := make([]int, 5)     // creates a slice with length 5 and capacity 5, initialized with zero values
	sliceB := make([]int, 5, 10) // slice with length 5 and capacity 10

	fmt.Println("A:", sliceA, "B:",sliceB)

	// Slicing an Array
	normalArray := [5]int{70, 80, 90, 100, 200}
	sliceOfArray := normalArray[1:4] // creates a slice with elements normalArray[1] to normalArray[3], i.e., [80, 30, 400]

	// Using a Slice Literal
	sliceLiteral := []int{1, 2, 3, 4, 5} // Creates a slice with elements 1 to 5

	fmt.Println("Slice of the array:", sliceOfArray)
	fmt.Println("Slice literal:", sliceLiteral)

	G := []int{40, 50, 60}
	G = append(G, 70, 80) // Add more two values to the slice 'g'
	fmt.Println("Slice g after appending:", G)


	sliceOfTheSlice := G[0:4] // Adds values in index 1 - 3

	destination := make([]int, len(sliceOfTheSlice)) // Create a slice with the same length(3)
	copy(destination, sliceOfTheSlice)               // Copy elements from 'sliceOfTheSlice' to 'destination'

	fmt.Println("SliceOfTheSlice:", sliceOfTheSlice, "Length:", len(sliceOfTheSlice), "Capacity:", cap(sliceOfTheSlice))

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

func mapping() {
	map1 := make(map[string]int)
	map1["Ajin"] = 83
	map1["The Horizon"] = 24
	map1["Berserk"] = 376
	fmt.Println(map1)

	map2 := map[string]int{"a": 1, "b": 2}

	someValue := map2["c"]
	fmt.Println(someValue)

	// Check if a key exists
	chapters, ok := map1["Bearserk"] // 'ok' is true or false if the specified key exists or not
	if ok {
		fmt.Println("Berserk chapters:", chapters)
	} else {
		fmt.Println("Berserk not found", chapters)
	}

	// Delete a key
	delete(map1, "The Horizon")
	fmt.Println("After deleting The Horizon:", map1)

	// Iterate over the map
	for manga, chapters := range map1 { // manga = key name | chapters = value
		fmt.Println(manga, "has", chapters, "chapters")
	}

	myMap := map[string]string{"a": "The world", "b": "is yours"}

	for key, value := range myMap {
		fmt.Println("Key:", key, "Value:", value)
	}

	// Controlling Map Size with 'make()' function
	bookPage := make(map[string]int, 200)

	// Adding elements to the map
	bookPage["Lord of the Rings"] = 1000
	bookPage["Duna"] = 700
	bookPage["Harry Potters"] = 600
	bookPage["Lord of the Flies"] = 224

	fmt.Print(bookPage)
}


func fast() {
	bookPage := make(map[string]int, 200)

	// Adding elements to the map
	bookPage["Lord of the Rings"] = 1000
	bookPage["Duna"] = 700
	bookPage["Harry Potter"] = 600
	bookPage["Lord of the Flies"] = 224
	bookPage["Crepusculo"] = 600

	fmt.Print(bookPage)
}

// MAKE FUNCTION AND LITERALS
func makeAndLiteral() {
	// <== MAKE FUNCTION ==> :

	theSliceOther := make([]int, 2) // length = 2 | capacity = 2

	// Slice with length 10 and capacity 20
	theSlice := make([]int, 10, 20)

	// Map with string keys and int values
	theMap := make(map[string]int)

	// Unbuffered channel
	theChannel := make(chan int)

	fmt.Println(theSliceOther ,theSlice, theMap, theChannel)

	// The make function does not return a pointer, it returns the initialized value directly




	// <== LITERALS ==>

	// Slice literal with three elements
	sliceLiteral := []int{1, 2, 3}
	fmt.Println("sliceLiteral length: ",len(sliceLiteral), "sliceLiteral capacity:", cap(sliceLiteral))

	// Map literal with two key-value pairs
	mapLiteral := map[string]int{"first": 10, "third": 30}

	type Person struct {
		Name     string
		LastName string
		Age      int
	}
	user1 := Person{Name: "Felipe", LastName: "Magalhaes", Age: 22}

	fmt.Println("Literals values:", sliceLiteral, mapLiteral, user1)

}


// STRUCTS
func structures() {
	type Manga struct {
		Name string
		Author string
		Chapters int
		Cover string
	}
 
	naruto := Manga{
		Name: "Naruto",
		Author: "Masashi Kishimoto",
		Chapters: 700,
		Cover: "https//somesite/someimage.jpg",
	}

	fmt.Println(naruto)
}

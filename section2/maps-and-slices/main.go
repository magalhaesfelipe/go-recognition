package main

import (
	"fmt"
	"sort"
)

type Manga struct {
	Name     string
	Chapters int
}

func main() {
	//printMyMap()
	//otherMap()
	mySlices()

}

// MAPS
func printMyMap() {

	map1 := make(map[string]int)

	map1["cat1"] = 20
	fmt.Println(map1["cat1"])

	// Map with our own type for the value
	mangaMap := make(map[string]Manga)

	naruto := Manga{
		Name:     "Naruto",
		Chapters: 7000,
	}

	mangaMap["manga1"] = naruto

	fmt.Println(mangaMap["manga1"])
}

func otherMap() {
	// If you don't know which type the value of a map will be, you can use:
	randomMap := make(map[string]interface{})
	randomMap["whatever"] = 20

	fmt.Println(randomMap)
}

// SLICES
func mySlices() {
	var languages []string

	// using append to add values to the slice
	languages = append(languages, "Go")
	languages = append(languages, "Rust")
	languages = append(languages, "C++")
	languages = append(languages, "C#")

	sort.Strings(languages)
	fmt.Println(languages)

	// using the shorthand to declare a slice of strings, and the pre-populate it with some letters
	frameworks := []string{"b", "a", "c", "v", "e", "g", "f"}

	numbers := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	// using the function of the 'sort' package to sort ints(ascending)
	sort.Ints(numbers)

	// using the function of the 'sort' package to sort strings(alphabetically)
	sort.Strings(frameworks)

	fmt.Println(frameworks, numbers)

	// Selects the values from the index 3 until before the index 5
	fmt.Println(numbers[3:5])


	// fmt.Println(numbers[5:2])
}

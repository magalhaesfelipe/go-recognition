package main

import (
	"fmt"
	"sort"
)


type Manga struct {
	Name string
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
		Name: "Naruto",
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
	var languages []int

	languages = append(languages, 9)
	languages = append(languages, 8)
	languages = append(languages, 7)
	languages = append(languages, 90)
	



	sort.Ints(languages)

	fmt.Println(languages)

	
	frameworks := []string{"a", "b", "c", "d", "e", "g", "f"}
	sort.Strings(frameworks)

	fmt.Println(frameworks)

}







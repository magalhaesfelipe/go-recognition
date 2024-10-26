// JSON = Javascript Object Notation
package main

import (
	"encoding/json"
	"fmt"
)

type Manga struct {
	// Preparing this type/struct to receive json data when instantiated
	Name     string `json:"name"`
	Author   string `json:"author"`
	Chapters int    `json:"chapters"`
}

func main() {

	// json data
	mangaData := `
[
	{
		"name": "Ajin",
		"author": "Unknow",
		"chapters": 83
	},
	{
		"name": "One Punch Man",
		"author": "One, Yusuke Murata",
		"chapters": 200
	}	
]`

	// READ JSON AND CONVERT INTO A STRUCT/TYPE

	// Define a variable of type Manga
	var z []Manga

	// Unmarshing the json data, and passing it to the variable z
	// 'Unmarshal' needs to receive the json converted to bytes, and a pointer to the variable where the result will be stored
	err := json.Unmarshal([]byte(mangaData), &z)

	// If there is a error caught, 'err' it's not null(nil), so print the error
	if err != nil {
		fmt.Println("Error unmarshalling json", err)
	}

	fmt.Println("Unmarshalled:", z)

	// CONVERTING DATA FROM STRUCT TO JSON

	// Create a slice
	var sliceOfMangas []Manga

	// Create variables of type Manga to then append into the slice
	var manga1 Manga
	manga1.Name = "Berserk"
	manga1.Author = "Kentaru Miura"
	manga1.Chapters = 489

	var manga2 Manga
	manga2.Name = "Dragon Ball"
	manga2.Author = "Akira Toriyama"
	manga2.Chapters = 297

	// Keep the previous values of the slice and add the other two we passed
	sliceOfMangas = append(sliceOfMangas, manga1, manga2)

	fmt.Println(sliceOfMangas)

	// covert the 'sliceOfMangas' variable to JSON

	// '.MarshalIndent' returns a slice of bytes
	myJsonVariable, err := json.MarshalIndent(sliceOfMangas, "", "    ")

	if err != nil {
		fmt.Println("error marshalling:", err)
	}

	// Using the 'string()' method to convert the slice of bytes to an indented JSON
	fmt.Println(string(myJsonVariable))

}

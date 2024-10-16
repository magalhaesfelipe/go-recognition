package mangas

import (
	"fmt"
	"time"
)

// When we define a type, variable or function with a capital letter, it's visible outside the package
type Manga struct {
	Name             string
	Author           string
	Chapters         int
	Released         time.Time
	OriginalLanguage string
}

// global variable
var MiddleName string = "Shippuden"

func x() {
	manga1 := Manga{
		Name:             "Naruto",
		Author:           "Masashi Kishimoto",
		Chapters:         700,
		OriginalLanguage: "Japanese",
	}

	fmt.Println(manga1.Released)
}

// Function available outside this package
func PrintMyManga(name string) {
	fmt.Println("My printed manga is:", name)
}

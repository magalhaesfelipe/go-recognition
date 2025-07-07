package goTypes

import "fmt"

func ShowStrings() {
	fmt.Printf("Hello, %s!!", "Felipe")

	var anotherString string = "Greetings and \"Salutions\""
	fmt.Println(anotherString)

	var food1 string = "carrot"
	var food2 string = "tomato"
	var food3 string = "potato"

	var likeCarrrots bool // The zero value of a bool type is 'false'
	var likeTomatoes bool = true
	likePotatoes := false

	fmt.Printf("Do I like %s?: %t. ", food1, likeCarrrots)
	fmt.Printf("Do I like %s?: %t. ", food2, likeTomatoes)
	fmt.Printf("Do I like %s?: %t. ", food3, likePotatoes)
}
package main

import (
	"fmt"
	"project1/my-utils"
)

func main() {
	deckard := project1.Robot{
		FabricNumber: "FAB0499",
		StillProducing: true,
	}


	fmt.Println(deckard)
	project1.Sum(1, 2)
}
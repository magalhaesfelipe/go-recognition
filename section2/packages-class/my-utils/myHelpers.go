package project1

import "fmt"

type Robot struct {
	FabricNumber string
	StillProducing bool
}

func Sum(x int, y int) {
	var result = x + y

	fmt.Println("This is the sum of ", x, "and", y, ":", result)
}

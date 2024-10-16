package main

import "log"

func main() {
	g := 200


	// switch in Go breaks out as soon as it matches one case, it doesn't check all the other cases/conditions
	switch {
	case g > 1000:
		log.Println("G is greater than 1000")

	case g < 300:
		log.Println("G is less than 300")

	case g < 600:
		log.Println("G is less than 600")

	default:
		log.Println("G is less than 1000 and greater than 600")

	}

}

package main

import (
	"channels-module/helpers"
	"fmt"
)

const numberPool = 3000
func calculateSomeValue(intChannel chan int) {
	randomNumber := helpers.GenerateRamdomNumber(numberPool)

	// Pass the random number to the channel
	intChannel <- randomNumber
}

func main() {
	intChannel := make(chan int)

	// Whatever comes after the keyword 'defer' will be executed as soon as the current function is done
	defer close(intChannel) 

	go calculateSomeValue(intChannel)

	// Getting a value from the channel
	num := <-intChannel
	fmt.Println(num)
}

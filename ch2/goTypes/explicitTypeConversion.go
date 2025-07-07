package goTypes

import (
	"fmt"
	"strconv"
)

// Types need to be Explicit converted in Go

func MyConversion() {
	var m int = 35
	var n float64 = 50.5
	var s string = "abcdefg "

	var sum1 int = int(n) + m
	var sum2 float64 = float64(m) + n

	s += strconv.FormatFloat(n, 'f', -1, 64)

	fmt.Println(s)
	fmt.Println("This is the sum1 result: ", sum1)
	fmt.Println("This is the sum2 result: ", sum2)

}

func MyConversion2() {
	var h int = 30
	var j byte = 30

	var sum3 = h + int(j)
	var sum4 = j + byte(h)

	fmt.Println("This is the sum3 result: ", sum3)
	fmt.Println("This is the sum4 result: ", sum4)

	var z float64 = 66
	var w float64 = 300.5 * 5

	fmt.Println(z + w)

	var thousand byte = 255

	fmt.Println(thousand)
}

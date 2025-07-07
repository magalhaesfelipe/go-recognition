package exercises

import "fmt"

func Ex1() {
	i := 20
	var f float64 = float64(i)

	fmt.Printf("%d %.1f", i, f)
}

func Ex2() {
	const value = 50
	i := value
	var f float64 = value

	fmt.Printf("%d \n", i)
	fmt.Printf("%.2f", f)

}

func Ex3() {
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615

	b += 1
	smallI += 1
	bigI += 1

	fmt.Println(b, smallI, bigI)
}

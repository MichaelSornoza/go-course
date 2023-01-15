package main

import (
	"fmt"
)

func main() {

	x := 10
	y := 50

	result := x + y

	fmt.Println(result)

	result = y - x

	fmt.Println(result)

	result = x * y

	fmt.Println(result)

	result = y / x

	fmt.Println(result)

	result = y % x

	fmt.Println(result)

	x++
	fmt.Println(x)

	x--
	fmt.Println(x)
}

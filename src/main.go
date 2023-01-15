package main

import "fmt"

func main() {
	// Constants

	const pi float64 = 3.14
	const pi2 = 3.14

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	// Int variables

	base := 12
	var height int = 14
	var area int

	area = base * height

	fmt.Println("Area:", area)

	// Zero Values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Area de un cuadrado

	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("Area del cuadrado:", areaCuadrado)
}

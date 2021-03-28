package main

import "fmt"

func main() {
	//Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.15

	fmt.Println("pi: ", pi)
	fmt.Println("pi2: ", pi2)

	//Declaración de variables enteras

	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	//zero values

	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Calcular el area de un cuadrado

	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado: ", areaCuadrado)

	x := 10
	y := 50

	//suma
	result := x + y
	fmt.Println("Suma: ", result)
	//Resta
	result = y - x
	fmt.Println("Resta: ", result)
	//Multiplicacion
	result = x * y
	fmt.Println("Multiplicacion: ", result)
	//División
	result = y / x
	fmt.Println("Division: ", result)
	//Modulo
	result = y % x
	fmt.Println("Modulo: ", result)
	//Incremental
	x++
	fmt.Println("Incremental: ", x)
	//Decremental1
	x--
	fmt.Println("Decremental: ", x)

}

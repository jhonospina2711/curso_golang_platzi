package main

import (
	"fmt"
	"math"
)

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

// función area de un circulo

func areaCirculo(radioCirculo int) float64 {
	areaCirculo := math.Pi * math.Pow(float64(radioCirculo), float64(2))
	return areaCirculo
}

func main() {
	normalFunction("Hola Mundo")
	tripleArgument(1, 2, "hola")
	value := returnValue(2)
	fmt.Println("Value: ", value)

	value1, _ := doubleReturn(2)
	fmt.Println("Value1: ", value1)

	var radioCirculo int
	fmt.Println("Seleccione el valor del radio del circulo: ")
	fmt.Scanf("%d\n", &radioCirculo)
	areaCircunferencia := areaCirculo(radioCirculo)
	fmt.Printf("El Area del circulo es: %.2f", areaCircunferencia)

}

package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) {

	//var LowerText string
	var textReverse string
	LowerText := strings.ToLower(text)

	for i := len(LowerText) - 1; i >= 0; i-- {
		textReverse += string(LowerText[i])

	}

	if LowerText == textReverse {
		fmt.Println("Es un palindromo")
	} else {
		fmt.Println("No es Palindromo")
	}
}

func main() {

	slice := []string{"hola", "que", "hace"}

	fmt.Printf("********** Slice devuelve indices y valor ********\n")
	for i, valor := range slice {
		fmt.Println(i, valor)
	}

	// Si solo te interesa saber el valor debes de realizar lo siguiente
	fmt.Printf("********** Slice devuelve Valores ********\n")
	for _, valor := range slice {
		fmt.Println(valor)
	}

	// En el caso en que solo queramos conocer el indice
	fmt.Printf("********** Slice devuelve indices ********\n")
	for i := range slice {
		fmt.Println(i)
	}

	// Saber si una palabra es palindromo
	fmt.Printf("********** Es un palindromo ********\n")

	/*
			var valor string
		fmt.Println("Ingrese la palabra que desea validar si es palindromo: ")
		fmt.Scanf("%v\n", &valor)
		fmt.Print("La palabra ingresada: ")
		isPalindromo(valor)
	*/

	i := 1
	var valor string
	for i > 0 {
		fmt.Printf("********** Es un palindromo ********\n")
		fmt.Println("Ingrese la palabra que desea validar: ")
		fmt.Scanf("%v\n", &valor)
		fmt.Print("La palabra ingresada: ")
		isPalindromo(valor)
		fmt.Printf("\n")
		i++
	}

}

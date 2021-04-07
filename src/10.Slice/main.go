package main

import "fmt"

func isPalindromo(text string) {

	var textReverse string

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])

	}

	if text == textReverse {
		fmt.Println("Es un palindromo")
	} else {
		fmt.Println("No es Palindromo")
	}
}

func main() {

	/*
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
	*/

	var valor string
	fmt.Println("Ingrese la palabra que desea validar si es palindromo: ")
	fmt.Scanf("%d\n", &valor)
	fmt.Print("La palabra ingresada: ")
	isPalindromo(valor)

}

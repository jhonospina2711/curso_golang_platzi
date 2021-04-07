package main

import "fmt"

func main() {

	helloMessage := "Hello"
	worldMessage := "Platzi"

	//print

	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene mas de %d cursos\n", nombre, cursos)

	//sprintf

	message := fmt.Sprintf("%s tiene mas de %d cursos\n", nombre, cursos)
	println(message)

	//Tipos de datos
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)

}

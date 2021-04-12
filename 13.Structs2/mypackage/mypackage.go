package mypackage

import "fmt"

//CarPublic Car con acceso publico
type CarPublic struct {
	Brand string
	Year  int
}

// Se define la función carPrivate como privada
type carPrivate struct {
	brand string
	year  int
}

//Se define PrintMesage como publica Imprimir un mensaje
func PrintMessage(text string) {
	fmt.Println(text)

}

//Se define la función printMessagel en privado.
func printMessagel(text string) {
	fmt.Println(text)

}

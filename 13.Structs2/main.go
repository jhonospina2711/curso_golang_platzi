package main

import (
	"fmt"

	pk "github.com/jhonospina2711/curso_golang_platzi/13.Structs2/mypackage"
)

func main() {

	var myCar pk.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2020
	fmt.Println(myCar)

	//fmt.Println(myOtherCar)
	//La función "myOtherCar" arroja error, ya que el carPrivate es de tipo privado
	//var myOtherCar pk.carPrivate

	pk.PrintMessage("Hola Platzi")

	//Este llamada a la función printMessagel genera error
	//porque no existe la función se declaro como publica es decir
	//La primera letra en mayuscula.
	//pk.printMessagel("Hola")

}

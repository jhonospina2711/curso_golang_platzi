package main

import "fmt"

func main() {

	// Defer -> Se utiliza cuando se abre una conexión a base de datos, para que al final
	//Se cierre la conexión.
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	//Continue y break

	for i := 0; i < 10; i++ {
		fmt.Println(i)

		//Continue

		if i == 2 {
			fmt.Println("Es 2")
			continue
		}
		// Break

		if i == 8 {
			fmt.Println("Break")
			break
		}

	}

}

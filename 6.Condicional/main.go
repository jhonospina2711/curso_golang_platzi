package main

import (
	"fmt"
)

func esPar(numero int) string {

	if numero%2 == 0 {
		return "Es par"
	} else {
		return "No es par"
	}

}

func validarUsuario(usuario, contraseña string) string {
	usuarioReg := "platzi"
	passReg := "platzi"
	if usuario == usuarioReg && contraseña == passReg {
		return "Info!... Usuario y Contraseña Valido"

	} else {
		return "ERROR!... El usuario o contraseña No son invalidos"
	}

}

func main() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	//With and

	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es verdad, AND")
	} else {
		fmt.Println("Es Falso, AND")
	}

	//With Or

	if valor1 == 0 || valor2 == 3 {
		fmt.Println("Es verdad, OR")
	} else {
		fmt.Println("Es Falso, OR")
	}

	fmt.Printf("********** Reto es Par ********\n")
	var numero int
	fmt.Println("Ingrese el número que desea validar si es par: ")
	fmt.Scanf("%d\n", &numero)
	validaPar := esPar(numero)
	fmt.Println("El número ingresado: ", validaPar)

	fmt.Printf("********** Reto Autenticación ********\n")
	var user string
	var password string

	fmt.Print("Ingrese el usuario: ")
	fmt.Scanf("%v\n", &user)
	fmt.Print("Ingrese el password: ")
	fmt.Scanf("%v\n", &password)
	validaUsuario := validarUsuario(user, password)
	fmt.Println("Info: ", validaUsuario)

}

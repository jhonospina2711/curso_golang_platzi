package main

import "fmt"

func main() {
	//Array
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))

	//Slice

	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// Métodos en el slice
	//Que inprima los elementos del indice cero
	fmt.Println(slice[0])
	//Que inprima los elementos hasta el indice 3
	fmt.Println(slice[:3])
	//Que inprima los elementos del indice 2 y los elementos del indice 4
	fmt.Println(slice[2:4])
	//Que inprima del elemento 4 en adelante
	fmt.Println(slice[4:])

	//Append
	slice = append(slice, 7)
	fmt.Println(slice)

	//Append nueva list
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

}

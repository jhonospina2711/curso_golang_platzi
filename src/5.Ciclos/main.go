package main

import "fmt"

func main() {
	fmt.Printf("********** For condicional Incremento ********\n")
	//For condicional
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Printf("********** For condicional decremento ********\n")

	//For condicional decremento

	for j := 10; j >= 0; j-- {
		fmt.Println(j)
	}

	fmt.Printf("********** For While ********\n")

	//For while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	fmt.Printf("\n")

	//For forever

	/*
		counterForever := 0
		for {
			fmt.Println(counterForever)
			counterForever++
		}
	*/
}

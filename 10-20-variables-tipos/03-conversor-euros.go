package main

import "fmt"

const dollar float64 = 1.14

var euros float64

func main() {

	fmt.Print("Ingrese cantida de euros: ")
	fmt.Scanln(&euros)

	fmt.Println(euros, "€ son", euros*dollar, "$.")

}

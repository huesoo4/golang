package main

import "fmt"

var nombre string = "Jose Manuel"
var edad int = 25
var altura float64 = 1.70
var estudiando bool

func main() {
	estudiando = true //Le asignamos el valor al booleano dentro la función
	fmt.Println("Nombre:", nombre)
	fmt.Println("Edad:", edad)
	fmt.Println("Altura:", altura)
	if estudiando == true {
		fmt.Println("Estudiando: Sí")
	} else {
		fmt.Println("Estudiando: No")
	}
}

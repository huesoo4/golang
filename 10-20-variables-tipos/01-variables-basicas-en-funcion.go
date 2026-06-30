package main

import "fmt"

func main() {
	nombre := "Jose Manuel"
	edad := 25
	altura := 1.70
	estudiando := true

	fmt.Println("Nombre:", nombre)
	fmt.Println("Edad:", edad)
	fmt.Println("Altura:", altura)
	if estudiando {
		fmt.Println("Estudiando: Sí")
	} else {
		fmt.Println("Estuandianto: No")
	}
}

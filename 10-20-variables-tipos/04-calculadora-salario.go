// Para realizar este ejercicio ver las explicaciones en la carpeta inputs-operaciones

package main

import "fmt"

const porcentaje float64 = 15

var sueldo float64
var neto float64
var total float64

func main() {

	fmt.Print("Introduce el sueldo: ")
	fmt.Scanln(&sueldo)

	neto = sueldo * porcentaje / 100

	total = sueldo - neto

	fmt.Println("El total de sueldo neto es:", total)
}

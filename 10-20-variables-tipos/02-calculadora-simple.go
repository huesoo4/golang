package main

import "fmt"

var num1 int = 15
var num2 int = 2

func main() {
	fmt.Println("CALCULADORA BÁSICA")
	fmt.Println("Numero1 =", num1, "|| Numero2 =", num2)
	fmt.Println("Suma:", num1+num2)
	fmt.Println("Resta:", num1-num2)
	fmt.Println("Multiplicación:", num1*num2)
	fmt.Println("División:", num1/num2) /* Esta división nos devuelve un número entero porque se está haciendo
	la operación con dos números enteros */

	decimal1 := float64(num1) //Así convertimos enteros a decimal
	decimal2 := float64(num2)
	fmt.Println("División con decimales:", decimal1/decimal2)
}

// Por último también el operador % devuelve el resto de la división

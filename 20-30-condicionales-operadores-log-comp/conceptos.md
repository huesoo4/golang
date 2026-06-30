# Conceptos principales: Condicionales y operadores en Go

En Go, las estructuras condicionales permiten que un programa tome decisiones según se cumpla o no una condición.

Estos conceptos son fundamentales para controlar el flujo de ejecución de un programa.

---

## `if`

La palabra clave `if` se usa para ejecutar un bloque de código solo si una condición se cumple.

```go
package main

import "fmt"

func main() {
	edad := 25

	if edad >= 18 {
		fmt.Println("Eres mayor de edad")
	}
}
```

---

## `else`

La palabra clave `else` se usa para indicar qué debe ocurrir cuando la condición del `if` no se cumple.

```go
package main

import "fmt"

func main() {
	edad := 16

	if edad >= 18 {
		fmt.Println("Eres mayor de edad")
	} else {
		fmt.Println("Eres menor de edad")
	}
}
```

---

## `else if`

La estructura `else if` se usa cuando queremos comprobar varias condiciones diferentes.

```go
package main

import "fmt"

func main() {
	nota := 7

	if nota >= 9 {
		fmt.Println("Sobresaliente")
	} else if nota >= 7 {
		fmt.Println("Notable")
	} else if nota >= 5 {
		fmt.Println("Aprobado")
	} else {
		fmt.Println("Suspenso")
	}
}
```

---

## `switch`

La estructura `switch` se usa para evaluar varios casos posibles de una forma más limpia que muchos `else if`.

```go
package main

import "fmt"

func main() {
	dia := "lunes"

	switch dia {
	case "lunes":
		fmt.Println("Inicio de semana")
	case "viernes":
		fmt.Println("Casi fin de semana")
	case "sábado", "domingo":
		fmt.Println("Fin de semana")
	default:
		fmt.Println("Día normal")
	}
}
```

---

## Operadores de comparación

Los operadores de comparación sirven para comparar valores.

| Operador | Significado       | Ejemplo      |
| -------- | ----------------- | ------------ |
| `==`     | Igual que         | `edad == 18` |
| `!=`     | Distinto de       | `edad != 18` |
| `>`      | Mayor que         | `edad > 18`  |
| `<`      | Menor que         | `edad < 18`  |
| `>=`     | Mayor o igual que | `edad >= 18` |
| `<=`     | Menor o igual que | `edad <= 18` |

---

## Operadores lógicos

Los operadores lógicos sirven para combinar varias condiciones.

| Operador | Nombre | Significado                           |    |                                       |
| -------- | ------ | ------------------------------------- | -- | ------------------------------------- |
| `&&`     | AND    | Todas las condiciones deben cumplirse |    |                                       |
| `        |        | `                                     | OR | Al menos una condición debe cumplirse |
| `!`      | NOT    | Niega una condición                   |    |                                       |

---

# Idea clave

Los condicionales permiten que un programa tome decisiones.

En Go, las estructuras más utilizadas para tomar decisiones son:

```go
if
else
else if
switch
```

Y para crear condiciones usamos operadores de comparación y operadores lógicos.
::: 

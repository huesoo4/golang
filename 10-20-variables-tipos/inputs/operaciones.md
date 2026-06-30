# Pedir datos por teclado y operaciones matemáticas en Go

## 1. Pedir datos por teclado en Go

En Go podemos pedir datos al usuario usando el paquete `fmt`.

Para leer datos desde la terminal se puede usar:

```go
fmt.Scan()
```

o también:

```go
fmt.Scanln()
```

Ambos permiten guardar en una variable lo que el usuario escriba por teclado.

---

## 2. Ejemplo básico con `fmt.Scanln`

```go
package main

import "fmt"

func main() {
	var nombre string

	fmt.Print("Introduce tu nombre: ")
	fmt.Scanln(&nombre)

	fmt.Println("Hola,", nombre)
}
```

### Explicación

```go
var nombre string
```

Creamos una variable llamada `nombre` de tipo `string`.

```go
fmt.Print("Introduce tu nombre: ")
```

Mostramos un mensaje al usuario.

```go
fmt.Scanln(&nombre)
```

Leemos el dato que escribe el usuario y lo guardamos en la variable `nombre`.

El símbolo `&` significa que le pasamos a Go la dirección de memoria de la variable para que pueda modificar su valor.

```go
fmt.Println("Hola,", nombre)
```

Mostramos el resultado por pantalla.

---

## 3. Pedir números por teclado

También podemos pedir números enteros usando variables de tipo `int`.

```go
package main

import "fmt"

func main() {
	var edad int

	fmt.Print("Introduce tu edad: ")
	fmt.Scanln(&edad)

	fmt.Println("Tu edad es:", edad)
}
```

En este caso, si el usuario introduce `25`, ese valor se guarda dentro de la variable `edad`.

---

## 4. Pedir números decimales

Para números con decimales usamos `float64`.

```go
package main

import "fmt"

func main() {
	var altura float64

	fmt.Print("Introduce tu altura: ")
	fmt.Scanln(&altura)

	fmt.Println("Tu altura es:", altura)
}
```

Ejemplo de entrada:

```text
1.70
```

---

# Operaciones matemáticas en Go

Go permite hacer operaciones matemáticas básicas usando operadores.

## 1. Suma

```go
resultado := 10 + 5
```

Resultado:

```text
15
```

---

## 2. Resta

```go
resultado := 10 - 5
```

Resultado:

```text
5
```

---

## 3. Multiplicación

```go
resultado := 10 * 5
```

Resultado:

```text
50
```

---

## 4. División

```go
resultado := 10 / 5
```

Resultado:

```text
2
```

---

## 5. Módulo o resto

El operador `%` devuelve el resto de una división.

```go
resultado := 10 % 3
```

Resultado:

```text
1
```

Porque `10 / 3` da `3` y sobra `1`.


---

# Cuidado con la división entre enteros

En Go, si divides dos números `int`, el resultado también será un `int`.

Ejemplo:

```go
resultado := 5 / 2
fmt.Println(resultado)
```

Resultado:

```text
2
```

No devuelve `2.5`, porque ambos números son enteros.

Para obtener decimales, debemos usar `float64`.

```go
package main

import "fmt"

func main() {
	var numero1 float64
	var numero2 float64

	fmt.Print("Introduce el primer número: ")
	fmt.Scanln(&numero1)

	fmt.Print("Introduce el segundo número: ")
	fmt.Scanln(&numero2)

	division := numero1 / numero2

	fmt.Println("Resultado:", division)
}
```

Si introducimos:

```text
5
2
```

El resultado será:

```text
2.5
```

---

# Conversión de tipos

En Go no se pueden mezclar directamente tipos diferentes en una operación.

Ejemplo incorrecto:

```go
var edad int = 25
var altura float64 = 1.70

resultado := edad + altura
```

Esto da error porque `edad` es `int` y `altura` es `float64`.

Para solucionarlo, hay que convertir el tipo:

```go
resultado := float64(edad) + altura
```

Ahora sí funciona, porque ambos valores son `float64`.

---

# Ejemplo práctico: calcular salario neto

```go
package main

import "fmt"

func main() {
	var salarioBruto float64
	var retencion float64

	fmt.Print("Introduce el salario bruto: ")
	fmt.Scanln(&salarioBruto)

	fmt.Print("Introduce la retención en porcentaje: ")
	fmt.Scanln(&retencion)

	descuento := salarioBruto * retencion / 100
	salarioNeto := salarioBruto - descuento

	fmt.Println("Salario bruto:", salarioBruto)
	fmt.Println("Retención:", retencion, "%")
	fmt.Println("Descuento:", descuento)
	fmt.Println("Salario neto:", salarioNeto)
}
```

---

## Explicación

Si el usuario introduce:

```text
2000
15
```

El programa calcula:

```go
descuento := 2000 * 15 / 100
```

Resultado:

```text
300
```

Después calcula:

```go
salarioNeto := 2000 - 300
```

Resultado:

```text
1700
```

---

# Resumen rápido

| Concepto         | Explicación                                   |
| ---------------- | --------------------------------------------- |
| `fmt.Scanln()`   | Lee datos escritos por teclado                |
| `&variable`      | Permite guardar el dato dentro de la variable |
| `int`            | Números enteros                               |
| `float64`        | Números decimales                             |
| `+`              | Suma                                          |
| `-`              | Resta                                         |
| `*`              | Multiplicación                                |
| `/`              | División                                      |
| `%`              | Resto de una división                         |
| `float64(valor)` | Convierte un valor a decimal                  |

---

# Idea clave

Para pedir datos al usuario en Go, primero declaramos una variable y después usamos `fmt.Scanln(&variable)`.

Para hacer operaciones matemáticas, usamos operadores como `+`, `-`, `*`, `/` y `%`.

Si queremos trabajar con decimales, debemos usar `float64`.

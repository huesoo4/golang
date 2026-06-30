# Variables, constantes y tipos básicos en Go

## `var`

La palabra clave `var` se usa para declarar variables en Go.

Una variable es un espacio donde guardamos un valor que puede cambiar durante la ejecución del programa.

```go
var edad int = 25
var nombre string = "Jose Manuel"
```

También se puede declarar una variable sin asignarle valor inicial:

```go
var ciudad string
```

En ese caso, Go le asigna un valor por defecto.

---

## `:=`

El operador `:=` se usa para declarar una variable de forma corta.

Go detecta automáticamente el tipo de dato según el valor asignado.

```go
edad := 25
nombre := "Jose Manuel"
```

Esto es equivalente a:

```go
var edad int = 25
var nombre string = "Jose Manuel"
```

Importante: `:=` solo puede usarse dentro de funciones.

---

## `const`

La palabra clave `const` se usa para declarar constantes.

Una constante es un valor que no puede cambiar durante la ejecución del programa.

```go
const pais = "España"
const iva = 21
```

Si intentamos modificar una constante, Go dará error.

```go
pais = "Francia" // Error
```

---

## `int`

El tipo `int` se usa para representar números enteros.

Es decir, números sin decimales.

```go
edad := 25
usuarios := 100
```

Ejemplos de valores `int`:

```go
10
25
100
-5
```

---

## `float64`

El tipo `float64` se usa para representar números decimales.

```go
precio := 19.99
temperatura := 36.5
```

Se utiliza cuando necesitamos trabajar con valores que pueden tener parte decimal.

---

## `string`

El tipo `string` se usa para representar texto.

Los valores de tipo `string` se escriben entre comillas dobles.

```go
nombre := "Jose Manuel"
ciudad := "Sevilla"
```

Ejemplos de `string`:

```go
"Hola"
"DevOps"
"Golang"
```

---

## `bool`

El tipo `bool` se usa para representar valores booleanos.

Solo puede tener dos valores:

```go
true
false
```

Ejemplo:

```go
activo := true
servidorCaido := false
```

Se usa mucho en condiciones.

```go
if activo {
    fmt.Println("El servidor está activo")
}
```

---

## Inferencia de tipos

La inferencia de tipos significa que Go puede detectar automáticamente el tipo de una variable según el valor que le asignamos.

Ejemplo:

```go
nombre := "Jose Manuel"
edad := 25
precio := 19.99
activo := true
```

Go interpreta automáticamente:

```go
nombre // string
edad   // int
precio // float64
activo // bool
```

Esto hace que el código sea más corto y limpio.

---

## Variables no usadas

En Go no se permite declarar variables que luego no se utilizan.

Ejemplo incorrecto:

```go
package main

func main() {
    edad := 25
}
```

Este código da error porque la variable `edad` se declara, pero no se usa.

Para solucionarlo, debemos usar la variable:

```go
package main

import "fmt"

func main() {
    edad := 25
    fmt.Println(edad)
}
```

Esto ayuda a mantener el código limpio y evitar variables innecesarias.

---

## Imports no usados

En Go tampoco se permite importar paquetes que no se utilizan.

Ejemplo incorrecto:

```go
package main

import "fmt"

func main() {
}
```

Este código da error porque se importa `fmt`, pero no se usa.

Para solucionarlo, debemos usar el paquete importado:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hola, Go")
}
```

O eliminar el import si no hace falta.

---

# Resumen rápido

| Concepto            | Explicación                                       |
| ------------------- | ------------------------------------------------- |
| `var`               | Declara variables                                 |
| `:=`                | Declara variables de forma corta                  |
| `const`             | Declara valores constantes                        |
| `int`               | Números enteros                                   |
| `float64`           | Números decimales                                 |
| `string`            | Texto                                             |
| `bool`              | Verdadero o falso                                 |
| Inferencia de tipos | Go detecta el tipo automáticamente                |
| Variables no usadas | Go da error si declaras una variable y no la usas |
| Imports no usados   | Go da error si importas un paquete y no lo usas   |

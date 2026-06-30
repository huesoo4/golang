# Comandos para Golang

---

## Ejecución desde terminal

Para ejecutar un archivo con Go desde la terminal sin ejecutar, accedemos al directorio donde se encuentra el archivo y ejecutamos:

```bash
go run nombre_archivo.go 
```

Este comando compila de manera temporal y ejecuta directamente, pero no crea un binario

---

## Compilación desde terminal

Para poder compilar el archivo nuevamente accedemos al directorio donde se encuentra el fichero y ejecutamos:

```bash
 go build nombre_archivo.go
```

Y para poder ejecutarlo una vez compilado:

```bash
./nombre_archivo
```
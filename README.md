# Ruta Práctica de Aprendizaje de Go

Repositorio dedicado a mi aprendizaje práctico de **Go**, también conocido como **Golang**.

El objetivo de este repositorio es avanzar desde los fundamentos del lenguaje hasta la creación de herramientas reales orientadas a un perfil **DevOps**, trabajando con ejercicios progresivos, proyectos pequeños, proyectos medianos y un proyecto final de portfolio.

---

## Objetivo del repositorio

Aprender Go mediante práctica real, no solo teoría.

Esta ruta está diseñada para desarrollar habilidades útiles en:

* Desarrollo de herramientas de terminal.
* Automatización de tareas.
* Lectura y análisis de logs.
* Trabajo con archivos y JSON.
* Creación de APIs básicas.
* Concurrencia.
* Testing.
* Organización profesional de proyectos.
* Desarrollo de herramientas útiles para sistemas y DevOps.

---

## ¿Por qué Go?

Go es un lenguaje muy utilizado en entornos de infraestructura, backend, cloud y DevOps.

Muchas herramientas importantes del ecosistema DevOps están escritas en Go, como:

* Docker
* Kubernetes
* Terraform
* Prometheus
* Grafana
* Traefik
* Helm

Aprender Go me permite entender mejor cómo funcionan este tipo de herramientas y, además, crear mis propias utilidades para automatización, monitorización y administración de sistemas.

---

## Estructura de la ruta

La ruta está dividida en niveles del **0 al 100**, avanzando de forma progresiva.

```text
go-learning-path/
├── 00-10-primer-contacto/
├── 10-20-variables-tipos/
├── 20-30-condicionales/
├── 30-40-bucles-slices/
├── 40-50-maps-funciones/
├── 50-60-structs-metodos/
├── 60-70-errores-modulos-paquetes/
├── 70-80-archivos-json-logs/
├── 80-90-http-apis-clientes/
├── 90-100-concurrencia-testing/
├── proyecto-final/
└── README.md
```

---

# Nivel 0-10: Primer contacto con Go

## Objetivo

Entender cómo se estructura un programa básico en Go.

## Conceptos principales

* `package main`
* `import`
* `func main()`
* `fmt.Println`
* Archivos `.go`
* `go run`
* `go build`
* Diferencia entre ejecutar y compilar

## Proyecto

### Tarjeta de presentación por terminal

Programa básico que muestra una presentación personal desde la terminal.

Debe mostrar:

* Nombre
* Edad
* Ciudad
* Tecnologías en aprendizaje
* Objetivo profesional

Ejemplo:

```text
==============================
 Perfil DevOps
==============================
Nombre: Jose Manuel
Edad: 25
Ciudad: Sevilla
Tecnologías: Linux, Docker, Go, Python
Objetivo: Convertirme en DevOps Junior
==============================
```

## Aprendizaje esperado

Al terminar este nivel, debo saber crear, ejecutar y compilar un programa básico en Go.

---

# Nivel 10-20: Variables, constantes y tipos básicos

## Objetivo

Aprender a guardar información en variables y entender los tipos básicos del lenguaje.

## Conceptos principales

* `var`
* `:=`
* `const`
* `int`
* `float64`
* `string`
* `bool`
* Inferencia de tipos
* Variables no usadas
* Imports no usados

## Proyecto

### Calculadora de salario neto simple

Programa que calcula un salario neto aproximado a partir de un salario bruto y una retención.

Ejemplo:

```text
Salario bruto: 2000
Retención: 15%
Salario neto: 1700
```

## Aprendizaje esperado

Al terminar este nivel, debo saber trabajar con variables, constantes, tipos básicos y operaciones matemáticas simples.

---

# Nivel 20-30: Condicionales

## Objetivo

Aprender a tomar decisiones dentro de un programa.

## Conceptos principales

* `if`
* `else`
* `else if`
* `switch`
* Operadores de comparación
* Operadores lógicos

## Proyecto

### Evaluador de estado de servidor

Programa que evalúa el estado de los recursos principales de un servidor.

Datos de ejemplo:

```go
cpu := 75
ram := 82
disk := 91
```

Salida esperada:

```text
CPU: Advertencia
RAM: Crítico
Disco: Crítico

Estado general del servidor: Crítico
```

## Aprendizaje esperado

Al terminar este nivel, debo saber usar condicionales para aplicar reglas lógicas a problemas reales.

---

# Nivel 30-40: Bucles y slices

## Objetivo

Aprender a trabajar con listas de datos.

En Go, los `slices` son una de las estructuras más usadas. Son parecidos a las listas en Python.

## Conceptos principales

* `for`
* `range`
* `break`
* `continue`
* Slices
* `append`
* Recorrido de listas
* Contadores
* Acumuladores

## Proyecto

### Analizador de uso de CPU

Programa que analiza varios valores de CPU y muestra su estado.

Ejemplo:

```go
cpus := []int{25, 40, 65, 82, 91, 30}
```

Salida esperada:

```text
CPU 25% - Normal
CPU 40% - Normal
CPU 65% - Advertencia
CPU 82% - Crítico
CPU 91% - Crítico
CPU 30% - Normal

Total críticos: 2
Media CPU: 55.5%
```

## Aprendizaje esperado

Al terminar este nivel, debo saber recorrer listas, filtrar datos, contar elementos y calcular valores agregados.

---

# Nivel 40-50: Maps y funciones

## Objetivo

Aprender a organizar datos clave-valor y reutilizar código mediante funciones.

## Conceptos principales

* `map`
* Crear mapas
* Leer valores
* Añadir valores
* Eliminar valores
* Comprobar si una clave existe
* Funciones
* Parámetros
* Valores de retorno

## Proyecto

### Inventario simple de servidores

Programa que gestiona un inventario básico de servidores.

Ejemplo:

```go
servidores := map[string]string{
    "web01": "activo",
    "db01":  "activo",
    "old01": "apagado",
}
```

El programa debe:

* Mostrar todos los servidores.
* Contar servidores activos.
* Contar servidores apagados.
* Evaluar estados mediante funciones.
* Mostrar un resumen final.

## Aprendizaje esperado

Al terminar este nivel, debo saber organizar datos con maps y separar lógica usando funciones.

---

# Nivel 50-60: Structs y métodos

## Objetivo

Aprender a crear estructuras de datos propias.

En Go no se trabaja con clases como en otros lenguajes. En su lugar, se usan mucho los `structs`.

## Conceptos principales

* `struct`
* Campos
* Instanciar structs
* Slices de structs
* Métodos
* Métodos con puntero

## Proyecto

### Monitor básico de servidores

Programa que modela varios servidores usando structs.

Cada servidor debe tener:

* Nombre
* IP
* CPU
* RAM
* Disco
* Estado

Ejemplo:

```go
type Servidor struct {
    Nombre string
    IP     string
    CPU    int
    RAM    int
    Disco  int
}
```

El programa debe mostrar el estado de cada recurso y un resumen final.

## Aprendizaje esperado

Al terminar este nivel, debo saber modelar datos reales usando structs, slices de structs y métodos.

---

# Nivel 60-70: Errores, módulos y paquetes

## Objetivo

Aprender a organizar proyectos de forma más profesional.

## Conceptos principales

* `error`
* `nil`
* `errors.New`
* `fmt.Errorf`
* `go mod init`
* `go mod tidy`
* Paquetes propios
* Funciones exportadas
* Separación de archivos y carpetas

## Proyecto

### Monitor de servidores organizado por paquetes

Refactorización del monitor anterior usando módulos y paquetes.

Estructura recomendada:

```text
monitor-go/
├── go.mod
├── main.go
├── monitor/
│   └── monitor.go
└── server/
    └── server.go
```

El proyecto debe incluir:

* Un paquete para servidores.
* Un paquete para evaluación de recursos.
* Manejo de errores.
* Código limpio.
* Funciones exportadas.
* Formateo con `go fmt ./...`.

## Aprendizaje esperado

Al terminar este nivel, debo saber crear un módulo en Go, dividir código en paquetes y manejar errores de forma explícita.

---

# Nivel 70-80: Archivos, JSON y logs

## Objetivo

Aprender a trabajar con datos externos.

Esto es fundamental para DevOps porque muchas herramientas trabajan con logs, JSON, ficheros de configuración y salidas de comandos.

## Conceptos principales

* Leer archivos
* Escribir archivos
* `os.ReadFile`
* `os.WriteFile`
* `encoding/json`
* `json.Marshal`
* `json.Unmarshal`
* Procesamiento de logs

## Proyecto

### Analizador de logs

Herramienta en Go que lee un archivo de logs y genera un resumen.

Archivo de entrada:

```text
INFO servicio iniciado
ERROR conexión rechazada
WARNING uso de disco alto
ERROR timeout
INFO backup completado
ERROR fallo autenticación
```

Salida esperada:

```text
Resumen de logs
----------------
INFO: 2
WARNING: 1
ERROR: 3

Estado general: Crítico
```

## Aprendizaje esperado

Al terminar este nivel, debo saber leer archivos, procesar texto, contar eventos y generar reportes simples.

---

# Nivel 80-90: HTTP, APIs y clientes

## Objetivo

Aprender a crear y consumir APIs.

Esto es clave para backend, DevOps, automatización y herramientas cloud.

## Conceptos principales

* `net/http`
* Crear servidor HTTP
* Rutas
* Métodos GET y POST
* JSON en HTTP
* Cliente HTTP
* Status codes
* APIs REST básicas

## Proyecto

### API de inventario DevOps

API REST en Go para consultar servidores.

Endpoints mínimos:

```text
GET /health
GET /servers
GET /servers/{name}
```

Ejemplo de datos:

```go
servidores := []Servidor{
    {Nombre: "web01", IP: "192.168.1.10", CPU: 40, RAM: 50, Disco: 60},
    {Nombre: "db01", IP: "192.168.1.20", CPU: 90, RAM: 85, Disco: 92},
}
```

## Aprendizaje esperado

Al terminar este nivel, debo saber crear una API básica, devolver JSON y estructurar endpoints sencillos.

---

# Nivel 90-100: Concurrencia, testing y proyecto final

## Objetivo

Entrar en partes más potentes de Go.

Go destaca por su modelo de concurrencia, útil para ejecutar varias tareas al mismo tiempo.

## Conceptos principales

* Goroutines
* `go func()`
* Channels básicos
* `sync.WaitGroup`
* Tests
* `go test`
* Benchmarks básicos
* Proyecto final completo

## Proyecto

### Checker concurrente de URLs

Programa que comprueba varias URLs de forma concurrente.

Ejemplo:

```go
urls := []string{
    "https://google.com",
    "https://github.com",
    "https://openai.com",
}
```

Primero se debe hacer de forma secuencial y después usando goroutines y `sync.WaitGroup`.

## Aprendizaje esperado

Al terminar este nivel, debo saber lanzar tareas concurrentes y escribir tests básicos para validar funciones importantes.

---

# Proyecto final: Go DevOps Monitor

## Descripción

Herramienta CLI escrita en Go para analizar el estado de servidores a partir de un archivo JSON y generar un reporte de salud del sistema combinando métricas de CPU, RAM, disco y logs.

## Funcionalidades mínimas

La herramienta debe poder:

1. Leer un archivo JSON con servidores.
2. Evaluar CPU, RAM y disco.
3. Mostrar el estado de cada servidor.
4. Leer un archivo de logs.
5. Contar errores, warnings e info.
6. Generar un resumen final.
7. Guardar un reporte en un archivo.
8. Tener el código organizado por paquetes.
9. Tener tests para las funciones principales.

## Estructura recomendada

```text
devops-monitor/
├── go.mod
├── main.go
├── data/
│   ├── servers.json
│   └── logs.txt
├── internal/
│   ├── monitor/
│   │   └── monitor.go
│   ├── logs/
│   │   └── logs.go
│   └── report/
│       └── report.go
└── tests/
```

## Ejemplo de `servers.json`

```json
[
  {
    "nombre": "web01",
    "ip": "192.168.1.10",
    "cpu": 40,
    "ram": 55,
    "disco": 60
  },
  {
    "nombre": "db01",
    "ip": "192.168.1.20",
    "cpu": 91,
    "ram": 88,
    "disco": 93
  }
]
```

## Ejemplo de `logs.txt`

```text
INFO servicio iniciado
ERROR fallo en base de datos
WARNING uso de disco alto
ERROR timeout
INFO backup completado
```

## Salida esperada

```text
DevOps Monitor
==============

Servidores
----------

web01 - 192.168.1.10
CPU: 40% - Normal
RAM: 55% - Normal
Disco: 60% - Advertencia
Estado general: Advertencia

db01 - 192.168.1.20
CPU: 91% - Crítico
RAM: 88% - Crítico
Disco: 93% - Crítico
Estado general: Crítico

Logs
----

INFO: 2
WARNING: 1
ERROR: 2

Resumen final
-------------

Servidores críticos: 1
Estado logs: Advertencia
Estado global: Crítico
```

---

## Extras opcionales

Cuando la versión básica esté terminada, se pueden añadir mejoras como:

* Argumentos por terminal.
* Flags para elegir archivo JSON.
* Flags para elegir archivo de logs.
* Exportar reporte en JSON.
* Endpoint HTTP `/health`.
* Endpoint HTTP `/report`.
* Dockerizar la aplicación.
* Crear GitHub Actions para ejecutar tests.
* Añadir documentación técnica más completa.

---

# Ruta resumida por semanas

## Semana 1: Fundamentos

Objetivo:

* Sintaxis básica.
* Variables.
* Tipos.
* Condicionales.
* Bucles.
* Slices.

Proyectos:

* Tarjeta de presentación.
* Calculadora simple.
* Evaluador de servidor.
* Analizador de CPU.

---

## Semana 2: Datos y organización

Objetivo:

* Maps.
* Funciones.
* Structs.
* Métodos.
* Módulos.
* Paquetes.

Proyectos:

* Inventario de servidores.
* Monitor básico de servidores.
* Monitor organizado por paquetes.

---

## Semana 3: Archivos y JSON

Objetivo:

* Leer archivos.
* Escribir archivos.
* Procesar logs.
* Leer JSON.
* Convertir structs a JSON.

Proyectos:

* Analizador de logs.
* Lector de inventario desde JSON.

---

## Semana 4: HTTP y APIs

Objetivo:

* Crear servidor HTTP.
* Devolver JSON.
* Crear endpoints.
* Consumir APIs.

Proyectos:

* API de inventario DevOps.
* Endpoint `/health`.
* Endpoint `/servers`.

---

## Semana 5: Concurrencia y testing

Objetivo:

* Goroutines.
* WaitGroup.
* Tests.
* Proyecto final.

Proyectos:

* Checker concurrente de URLs.
* Tests de monitorización.
* DevOps Monitor CLI.

---

# Reglas de práctica

## 1. No copiar y pegar sin entender

Puedo consultar ejemplos, pero después debo escribir el código por mi cuenta.

---

## 2. Separar cada ejercicio en una carpeta

Para evitar errores como:

```text
main redeclared in this block
```

Se recomienda usar una estructura como esta:

```text
go-practica/
├── 01-hola-mundo/
│   └── main.go
├── 02-variables/
│   └── main.go
├── 03-condicionales/
│   └── main.go
└── 04-slices/
    └── main.go
```

---

## 3. Usar comandos básicos de Go

```bash
go run .
```

```bash
go fmt ./...
```

```bash
go test ./...
```

```bash
go mod tidy
```

---

## 4. Leer los errores de compilación

Errores normales al principio:

```text
declared and not used
```

Significa que declaré una variable y no la usé.

```text
imported and not used
```

Significa que importé un paquete y no lo usé.

```text
undefined: Println
```

Seguramente quería escribir:

```go
fmt.Println()
```

```text
main redeclared in this block
```

Significa que tengo varias funciones `main` dentro de la misma carpeta.

---

# Checklist de progreso

## Nivel básico

* [ ] Sé crear un `main.go`.
* [ ] Sé usar `fmt.Println`.
* [ ] Sé declarar variables.
* [ ] Sé usar `int`, `float64`, `string` y `bool`.
* [ ] Sé usar `if`, `else` y `switch`.
* [ ] Sé usar bucles `for`.
* [ ] Sé recorrer slices con `range`.

---

## Nivel intermedio

* [ ] Sé crear funciones.
* [ ] Sé devolver valores desde funciones.
* [ ] Sé manejar errores básicos.
* [ ] Sé usar maps.
* [ ] Sé crear structs.
* [ ] Sé crear métodos.
* [ ] Sé organizar código en paquetes.
* [ ] Sé crear un módulo con `go mod init`.

---

## Nivel práctico DevOps

* [ ] Sé leer archivos de texto.
* [ ] Sé contar errores en logs.
* [ ] Sé leer JSON.
* [ ] Sé convertir structs a JSON.
* [ ] Sé crear una API básica.
* [ ] Sé hacer peticiones HTTP.
* [ ] Sé crear una herramienta CLI simple.
* [ ] Sé compilar un binario ejecutable.

---

## Nivel avanzado inicial

* [ ] Sé lanzar goroutines.
* [ ] Sé usar `sync.WaitGroup`.
* [ ] Sé escribir tests básicos.
* [ ] Sé ejecutar `go test ./...`.
* [ ] Sé estructurar un proyecto mediano.
* [ ] Sé crear un README técnico.
* [ ] Sé dockerizar una app simple en Go.

---

# Tecnologías trabajadas

* Go
* CLI
* JSON
* HTTP
* APIs REST
* Testing
* Docker
* GitHub Actions
* Automatización
* Monitorización básica
* Procesamiento de logs

---

# Qué demuestra este repositorio

Este repositorio demuestra una progresión práctica en Go aplicada a un perfil DevOps.

Al completarlo, habré trabajado:

* Sintaxis básica de Go.
* Control de flujo.
* Estructuras de datos.
* Funciones.
* Structs.
* Métodos.
* Manejo de errores.
* Módulos y paquetes.
* Lectura y escritura de archivos.
* JSON.
* APIs HTTP.
* Concurrencia.
* Testing.
* Organización profesional de proyectos.
* Creación de herramientas útiles para sistemas.

---

# Objetivo final

El objetivo final no es memorizar toda la sintaxis de Go, sino construir proyectos reales y entender por qué funcionan.

Go se aprende muy bien practicando, porque es un lenguaje directo, claro y muy enfocado a construir herramientas eficientes.

Esta ruta busca convertirme en una persona capaz de usar Go para crear herramientas de automatización, monitorización, APIs sencillas y utilidades prácticas dentro de un entorno DevOps.

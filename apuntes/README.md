# Apuntes de Golang

Go o Golang es un lenguaje compilado desarrollado en Google, está inspirado en el la sintaxis de C e intenta ser dinámico como Python.

Go genera archivos .EXE portables a cualquier sistema operativo y versión (Su EXE contiene todo lo necesario para ejecutarse).

Go se creo pensando en los procesadores modernos los cuales poseen mas de un núcleo lo que permite que puedan ejecutar procesos multi-hilo, lenguajes antiguos como C por ejemplo no tenían en cuenta esto ya que en ese tiempo no existía el hardware.

Es un lenguaje fuertemente tipado, las variables deben ser declaradas antes de ser utilizadas.

Go obliga al desarrollador a llevar buenas prácticas, por ejemplo no nos deja declarar variables que no se utilizan o cuando declaramos alguna función o método global, go no obliga a comentarlas. Go no te deja compilar si no respetamos estas normas.

Go demostró ser el lenguaje ideal para grandes desarrollos con miles y miles usuarios concurrentes y millones de transacciones por segundo.

En Go no hace falta usar `;` punto y coma.

Las funciones en Go pueden devolver mas de un valor. Esto es algo particular de este lenguaje, que es muy util.

Se pueden desarrollar en Go tanto instrucciones sincrónicas como asincrónicas.

Solo existe la instrucción `FOR` para iteraciones, no existe While, ni Do Until, ni nada similar.

Go no es un lenguaje orientado a objetos y resuelve lo que conocemos como POO, con Estructuras, funciones, métodos e interfaces

El scope de variables, métodos y funciones se determinan con nombres en mayúsculas y minúsculas. Si una variables o función comienza con una letra en minúscula es privada, si comienza con una letra en mayúscula es pública.

## Sintaxis

Algo peculiar de Go es que nos elimina las cosas que tenemos declaradas pero que no estamos usando.

Para hacer un Hola mundo en Go, podemos escribir:

```go

package main

import "fmt"

fmt.Println("Hola mundo")
```

Para ejecutar el programa podemos escribir en la consola:

```bash
go run main.go
```

Lo que hace este comando es compilar en memoria el programa y ejecutarlo, no nos genera un archivo .EXE.

Si lo que quiero es compilar el programa para poder ejecutarlo en cualquier otro computador, debo generar un archivo .EXE, para esto debo escribir el comando:

```bash
go build main.go
```

Con esto generamos un binario que podemos ejecutar en cualquier computador.

Este comando compila el programa pero no lo ejecuta.

### Variables

En Go existen tres tipos de variables, las numéricas, las cadenas y las booleanas, estos son los tres tipos principales, cada uno de estos tipos tiene sub divisiones, por ejemplo las numéricas tienen los enteros, las sin signo, las con signo y las flotantes.

Recordar que Go es un lenguaje tipado por lo que debemos indicar de que tipo es la variable que estamos declarando.

```go
var nombre string = "Juan"
var edad int = 23
var esVerdadero bool = true
```

También tenemos los tipos:

```go
var decimal float32 = 1.23
var decimal2 float64 = 1.23
var entero8 int8 = 127
var entero16 int16 = 32767
var entero32 int32 = 2147483647
var entero64 int64 = 9223372036854775807

// También están los enteros con y sin signo

var sinSigno uint = 1
var sinSigno2 uint8 = 255
var sinSigno3 uint16 = 65535
var sinSigno4 uint32 = 4294967295
var sinSigno5 uint64 = 18446744073709551615
```

Los tipos de datos numéricos están sub divididos por la cantidad de bits que ocupan en memoria (8, 16, 32, 64) y el tipo de dato que representan (int, uint, float32, float64), esto condiciona el tamaño del numero que se puede almacenar en la variable.

La "u" en `uint` significa unsigned, es decir que no tiene signo, por lo que solo puede almacenar valores positivos.

Hay otra forma de declarar una variable con el operador `:=`, con este operador Go infiere el tipo de dato de la variable, esta asignación solo la puedo usar una vez por variable.

```go
package main

func main() {
    numero := 1
}
```

También podemos declarar varias variables en una sola línea:

```go
package main

func main() {
    var numero, numero1 int = 1, 2

    num3, num4, num5, texto := 1, 2, 3, "Hola"
}
```

Un dato a tener en cuenta es que en `Go` no existe el concepto de `null`, el cual seria el valor por defecto de una variable que no se inicializó con algún valor en otro lenguaje.

En `Go` las variables `int` por defecto son `0`, `string` por defecto es `""` (string vacío) y `bool` por defecto es `false`.

Otra peculiaridad de `Go` es que no puedo asignar el valor de una variable `float` a una variable `int`, por ejemplo:

```go
package main

func main() {
    var numero float32 = 1.23
    var numero2 int = numero
}
```

Go es uno de los lenguajes más severos con los tipos de variables, esto significa que no puedo mover un flotante a un entero o viceversa, esto incluso va mas allá, por ejemplo si tengo una variable de tipo `int64` y quiero mover su valor a un entero `int` genérico, tampoco puedo hacerlo, solo puedo mover valores entre variables del mismo tipo, exactamente del mismo tamaño.

```go
package main

func main() {
    var numero int64 = 1234567890
    var numero2 int = numero // Error
}
```

Para poder hacer este tipo de cosas existen formas, por ejemplo convertir la variables:

```go
package main

func main() {
    var numero int64 = 1234567890
    var numero2 int = int(numero)
}
```

Los tipos de datos no son solo eso, sino que también son funciones, en `Go` prácticamente todo son funciones.

Esto de convertir los tipos de datos tampoco es tan flexible, por ejemplo si trato de convertir un `int` a un `string` no me dará error, pero al momento de imprimir el valor por consola no me devolverá nada.

Para convertir un `int` a un `string` debo usar la función `strconv.Itoa`:

```go
package main

func main() {
    var moneda float64 = 0
    var numero2 string = strconv.Itoa(int(moneda)) // la función `strconv.Itoa` solo convierte enteros genéricos a strings
}
```

## Condicionales

En `Go` no es necesario poner paréntesis al momento de escribir una condición, por ejemplo:

```go
package main

import "fmt"

var estado bool

func main() {
	estado = true

	if estado == true {
		fmt.Println(estado)
	} else {
		fmt.Println("El estado es false")
	}
}
```

`Go` también me permite hacer cosas que no puedo hacer en otros lenguajes como por ejemplo asignar una variable en medio de un `if`.

```go
package main

import "fmt"

var estado bool

func main() {
	estado = true

	if estado = false; estado == true {
		fmt.Println(estado)
	} else {
		fmt.Println("El estado es false")
	}
}
```

También existe el `else if`, para concatenar varias condiciones.

### Switch

```go
package main

import "fmt"

func main() {

	switch numero := 5; numero {
	case 1:
		fmt.Println("uno")
	case 2:
		fmt.Println("dos")
	case 3:
		fmt.Println("tres")
	case 4:
		fmt.Println("cuatro")
	case 5:
		fmt.Println("cinco")
	default:
		fmt.Println("mayor a 5")
	}
}
```

A diferencia de otros lenguajes no es necesario romper el switch con un `break`.

## Mostrar y leer datos por consola

```go
package main

import (
	"fmt"
	"os"
	"bufio"
)

var numero int
var numero1 int

func main() {
	fmt.Println("Introduce un numero")
	fmt.Scanln("%d", &numero)

	fmt.Println("Ingrese otro numero")
	fmt.Scanln("%d", &numero1)
}
```

En `Go` hay una diferencia entre usar `fmt.ScanF` y `fmt.Scanln`, si es que estamos usándolo en windows, esto es debido a como maneja windows los saltos de línea, lo cual seria `\n\r`, mientras que en Linux es solo `\n`, por lo que si usamos `fmt.ScanF` en windows, el programa no funcionaría, porque tomara el `\r` como el segundo parámetro, esto se soluciona con `fmt.Scanln`.
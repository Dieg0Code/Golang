# Apuntes de Golang

Go o Golang es un lenguaje compilado desarrollado en Google, está inspirado en el la sintaxis de C e intenta ser dinámico como Python.

Go genera archivos .EXE portables a cualquier sistema operativo y versión (Su EXE contiene todo lo necesario para ejecutarse).

Go se creo pensando en los procesadores modernos los cuales poseen mas de un núcleo lo que permite que puedan ejecutar procesos multi-hilo, lenguajes antiguos como C por ejemplo no tenían en cuenta esto ya que en ese tiempo no existía el hardware.

Es un lenguaje fuertemente tipado, las variables deben ser declaradas antes de ser utilizadas, sino lo hacemos Golang nos borra la variable.

Go obliga al desarrollador a llevar buenas prácticas, por ejemplo no nos deja declarar variables que no se utilizan o cuando declaramos alguna función o método global, go nos obliga a comentarlas. Go no te deja compilar si no respetamos estas normas.

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

// También están los enteros sin signo

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
    var numero2 int = numero // Error
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

Una alternativa es usar `bufio.NewScanner` para leer los datos por consola, por ejemplo:

```go
package main

import (
	"fmt"
	"os"
	"bufio"
)

var numero int
var numero1 int
var result int
var leyenda string

func main() {
	fmt.Println("Introduce un numero")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	result = numero + numero1
	fmt.Println(leyenda, result)
}
```

## Ciclos en Go

En Go solo existe el ciclo `for` para iterar sobre una secuencia de datos, esto es debido a que en `go` esta sentencia está enriquecida de tal manera que puede ser usada para reemplazar cualquier otro ciclo.

```go
package main

import "fmt"

func main() {

	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

}
```

Esta es una forma de crear un ciclo for, pero podemos hacerlo también de una manera más tradicional:

```go

func main() {

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

}
```

En `Go` no es necesario poner los paréntesis al escribir un for.

Podemos romper un ciclo for con un `break`:

```go
func main() {

	num := 0
	for {
		fmt.Println("continue")
		fmt.Println("ingrese un numero")
		fmt.Scanln(&num)
		if num == 0 {
			break
		}
	}
}
```

Escribir un ``for`` de esta manera, sin condiciones, hace que sea infinito, es por eso que debemos poner un `break` en alguna condicional, a menos que queramos que sea infinito a propósito.

Así mismo, como existe el `break`, podemos romper un ciclo for con un `continue`:

```go
func main() {

	var i := 0

	for i < 10 {
		fmt.Println("\n Valor de i: %d", i)
		if i == 5 {
			fmt.Println("Multiplicamos por 2 \n")
			i = i * 2
			continue
		}

		fmt.Println("Paso por aquí")
		i++
	}

}
```

El `continue` también rompe el ciclo `for`, pero no de la misma manera, el `continue` lo que hace es que reinicia el ciclo, es decir, vuelve al principio sin dejar que se ejecuten las instrucciones que están debajo del `continue`.

En `Go` también existe algo llamado `GOTO` el cual es una forma de nombrar una porción de código con el fin de poder llamarla luego en base el nombre que le demos.

```go
func main() {

	var i int = 0

	RUTINA:
		for i < 10 {
			if i == 4 {
				i = i + 2
				fmt.Println("Voy a RUTINA sumando 2 a i")
				goto RUTINA
			}

			fmt.Printf("Valor de i: %d\n", i)
			i++
		}
}

```

De esta forma, usando la instrucción `goto`, podemos volver a ejecutar esa porción de código que etiquetamos con el nombre de `RUTINA` en este caso.

## Funciones en Go

Las funciones en `Go` se declaran usando la palabra reservada `func`, seguida del nombre de la función, los parámetros de la función, el tipo de dato que retornará y el cuerpo de la función.

En `Go` todo son funciones, no existen los métodos, en Go estos son funciones en si mismos que solo no devuelven ningún valor.

```go
package main

import "fmt"

func main() {

	fmt.Println(uno(5)) // 6	
}

func uno(numero int) int {
	return numero + 1
}
```

En `Go` existe la peculiaridad de que una función puede devolver dos valores:

```go
func dos(numero int) (int, bool) {
	if numero == 5 {
		return 1, true
	} else {
		return 0, false
	}
}
```

Entonces, para poder mostrar el resultado de la función no podemos solo imprimir el valor con `fmt.Println`, sino que debemos capturar los valores de retorno en dos variables distintas:

```go
func main() {

	numero, estado := dos(5)
	fmt.Println(numero, estado) // 1 true
}
```

Otra variante al declarar una función en go es poner un nombre al valor de retorno, por ejemplo:

```go
func tres(numero int) (z int) {
	z = numero * 2
	return
}
```

De esta forma no es necesario especificar el valor de retorno, con solo el `return` se entiende que estamos devolviendo el valor `z` que se declaró al principio de la función.

### Parámetros variables

Los parámetros variables son parámetros dinámicos en donde yo no sé el número de parámetros de entrada.

Hay una falencia que tiene `Go` que es que no tiene sobrecarga de funciones o sobrecarga de métodos, a diferencia de lenguajes como C# o Java, en donde podemos escribir el mismo método o función N veces con diferentes tipos de parámetros de entrada. Esto es muy útil en programación porque permite que uno ejecute distinto tipo de código llamando al mismo método, pero dependiendo del parámetro que le pasemos va a comportarse de una manera u otra.

`Go` lo que si tiene implementado y sirve como una forma de polimorfismo son los parámetros variables, por ejemplo:	

```go
func Calculo(numero ...int) int {
	total := 0

	for _, num := range numero {
		total += num
	}

	return total
}
```

La palabra reservada `range` me devuelve siempre dos valores, lo que coloque a continuación de range debe ser una lista, un vector, etc.

En este caso, al ser parámetros variables, los convierte en una lista e itera por sobre cada uno de los elementos.


El primer valor que devuelve range es el numero del elemento, nos dice si es el elemento uno, dos, tres, etc., en este caso no necesitamos ese dato por lo que puedo nombrarlo como `_`. En `Go` esto se usa para nombrar a los datos que no usaremos. Go no nos permite saltarnos pasos, si una función devuelve dos valores no puedo ignorar uno de estos.


```go
func main() {

	fmt.Println(Calculo(1, 2, 3, 4, 5)) // 15
}
```

Lo que hace nuestra función es sumar todos los valores que le pasemos como parámetro y nos devuelve el total.

Si queremos, también podemos mostrar el valor del item en el que estamos iterando:

```go
func Calculo(numero ...int) int {
	total := 0

	for i, num := range numero {
		fmt.Printf("El valor de i es: %d\n", i)
		fmt.Printf("El valor de num es: %d\n", num)
		total += num
	}

	return total
}

func main() {

	fmt.Println(Calculo(1, 2, 3, 4, 5)) // el valor de i es: 0 
										// el valor de num es: 1
										// ... 1, 2, 3, 4, 5
	total := Calculo(1, 2, 3, 4, 5)
	fmt.Println(total) // 15

}
```
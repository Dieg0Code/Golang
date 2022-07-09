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

## Funciones anónimas y closures

Las funciones anónimas son funciones que no tienen nombre, en go podemos crear una variable y ponerle que sea de tipo `func`, con esto podemos aislar el código de la función dentro de una variable y así luego poder modificar ese código en `tiempo de ejecución`, cosa que no puedo hacer con una función normal.

Por ejemplo, no puedo crear una función suma y luego modificarla en tiempo ded ejecución, para esto tenemos que crear una función anónima:

```go
package main

import "fmt"

var Calculo func(int, int) int = func(a int, b int) int {
	return a + b
}

func main() {

	fmt.Println("Sumo 5 + 7 = %d \n", Calculo(5, 7)) // Sumo 5 + 7 = 12

}
```

Para crear una variable de tipo `func` debemos primero declarar cuanto parámetros recibirá, los tipos y el tipo del valor de retorno, luego a la derecha de `=` declaramos la función es sí misma, nombramos los parámetros, ponemos el tipo de valor de retorno y luego el cuerpo de la función.

De esta forma estamos creando una función anónima cuyo resultado es almacenado en la variable `Calculo`.

Pero ¿cual es la gracia de las funciones anónimas? si todo eso podemos hacerlo con una función normal. Lo especial de esto es que puedo redefinir `Calculo` cosa que no puedo hacer con una función normal.

Por ejemplo:

```go
package main

import "fmt"

var Calculo func(int, int) int = func(a int, b int) int {
	return a + b
}


func main() {

	Calculo = func(a int, b int) int {
		return a - b
	}

	fmt.Println("Restamos 6 - 4 = %d \n", Calculo(6, 4)) // Restamos 6 - 4 = 2

}
```

Esto es lo mas parecido que tenemos en Go a una sobrecarga, con la diferencia de que en otros lenguajes podemos cambiar la cantidad de parámetros, pero en go no podemos, en go tenemos que respetar la misma cantidad de parámetros de entrada, el tipo de parámetros de entrada y solo así puedo redefinir el cuerpo de la función.

Podemos también crear una función normal cuyo valor de retorno sea una función anónima.

```go

func Operaciones() {
	resultado := func() int {

		var a int = 23
		var b int = 13

		return a + b
	}

	fmt.Println(resultado()) // 36
}

func main() {

	Operaciones() // 36

}

```

## Closures

Los `closures` son un concepto de funciones anónimas que tienen que ver con la protección y el aislamiento de código.

Una de las características de los closures es que pueden acceder al valor de variables creadas por fuera de scope de la función.

Una closure es una función cuyo valor de retorno es una función anónima.

```go

func Tabla(valor int) func() int {

	numero := valor
	secuencia := 0

	return func() int {
		secuencia++
		return numero * secuencia
	}

}
```

Tenemos dos `return`, uno es el de la función anónima y el otro es el de la función normal.

```go
func main() {

	// CLosure
	tablaDel := 2
	MiTabla := Tabla(tablaDel)
	
	// la variable "MiTabla" se convierte en una función, debido
	// a que el valor de retorno de la función "Tabla"
	// es una función anónima.	

	for i := 1; i < 11; i++ {
		fmt.Println(MiTabla())
	}

}

// El resultado es:
// 2
// 4
// 6
// 8
// 10
// 12
// ...
```

Cuando ejecutamos la función `Tabla()` lo que estamos ejecutando es el `return` de esta, lo cual en este caso es una función anónima, por lo que en realidad estamos ejecutando esa parte del código de la función `Tabla()`.

En la primera ejecución de la función se va a ejecutar la primera parte, en donde declaramos la variable `numero` y la variable `secuencia`.

Cuando declaramos `MiTabla`:

```go
tablaDel := 2

MiTabla := Tabla(tablaDel)
```

No toma todo el código de la función `Tabla()`, sino que toma lo que le devuelve esta función.

## Arreglos estáticos y Slices



En Go los arreglos se pueden declarar especificando el tipo de datos de la colección y la cantidad de elementos que contiene.

```go
package main

import "fmt"

var tabla [10]int

func main() {
	fmt.Println(tabla) // [0 0 0 0 0 0 0 0 0 0]

	tabla[0] = 1
	tabla[1] = 2

	fmt.Println(tabla) // [1 2 0 0 0 0 0 0 0 0]
}
```

En go los tipos de datos se inicializan con valores por defecto, por ejemplo}: los de tipo `int` se inicializan con `0` y los de tipo `string` se inicializan con `""` una cadena vacía y los de tipo `bool` se inicializan con `false`.

Puedo declarar un arreglo y inicializarlo con valores de inmediato.

```go
package main

import "fmt"

func main() {

	tabla := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

}
```

Por lo generar para recorrer los elementos de una colección usamos el ciclo `for`.


```go
package main

import "fmt"

func main() {

	tabla := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

}
```

La función `len()` devuelve la cantidad de elementos de un arreglo, en este caso le decimos al ciclo que siga mientras ` i < len(tabla) `, es decir mientras que `i` sea menor a la cantidad de elementos del arreglo.

Podemos también tener matrices, es decir un arreglo de arreglos.

```go
package main

import "fmt"

var matriz [5][7]int // lo cual significa que tenemos una matriz
					 // de 5 filas y 7 columnas, es decir
					 // 5 arreglos de 7 elementos cada uno.

func main() {

	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz[i]); j++ {
			
			fmt.Printf("%d ", matriz[i][j])

		}
	}

}
```

Las matrices pueden ser recorridas usando un ciclo `for` anidado, uno que recorra las filas y otro que recorra las columnas.

### Slices

Los **vectores dinámicos**, a diferencia de los arreglos que son estáticos en los que debemos respetar la cantidad de elementos con la que se declaró el arreglo, no tenemos ese problema, podemos agregar cuantos elementos queramos.

Debemos tener en cuenta son dinámicos porque los puedo modificar durante el tiempo de ejecución.

```go
package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}	

}
```

Para indicar que es un arreglo dinámico tan solo debemos usar los `[]` sin ninguna longitud dentro, con esto Go entiende que lo que declaramos fue un arreglo de enteros de longitud variable.

```go
package main

import "fmt"

func variante2() {
	elementos := [5]int{1,2,3,4,5}
	porcion := elementos[3:] // Con esta linea estamos diciendo
							 // que porción va a tomar todos los
							 // valores desde el 3 indices hasta el final.
							 // es decir [4,5]
	porcion := elementos[2:5] // desde el indice 2 hasta el 5
}

func main() {



}
```

Los slices nos permiten tomar porciones de arreglos mas grandes, indicándoles el indice de inicio y el indice de fin.

```go
package main

import "fmt"

func variante3() {

	elementos := make([]int, 5, 20)
	fmt.Println(elementos) // [0 0 0 0 0]
	fmt.Println("Largo %d, Capacidad %d", len(elementos), cap(elementos)) // Largo 5, Capacidad 20
}
```

Con `make()` estamos creando un arreglo dinámico de enteros, el cual va a tener una longitud de 5 elementos, pero una capacidad de 20.

```go
package main

import "fmt"

func variante4() {

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Println("Largo %d, Capacidad %d", len(nums), cap(nums)) // Largo 100, Capacidad 128
}
```

## Mapas

Un mapa es una estructura que se puede serializar, como un vector, es decir podemos guardar una colección de elementos que podemos recorrer luego.

Un mapa es una colección de pares clave-valor.

```go
package main

import "fmt"

func main() {
	countries := make(map[string]string)
	fmt.Println(countries) // map[]


	countries["MX"] = "Mexico"
	countries["US"] = "Estados Unidos"
	countries["AR"] = "Argentina"

	fmt.Println(countries["MX"]) // Mexico
}
```

Esto sirve para poder guardar colecciones de datos en donde luego la búsqueda va a ser mas sencilla, a diferencia de un vector en donde solo tengo el indice para acceder a los datos, en los maps tengo la clave.

Puedo también declarar la cantidad de elementos que voy a tener en el mapa usando `make()`.

```go
countries := make(map[string]string, 10)
```

```go
package main

import "fmt"

func main() {

	campeonato := map[string]int{
		"Barcelona": 39,
		"Real":      38,
		"Atletico":  37
	}

	fmt.Println(campeonato) // map[Atletico:37 Barcelona:39 Real:38]

}
```

Para agregar un elemento al mapa:

```go
// Busca si el mapa tiene la clave, sino la agrega junto con el valor
campeonato["Colo-Colo"] = 36
```

Para modificar un elemento del mapa:

```go
campeonato["Barcelona"] = 40
// Lo modificamos usando la clave
```

Para eliminar un elemento del mapa:

```go
// delete() recibe como argumento el mapa y la clave para eliminar el par
delete(campeonato, "Barcelona")
```

Para recorre un mapa:

```go
for equipo, puntaje := range campeonato {
	fmt.Println("El equipo %s tiene un puntaje de %d", equipo, puntaje)
}
```

Para recorrer una mapa podemos usar un ciclo `for` y el comando `range` el cual nos devuelve ambos valores.

Para saber si un elemento existe en un mapa:

```go
puntaje, existe := campeonato["Mineirao"]
fmt.Println("El puntaje capturado es %d, y el equipo existe %t", puntaje, existe) // El puntaje capturado es 0, y el equipo existe false
```

## Estructuras (POO en Go)

En Go no existe la programación orientada a objetos como la conocemos, en donde tenemos una clase con sus métodos, etc. no la herencia y el polimorfismo no funciona igual a como la conocemos en Java o algún otro lenguaje.

Go implementa una nueva forma de hacer la programación orientada a objetos, para eso se basa en las `struct`.

```go
package main

import (
	"fmt"
	"time"
)

type user struct {
	Id 			int
	Nombre 		string
	FechaAlta 	time.Time
	Status 		bool
}

func main() {

	User := new(user)
	User.Id = 1
	User.Nombre = "Juan"
	User.FechaAlta = time.Now()
	User.Status = true
	fmt.Println(User)

}
```

En Go no existe los métodos como tal, sino que debemos usar una función que tenga una variable que apunte a una estructura, de esa forma cada vez que mencionemos esa variable estaremos haciendo referencia directamente a esa estructura.

```go
// Creamos un nuevo archivo llamado user.go
package user

import "time"

type User struct {
	Id 			int
	Nombre 		string
	FechaAlta 	time.Time
	Status 		bool
}

func (usr *User) setUser(id int, nombre string, fechaAlta time.Time, status bool) {
	usr.Id = id
	usr.Nombre = nombre
	usr.FechaAlta = fechaAlta
	usr.Status = status
}

func (usr *User) getUser() {
	fmt.Println("Id: ", usr.Id)
	fmt.Println("Nombre: ", usr.Nombre)
	fmt.Println("Fecha de alta: ", usr.FechaAlta)
	fmt.Println("Status: ", usr.Status)
}

func main() {

	User := new(User)
	User.setUser(1, "Juan", time.Now(), true)
	User.getUser()

}
```

De esta forma podemos crear una especie de métodos

También podemos hacer una especie de herencia de la siguiente forma:

```go
package main

import (
	"fmt"
	"time"

	us "./user"
)

type pepe struct {
	us.User
}

func main() {
	user := new(pepe)
	user.setUser(1, "Juan", time.Now(), true)
	fmt.Println(user.User)
}
```

De esta forma la `struct` `pepe` hereda el método `setUser` de la `struct` `User`.

También es importante mencionar la forma en la que importamos paquetes en Go, en este caso lo hacemos con:

```go
import(
	us "./user"
)
```

Al importar el paquete `user` bajo el nombre `us` podemos acceder a todos los métodos de la `struct` y propiedades de `User`, lo que estamos haciendo es importar la carpeta `user`, go supone que dentro del paquete `user` hay un archivo llamado `user.go`, si lo nombrásemos de otra forma no funcionaría.

## Interfaces

Las interfaces nos permiten definir conductas, operaciones, comportamientos.

Si lo llevamos al mundo real, podemos decir que los humanos tenemos ciertos comportamientos, por ejemplo, el pensar, por lo que si creamos un objeto y decimos que tiene la capacidad de pensar, estaríamos diciendo que ese objeto implementa la interfaz "humano".

Las interfaces se tratan de crear funciones que definen que un objeto tenga una característica determinada, por ejemplo, puedo tener una estructura llamada perro, una llamada gato, las cuales no tienen nada que las una, pero puedo implementar una interfaz que sea del tipo "mamífero" la cual tenga ciertas acciones que son comunes a los perros y los gatos como comer, correr, etc. y puedo también tener funciones que devuelvan valores que son comunes a los perros y los gatos como por ejemplo una función que se llame cantidad de patas y me devuelva la cantidad de patas que tenga cada una de las especies.

```go
package main

import "fmt"

type humano interface {
	respirar()
	comer()
	pensar()
	sexo() string
}

type hombre struct {
	edad int
	altura float32
	peso float32
	respirando bool
	comiendo bool
}

type mujer struct {
	edad int 
	altura float32
	peso float32
	respirando bool
	comiendo bool
}

func (h *hombre) respirar()    { h.respirando = true }
func (h *hombre) comer()       { h.comiendo = true }
func (h *hombre) pensar()      { h.pensando = true }
func (h *hombre) sexo() string { return "Hombre" }

func (m *mujer) respirar()    { m.respirando = true }
func (m *mujer) comer()       { m.comiendo = true }
func (m *mujer) pensar()      { m.pensando = true }
func (m *mujer) sexo() string { return "Mujer" }

func HumanoRespirando(hu humano) {
	hu.respirar()
	fmt.Printf("Soy un/a %s, y estoy respirando \n", hu.sexo())
}

func main() {
	Pedro := new(hombre)
	HumanoRespirando(Pedro)

	Maria := new(mujer)
	HumanoRespirando(Maria)
}
```

En la `interface humano` definimos métodos que son comunes a todos los humanos, por ejemplo, respirar, comer, pensar, etc.

Con las `structs` hombre y mujer definimos las estructuras que luego implementan la `interface humano`.

```go
func (h *hombre) respirar()    { h.respirando = true }
func (h *hombre) comer()       { h.comiendo = true }
func (h *hombre) pensar()      { h.pensando = true }
func (h *hombre) sexo() string { return "Hombre" }

func (m *mujer) respirar()    { m.respirando = true }
func (m *mujer) comer()       { m.comiendo = true }
func (m *mujer) pensar()      { m.pensando = true }
func (m *mujer) sexo() string { return "Mujer" }
```

No hace falta indicar de forma explicita que las estructuras están implementando la interfaz humano con alguna palabra reservada como `implements`, go se da cuenta solo de que es así.

Si nos fijamos, estamos repitiendo varios atributos que realidad son los mismo en hombre y mujer, es por eso que podemos simplemente heredar a mujer todas las propiedades de hombre.

```go
package main

type hombre struct {
	edad int
	altura float32
	peso float32
	respirando bool
	comiendo bool
	esHombre bool
}

type mujer struct {
	hombre
}
```

No es necearía entonces implementar la interface para mujer, ya que la hereda de hombre.

Cambiamos un poco la función ``sexo()``:

```go
func (h *hombre) sexo() string {
	if h.esHombre {
		return "Hombre"
	}
	return "Mujer"
}
```

Y se debería comportar de la misma manera:

```go
func main() {
	Juan := new(hombre)
	Juan.esHombre = true
	fmt.Println("Juan")
	HumanoRespirando(Juan)

	Maria := new(mujer)
	Maria.esHombre = false
	fmt.Println("Maria")
	HumanoRespirando(Maria)
}
```

Lo que me permiten las interfaces es relacionar objetos del mismo tipo, esto me permite crear polimorfismo en mi programa.

Por ejemplo, puedo crear una interfaz aún mas abstrcta que englobe a humano y a animal, la interfaz **SerVivo**:

```go
package main

import "fmt"

type SerVivo interface {
	estaVivo() bool
}

type humano interface {
	respirar()
	comer()
	pensar()
	sexo() string
	estaVivo() bool
}

type animal interface {
	respirar()
	comer()
	EsCarnivoro() bool
	estaVivo() bool
}

type hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	esHombre   bool
	vivo       bool
}

type perro struct {
	respirando bool
	comiendo   bool
	carnivoro  bool
	vivo       bool
}

func (h *hombre) estaVivo() bool { return h.vivo }

func (p *perro) estaVivo() bool { return p.vivo }

func estoyVivo(sv SerVivo) bool {
	return sv.estaVivo()
}

func main() {
	totalCarnivoros := 0
	Dogo := new(perro)
	Dogo.carnivoro = true
	Dogo.vivo = true
	AnimalRespirando(Dogo)
	totalCarnivoros += AnimalesCarnivoros(Dogo)
	fmt.Printf("Total de animales carnivoros: %d \n", totalCarnivoros)

	fmt.Printf("Estoy vivo: %t \n", estoyVivo(Dogo))

}
```

## Manejo de archivos en Go

Tenemos diferentes maneras de leer y escribir archivos en Go.

Las principales bibliotecas que podemos usar para esto son:

```go
import (
	"bufio"
	"io/ioutil"
	"os"
)
```

Una forma de leer es usando `ioututil.ReadFile`:

```go
func leoArchivo() {
	archivo, err := ioutil.ReadFile("./archivo.txt")
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		fmt.Println(string(archivo))
	}
}
```

Las funciones del sistema siempre regresan dos valores, es este caso es un `archivo` y también puede regresar un `error`, es por eso que siempre debemos escribir:

```go
if err != nil {
	fmt.Println("Hubo un error")
}
```

Con lo que definimos un comportamiento si es que hay un error, porque ``err`` siempre debería ser nulo, por ende, si es diferente de nulo entonces...

En go nulo se representa con `nil`.

Otra forma de leer es usando `os.Open`:

```go
func leoArchivo1() {
	archivo, err := os.Open("./archivo.txt")
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		scanner := bufio.NewScanner(archivo)
		for scanner.Scan() {
			registro := scanner.Text()
			fmt.Printf("Linea > " + registro + "\n")
		}
	}
	archivo.Close()

}
```

Luego con un `for` escaneamos linea por linea el archivo, esto nos puede ser util se estamos buscando algo en especifico, por que si lo encontramos podemos actuar en consecuencia.

Para escribir en el archivo podemos usar `os.Create`:

```go
func graboArchivo() {
	archivo, err := os.Create("./nuevoArchivo.txt")
	if err != nil {
		fmt.Println("Hubo un error 1")
		return
	}
	fmt.Fprintln(archivo, "Esto es un nuevo archivo")
	archivo.Close()
}
```

Otra forma es:

```go
func graboArchivo1() {
	fileName := "./nuevoArchivo.txt"
	if !Append(fileName, "\n Esta es una segunda linea") {
		fmt.Println("[graboArchivo1] Hubo un error")
	}
}

func Append(archivo string, texto string) bool {
	arch, err := os.OpenFile(archivo, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println("[Append] Hubo un error")
		return false
	}

	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("[WriteString] Hubo un error 3")
		return false
	}

	return true
}
```

## Recursividad en Go

La recursividad es básicamente una función que se llama a sí misma, esto se hace para que la función se ejecute un número indeterminado de veces, a menos que especifiquemos una condición bajo la cual finalize.

```go
package main

import "fmt"

func main() {
	exponencia(2)
}

func exponencia(numero int) {
	if numero > 100 {
		return
	}
	fmt.Println(numero)
	exponencia(numero * 2)
}
```

De esta forma la función recibe inicialmente un 2, pasa por la condición y luego se llama a sí misma, ese 2 que inicialmente le pasamos lo multiplica por 2 y se ejecuta nuevamente, solo que ahora ``numero`` vale 4, luego valdrá 8 y así sucesivamente.

## Excepciones y manejo de errores (defer, panic, recover)

A diferencia de otros lenguajes, aquí no hay nada como un `try/catch` o un `exception`.

El `defer` es una instrucción que se va a ejecutar si o si cuando se detecta que una función se va por un `return`, por un error o por un fin de función.

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	archivo := "prueba.txt"
	f, err := os.OpenFile(archivo)

	defer f.Close()

	if err != nil {
		fmt.Println("Hubo un error al abrir el archivo")
		os.Exit(1)
	}

}
```

Cuando coloco una instrucción antecedida por `defer` esa instrucción no se ejecuta secuencialmente, se va a ejecutar recién cuando yo salga de la función.

Lo que estoy diciendo con:

```go
defer f.Close()
```

Es que antes de salir de la función, ya sea por un error o porque finalizó, que cierre el archivo.

### Panic y Recover

`Panic` es una forma de forzar un error el cual hace que el sistema aborte la ejecución del programa.

```go
func ejemploPanic() {
	a := 1
	if a == 1 {
		panic("Panic")
	}
}
```

Esto se hace cuando por validación encontramos que hay un dato que nos falta que es crucial y que no puede continuar el programa sin ese dato, para eso se usa `panic`.

Panic no me da chance de continuar con el programa, ya que aborta la ejecución, para poder tener el control después de un `panic` debo usar `recover`:

```go
func ejemploPanic() {
	defer func() {
		reco := recover()

		if reco != nil {
			log.Fatalf("Ocurrio un error que generó un panic: \n %v", reco)
		}
	}()

	a := 1
	if a == 1 {
		panic("Panic")
	}
}
```

Lo que hace `recover` es que si detecta un `panic` trae el resultado de `panic`, osea que si no hubo `panic` `reco` debería ser `nil`, si hubo `panic` `reco` debería ser un string con el mensaje de error.

Es por eso que podemos validar esto:

```go
if reco != nil {
	fmt.Println(reco)
}
```

## GoRoutines (ejecución asíncrona - Promesas en Go)

Las `GoRoutines` tiene que ver con el multi-threading y la programación asíncrona.

Si venimos de `NodeJS` sabemos que existen las promesas, los `async` y `await, que son la forma que tiene NodeJS de hacer que el código sea asíncrono. Sin embargo esta forma que tiene node no es tan clara de entender.

Go tiene una forma de hacer eso de una forma mucho mas amigable mediante las `GoRoutines`, las cuales son el equivalente a las promesas y a los `async` y `await` de NodeJS.

```go
package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	go miNombreLento("Diego")
	fmt.Println("Estoy en el main")
	var x string
	fmt.Scanln(&x)
}

func miNombreLento(nombre string) {
	letras := strings.Split(nombre, "") // Separo las letras del nombre
	for _, letra := range letra {
		time.Sleep(1000 * time.Millisecond)
		fmt.Printf("%s", letra)
	}
}
```

Con la instrucción `go` podemos crear una `GoRoutine` que se ejecute en paralelo con el código principal, es por eso que al ejecutar el código puedo para el programa antes de que se termine de mostrar el nombre.

En Go el runtime no se queda a esperar que termine la función asíncrona, mientras estoy ejecutando la función asíncrona, el sistema termina la ejecución y aborta el programa.

## Channels en Go (Diálogo entre GoRoutines)

Los channels son unos canales que permiten que una goroutine envíe información hacia otra función, ya sea main, otra goroutine, o cualquier otra función.

Esto sirve para que cada función paralela que se esté ejecutando pueda dialogar con otra para enviar información o para recibir información.

Un canal es un espacio de memoria en el que van a dialogar diferentes rutinas y cuando se aloje un valor en ese espacio de memoria, la rutina que está pidiendo un valor va a actuar en consecuencia. Ese espacio de memoria tiene que se de un tipo de dato, en definitiva lo que va a transportar el canal es una variable de ese tipo de dato. 

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan time.Duration)
	go bucle(canal)
	fmt.Println("llegue hasta aquí")

	msg := <-canal // Espero que alguien me envíe un mensaje
}

func bucle(canal chan time.Duration) {
	init := time.Now()
	for i := 0; i < 100; i++ {
	}

	final := time.Now()
	canal <- final.Sub(init)
}
```

Al escribir:

```go
canal := make(chan time.Duration)
```

Estoy creando una canal de tipo `time.Duration`, luego cuando le quiero pasar un dato a ese canal, debo escribir:

```go
init := time.Now()
for i := 0; i < 100; i++ {
}

final := time.Now()
canal <- final.Sub(init)
```

Con `<-` le paso un valor al canal, y luego para recibirlo, debo escribir:

```go
msg := <-canal
```

Des esta forma creo algo similar al `asyc-await` de NodeJS.

## Servidor Web en Go

Go cuenta con un paquete llamado `net/http` que permite crear un servidor web.

```html
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Document</title>
</head>
<body>
	<h1>Hola mundo</h1>
</body>
</html>
```

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", home)
	http.ListenAndServe(":3000"	, nil)
	fmt.Println("Server running on port 3000")

}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
```
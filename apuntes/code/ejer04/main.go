package main

import (
	"bufio"
	"fmt"
	"os"
)

var numero int
var numero1 int
var res int
var leyenda string

func main() {
	fmt.Println("Introduce un numero")
	fmt.Scanln(&numero)

	fmt.Println("Ingrese otro numero")
	fmt.Scanln(&numero1)

	fmt.Println("Descripcion: ")

	scan := bufio.NewScanner(os.Stdin)

	if scan.Scan() {
		leyenda = scan.Text()
	}

	res = numero + numero1
	fmt.Println(leyenda, res)
}

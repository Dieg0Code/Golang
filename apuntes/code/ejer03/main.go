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

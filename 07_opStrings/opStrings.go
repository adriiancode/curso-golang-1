package main

import "fmt"

func main() {
	var cadena string = "hola \"programadores\""

	//imprime toda la cadena
	fmt.Println(cadena[:])

	//[incluye:excluye]
	fmt.Println(cadena[1:3])

	//imprime hasta indice menor que 3 (2)
	fmt.Println(cadena[:3])
}

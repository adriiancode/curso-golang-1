package main

import (
	"fmt"
)

func main() {
	//unidimensional
	var arreglo [3]int
	fmt.Println(arreglo)

	//bidimensional
	var edades [3][2]int
	fmt.Println(edades)

	//asignar un valor
	arreglo[1] = 23
	fmt.Println(arreglo)

	//declarar e inicializar

	arreglo2 := [5]int{1, 3, 5, 7, 9}
	fmt.Println(arreglo2)

	//tamaño de acuerdo a valores que inicialize
	arreglo3 := [...]int{3, 2, 3}
	fmt.Println(len(arreglo3))

	//comparacion de arrays
	/* para ser iguales deben tener:
	- el mismo tamaño
	- el mismo tipo
	- el mismo contenido
	*/
	a := [3]int{1, 2, 3}
	b := [...]int{1, 2, 3}
	c := [3]int{1, 2, 4}
	d := [4]int{1, 2, 3}

	fmt.Println(a == b)
	fmt.Println(a == c)
	fmt.Println(a == d)

}

package main

import (
	"fmt"
)

func main() {
	/*un slice es similar a un array dinamico
	el tamaño puede cambiar en tiempo de
	ejecucion a diferencia del array*/
	slice := []int{1, 2, 3, 4}
	fmt.Println(slice)

	/* declaracion de slice usando make
	se pasan dos parametros, el tipo,
	el tamaño, y opcional la capacidad*/

	/*si no defino la capacidad tomara
	el mismo valor que el tamaño*/
	// j := make([]int, 5)
	// fmt.Println(j)
	// fmt.Println(len(j))
	// fmt.Println(cap(j))

	/*ejemplo con tamaño y cacidad*/
	// k := make([]int, 5, 10)
	// fmt.Println(k)
	// fmt.Println(len(k))
	// fmt.Println(cap(k))

	/*definimos un array*/
	ar := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(ar)

	/*cualquier cambio en un slice modificara
	el array original */
	a := ar[2:5]
	/*cuando se asigna de esta manera la capacidad
	de a es desde donde comieza a hasta donde termina ar*/
	fmt.Println(a)

	b := ar[3:5]
	fmt.Println(b)

	b[0] = 25
	fmt.Println(ar)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(cap(b))

}

/*un puntero es un tipo de vlor que contiene
la direccion en memoria de una variable*/

package main

import (
	"fmt"
)

func main() {
	a := 23
	fmt.Println("valor de a es : ", a)
	/*un punterose simboliza antecedientola variable por & */
	fmt.Println("direccion en memoria de a es : ", &a)

	b := &a
	/*ahora b apunta a la direccion de a */
	/*si una variable esta apuntando a una
	direccion existente,con * previo a la
	variable accedemos a su valor */
	fmt.Println("direccion de b: ", b)
	fmt.Println("valor de b: ", *b)
	// b = 45 //error, b es un puntero
	*b = 45
	fmt.Println(a)
	fmt.Println(*b)

	/*otra forma de llamar a un puntero es
	con la palabra clave new */
	d := new(int) // *int
	fmt.Println(d)
	*d = 11
	fmt.Println(*d)
}

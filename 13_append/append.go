package main

import (
	"fmt"
)

func main() {
	//declaramos slice x
	x := make([]byte, 4, 10)
	fmt.Println(x)

	//inicializamos slice x
	/*cuando colocamos una letra en comilla
	simple no se guarda la letra en si sino
	el numero de utf8 al que correspone el
	caracter colocado*/
	x = []byte{'H', 'O', 'L', 'A'}
	/*el marcador de posicion %q nos da la
	representacion del numero utf8 recibido*/
	fmt.Printf("slice x : %q \n", x)

	//imprimiendo cada elemnto del slice porseparado
	for i := 0; i < len(x); i++ {
		fmt.Printf("slice x[%d]: %q\n", i, x[i])
	}

	//si tratamos de asignar un espacio en blanco en el indice 5 del slice x
	// x[5] = ' ' //error
	/*no podemos acceder a indices mayores la tamaño
	del slice aun cuando tenga capacidad*/
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = append(x, ' ')
	fmt.Printf("slice x: %q", x)
	fmt.Printf("tamaño: %d - capacidad: %d \n", len(x), cap(x))

	fmt.Println("********")

	/*la funcion append duplica la capacidad del slice si
	el tamaño del mismo ya no es suficiente para guardar*/

	var y []int
	for i := 1; i <= 15; i++ {
		//agregamos el valor de i al slice y
		y = append(y, i)
		fmt.Println("slice y: ", y)
		fmt.Printf("Tam y: %d - Cap y: %d - Elem y: %d \n",
			len(y), cap(y), i)
	}
}

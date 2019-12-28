/* los maps en golang es lo que otros
lenguajes son diccionarios/objetos
con clave y valor*/

package main

import (
	"fmt"
)

func main() {

	//recibe el tipo de dato de la clave y del valor

	// x := make(map[string]string)
	// fmt.Println(x)

	/*se le puede pasar la capacidad tambien
	si se sabe cuanto contendra*/
	// y := make(map[string]string, 7)
	// fmt.Println(y)

	//asignar valores
	// x["nombre"] = "matias"
	// x["edad"] = "23"
	// fmt.Println(x)

	//acceder a los valores
	// fmt.Println(x["nombre"])

	edades := map[string]int{
		"ana":      23,
		"luis":     34,
		"maria":    21,
		"fernando": 22,
	}

	fmt.Println(edades)

	//eliminar elementos de un map pasandole su clave
	delete(edades, "ana")
	fmt.Println(edades)

	//iterar un map
	for index, edad := range edades {
		fmt.Printf("la edad de %s es %d \n", index, edad)
	}
}

package main

import (
	"fmt"
)

func main() {
	nombres := []string{
		"luis",
		"juan",
		"maria",
		"ana",
		"pedro",
	}

	//iterar un slice sin saber su tamaño con range()
	// for index, elemento := range nombres {
	// 	fmt.Printf("nombre %q esta en indice %d \n", elemento, index)
	// }

	/*el indice se puede reemplazar con guion bajo
	para que el compilador lo omita ej: for _,elmento*/
	// for _, elemento := range nombres {
	// 	fmt.Println(elemento)
	// }

	/*recibe solo el indice*/
	for index := range nombres {
		fmt.Println(index)
	}

	cadena := "hola mi gente, soy adrian"

	for index, letra := range cadena {
		fmt.Printf("la letra es %q en el indice %d\n", letra, index)
	}

}

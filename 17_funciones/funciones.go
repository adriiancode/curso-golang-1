package main

import "fmt"

//funcion que no retorna nada
// func imprimirNombre(nombre string) {
// 	fmt.Println("Fuera de main")
// 	fmt.Println("el nombre es: ", nombre)
// }

//funcion que si retona
// func sumar(a, b int) int {
// 	return a + b
// }

//otra formade de que retorne
// func resta(x, y int) (r int) {
// 	r = x - y
// 	return
// }

/*variadic funcionts son aquellas que
reciben un numero indefinido de parametros*/

func sumarVarios(numeros ...int) int {
	resultado := 0
	for _, numero := range numeros {
		// resultado = resultado + numero
		resultado += numero
	}
	return resultado
}

//el elype (...) siempre va en el parametro final
func sumarCadenas(cadena string, cadenas ...string) {
	for _, c := range cadenas {
		cadena += " "
		cadena += c
	}

	fmt.Println(cadena)
}

func main() {
	// imprimirNombre("juan")
	// fmt.Println(sumar(2, 3))
	// fmt.Println(resta(4, 2))
	// fmt.Println(sumarVarios(4, 2))               //6
	// fmt.Println(sumarVarios(4, 2, 4))            //10
	// fmt.Println(sumarVarios(10, 10, 10, 10, 10)) //50

	// fmt.Println(" ")

	// numeros := []int{
	// 	12,
	// 	6,
	// 	23,
	// }

	/*dara error porque recibe enteros no un slice de enteros*/
	// fmt.Println(sumarVarios(numeros)) //error

	/*el elipse (...) al final indica que omita
	el tipo slice y que se cncentre en el contenido*/
	// fmt.Println(sumarVarios(numeros...))
	sumarCadenas("hola", "mundo", "de", "programadores")

}

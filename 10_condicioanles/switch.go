package main

import (
	"fmt"
)

func main() {
	// nombre := "luis"

	//mala practica
	// if nombre == "hernan" {
	// 	fmt.Println("soy hernan")
	// } else if nombre == "maria" {
	// 	fmt.Println("soy maria")
	// } else if nombre == "ana" {
	// 	fmt.Println("soy ana")
	// } else {
	// 	fmt.Println("soy otra persona")
	// }

	//buena practica
	// switch nombre {
	// case "maria":
	// 	fmt.Println("soy maria")
	// case "juan":
	// 	fmt.Println("soy juan")
	// case "luis":
	// 	fmt.Println("soy luis")
	// default:
	// 	fmt.Println("soy otra persona")
	// }

	/* a diferencia de otros lenguajes
	el sitch en go permite colocar condicionales
	dentro del mismo case, y el break al terminar
	cada caso no es necesario*/
	edad := 12
	apellido := "martinez"
	switch {
	case edad == 10:
		fmt.Println("es menor de 10")
	case edad == 12:
		fmt.Println("la edad es 12")
		fallthrough
	case apellido == "martinez":
		fmt.Println("apellido es martinez")
	default:
		fmt.Println("algo paso")
	}

	/* con la palabra clave fallthrogh
	alfinal de un caso,le decimos al
	programa que siga evaluando los otros
	casos posteriores aun cuando ya se haya
	cumplido la condicion de un caso previo*/
}

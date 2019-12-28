package main

import "fmt"

type Persona struct {
	Nombre string
	Edad   int
}

//older determina quien es mayor y la diferencia de edad
func older(p1, p2 Persona) (string, int) {
	if p1.Edad > p2.Edad {
		return p1.Nombre, p1.Edad - p2.Edad
	}
	return p2.Nombre, p2.Edad - p1.Edad
}
func main() {

	//inicializamos 3 ariables de tipo Persona
	tom := Persona{"tom", 33}
	paulina := Persona{"paulina", 23}
	miguel := Persona{"miguel", 27}

	//llamamos a la funcion older

	t1Older, t1Diff := older(tom, paulina)
	t2Older, t2Diff := older(tom, miguel)
	t3Older, t3Diff := older(miguel, paulina)

	//imprimimos los resultados
	fmt.Printf("el mas viejo entre tom y paulina es: %s por %d años \n", t1Older, t1Diff)
	fmt.Printf("el mas viejo entre tom y miguel es: %s por %d años \n", t2Older, t2Diff)
	fmt.Printf("el mas viejo entre paulina y miguel es: %s por %d años \n", t3Older, t3Diff)
}

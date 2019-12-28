package main

import "fmt"

type Persona struct {
	Nombre string
	Edad   int
}

//aqui declaramos un campo anonimo
type Estudiante struct {
	Persona
	Carrera string
}

type Profesor struct {
	Estudiante
	Carrera string
}

func main() {
	alejandro := Estudiante{
		Persona{
			Nombre: "alejandro",
			Edad:   23,
		},
		"informatica",
	}

	fmt.Println(alejandro)
	fmt.Println(alejandro.Nombre)
	fmt.Println(alejandro.Edad)
	fmt.Println(alejandro.Carrera)

	/*en la declraacion multilinea todos los elementos
	llevan coma al final, hasta el ultimo*/
	pedro := Profesor{
		Estudiante{
			Persona{
				Nombre: "pedro",
				Edad:   45,
			},
			"informatica",
		},
		"contabilidad",
	}

	fmt.Println(pedro)
	fmt.Println(pedro.Carrera)
	fmt.Println(pedro.Estudiante.Carrera)

}

package main

import (
	"fmt"
)

type Persona struct {
	Nombre string
	Edad   int
}

func main() {
	var p Persona
	p.Nombre = "camilo"
	p.Edad = 32
	fmt.Println("Estructura p ", p)

	fmt.Println(p.Nombre)
	fmt.Println(p.Edad)

	//declarar e inicializar una estructura
	p2 := Persona{Nombre: "Rafael", Edad: 23}
	fmt.Println(p2)

	//otra forma si nos sabemos lo que lleva
	p3 := Persona{"lucas", 22}
	fmt.Println(p3)
}

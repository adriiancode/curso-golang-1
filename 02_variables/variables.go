package main

import "fmt"

func main() {
	//declaracion e inizializacion por separado
	var nombre string
	nombre = "juan"
	//se puede omtir el tipo
	var nombre2 = "roberto"
	fmt.Println(nombre)

	//declaracion e inicializacion a la vez
	//solo es posible hacer esto dentro de funciones o metodos
	//solo funciona si la variable es nueva
	edad := 34

	// edad := 54 //error
	// edad := 21 //error
	fmt.Println(edad)

	//declaracion e inicializacion multiple
	//se puede omitir el tipo
	var (
		ciudad string = "Riohacha"
		piso   int    = 1
	)
	fmt.Println(ciudad, piso)

	//otra forma de declaracion e inicializacion multiple
	apellido, anio, escasada := "martinez", 1983, true
	fmt.Println(apellido, anio, escasada)

}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	encabezado := `
// 	***************************
// 	Contador de numeros impares
// 	***************************
// 	`
// 	fmt.Println(encabezado)
// 	fmt.Println("Digita el numero donde quieres que inicie")
// 	var numero1 int
// 	fmt.Scanln(&numero1)
// 	fmt.Println("Digita el numero hasta donde quieres que termine")
// 	var numero2 int
// 	fmt.Scanln(&numero2)
// 	fmt.Println("ok...")
// 	for numero1 < numero2 {
// 		if (numero1%2)!=0 {
// 			fmt.Printf("%d es impar \n",numero1)
// 		}
// 		numero1++
// 	}
// 	encabezado = `
// 	****************
// 	Fin del Programa
// 	****************
// 	`
// 	fmt.Println(encabezado)
// 	encabezado = `
// 	********************
// 	Author: Adrian Gomez
// 	********************
// 	`
// 	fmt.Println(encabezado)

// }
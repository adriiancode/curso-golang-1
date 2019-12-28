package main

import "fmt"

func main() {
	a := 10
	b := 3
	/*
		Operadores aritmeticos
		+ suma
		- resta
		* multiplicacion
		/ division
		% residuo
+= aumento
-= decremento
++ aumento en uno
-- decremento en uno
	*/

	fmt.Println("10 + 2 es ", a+b)
	fmt.Println("10 - 2 es ", a-b)
	fmt.Println("10 * 2 es ", a*b)
	fmt.Println("10 / 2 es ", a/b)
	fmt.Println("10 % 2 es ", a%b)

	/*
		Operadores de comparacion
		== igual que
		!= diferente que
		< menor que
		<= menor o igual que
		< mayor que
		>= mayor o igual que
	*/

	fmt.Println("5 es igual a 5 ?", 5 == 5)
	fmt.Println("10 es diferente de 2?", 10 != 2)
	fmt.Println("5 es menor que 4 ?", 5 < 4)
	fmt.Println("7 es menor o igual que 7 ?", 7 <= 7)
	fmt.Println("5 es mayor que 4 ?", 5 > 4)
	fmt.Println("13 es mayor o igual que 13 ?", 13 >= 13)

	/*
		Operadores logicos:
		&& and (verdadero si todas las condiciones se cumplen)
		|| or (verdadero si una de las condiciones se cumplen)

	*/
	fmt.Println(2 > 1 && 4 > 3)
	fmt.Println(10 > 5 || 3 < 1)
	fmt.Println(10 < 5 || 3 < 1)
}

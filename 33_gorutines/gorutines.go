package main

import (
	"fmt"
	"time"
)

func fib(x int) int {
	if x < 2 {
		return x
	}
	//usando recursividad para llamarse asi misma
	return fib(x-1) + fib(x-2)

}

func animacion(retraso time.Duration) {
	for {
		for _, r := range `\|/-` {
			fmt.Printf("\r %c", r)
			time.Sleep(retraso)
		}
	}
}

/*una recursion sucede cuando una funcion se llama
a si misma haciendo que se produsca un bucle */
func main() {
	go animacion(100 * time.Millisecond)
	/*en este caso no se utiliza la estructura waitgroup
	para que el programa espere a que finalicen las
	gorutinas*/
	const n = 45
	resultado := fib(n)
	fmt.Printf("\r Fibonacci(%d) = %d\n", n, resultado)
}

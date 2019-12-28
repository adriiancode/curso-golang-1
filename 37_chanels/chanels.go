/*refactorizacion de codigo anterior y cambio de
chanels bidireccionales(pueden recibir y enviar)
 a undireccional(solor pueden recibir o mandar) chanels*/
package main

import (
	"fmt"
	"time"
)

/*
<-chan canal de entrada
chan<- canal de salida
*/

//generamos un ciclo y mandamos el indice por el canal
func generarNumeros(out chan<- int) {
	//enviamos 5 datos
	for i := 1; i <= 5; i++ {
		out <- i
	}
	close(out)
}

func elevarAlCuadrado(in <-chan int, out chan<- int) {
	for i := range in {
		out <- i * i
	}
	//los canales que solo reciben no hay que cerrarlos
	close(out)
}

func imprimir(in <-chan int) {
	for valor := range in {

		/*para no estar comprobando cada vez si  ya se cerro
		el canal hacemos un for range recorriendo el canal*/
		fmt.Println("cuadrado : ", valor)
		time.Sleep(1 * time.Second)
	}
}
func main() {
	numero := make(chan int)
	cuadrado := make(chan int)

	go generarNumeros(numero)
	go elevarAlCuadrado(numero, cuadrado)
	imprimir(cuadrado)

}

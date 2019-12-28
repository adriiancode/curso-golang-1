package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func imprimir(letra string) {
	//decrementa el contador de waitgroup en 1 al final por el defer
	defer wg.Done()

	for cantidad := 1; cantidad <= 10; cantidad++ {

		/*obtengo un numero  entero al azar entre 0 y 2000*/
		tiempoEspera := rand.Int63n(3000)

		/*dengo el programa durante el tiempo random previo*/
		time.Sleep(time.Duration(tiempoEspera) * time.Millisecond)

		fmt.Printf("letra %s - vuelta #%d \n", letra, cantidad)
	}
}

var wg sync.WaitGroup

func main() {

	/*ejecucion lineal (sin go rutinas)*/
	// saludar()
	// despedirse()

	/*la funcion Add del paquete sync especifica el numero de
	go rutinas que el waitgroup tendra*/
	wg.Add(2)
	fmt.Println("Iniciamos la go rutinas")

	/* al anteponer la pabra go a la funcion le decimos al compilador
	que se trata de una gorutina y que no depende de la ejecucion
	del programa principal main, entonces el programa las ejecuta de
	manera concurrente (similar a en paralelo pero no igual)
	*/
	go imprimir("A")
	go imprimir("B")

	fmt.Println("Esperando que finalicen")

	/*la funcion wait del paquete sync
	wg hace que el programa espere el numero
	de go rutinas especificados en wg.Add*/
	wg.Wait()
	fmt.Println("Fin del programa")

}

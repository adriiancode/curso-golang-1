/*los canales son un medio de comunicacion
entre dos gorutinas. Una de sus particulalidades
es que la gorutina receptora siempre espera a tener
el mensaje de la gorutina emisora*/

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	//cremos el canal para pasar strings
	ch := make(chan string)

	wg.Add(2)

	//llamamos las funciones como gorutinas
	go enviarPing(ch)
	go imprimirPing(ch)

	wg.Wait()
	fmt.Println("FIn del programa")

}

func enviarPing(c chan string) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		//enviamos la palabra ping 10 veces por el canal
		c <- "Ping"
	}
}
func imprimirPing(c chan string) {
	defer wg.Done()
	var contador int
	for i := 0; i < 10; i++ {
		contador++
		//recibiendo datos por el canal
		fmt.Println(<-c, " ", contador)

		/*hace que la funcion espere 1 minuto para poder
		observar el funcionamiento*/
		time.Sleep(1 * time.Second)
	}

}

/*Existen dos tipos de canales: los unbufered chanels
que pueden mandar y recibir un solo valor y los
bufered chanels que pueden mandar y recibir varios.
En este ejercisio usamos los unbufered*/

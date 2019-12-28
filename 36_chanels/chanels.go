package main

import (
	"fmt"
	"time"
)

func main() {
	numero := make(chan int)
	cuadrado := make(chan int)

	//generamos un ciclo y mandamos el indice por el canal
	go func() {
		//ciclo infinito con contador
		// for i := 0; ; i++ {
		// 	numero <- i
		// }

		//enviamos 5 datos
		for i := 1; i <= 5; i++ {
			numero <- i
		}
		/*la funcion close cierra un canal cuando ya termino
		de enviar o recibir datos*/
		close(numero)
	}()

	go func() {
		for i := range numero {
			/*cuando los chanel reciben informacion retornan
			no solo el dato que mandaron si no un booleano que
			indica si el canal esta abierto*/
			// i, ok := <-numero
			// if !ok {
			// 	//para salir del ciclo
			// 	break
			// }
			/*para no estar comprobando cada vez si  ya se cerro
			el canal hacemos un for range recorriendo el canal*/
			cuadrado <- i * i
		}
		//cierro el canal porque no le voy a mandar mas datos
		close(cuadrado)
	}()

	/*los canales tambien sirven para sincronizar gorutinas
	ya que los canales hacen que el programa main espere por el dato que viene de la gorutina, sin necesidad de utilizar el paquete sync*/

	for valor := range cuadrado {
		// valor, ok := <-cuadrado
		// if !ok {
		// 	break
		// }
		/*para no estar comprobando cada vez si  ya se cerro
		el canal hacemos un for range recorriendo el canal*/
		fmt.Println("cuadrado : ", valor)
		time.Sleep(1 * time.Second)
	}
}

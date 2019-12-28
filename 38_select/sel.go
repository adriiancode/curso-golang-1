package main

import (
	"fmt"
	"os"
	"time"
)

func leerEntrada(out chan<- []byte) {
	for {
		datos := make([]byte, 1024)
		n, _ := os.Stdin.Read(datos)
		if n > 0 {
			out <- datos
		}
	}
}

func main() {
	//creamos un chanel de manera diferente
	done := time.After(20 * time.Second)

	//creamos chanel de manera tardicional
	eco := make(chan []byte)
	go leerEntrada(eco)
	for {
		/*un select hace lo mismo que un switch pero en vez
		de evaluar una expresion para ver si es igual a otra
		lo que hace es que evalua cuando atravaes de un canal
		vienen datos*/
		select {
		//si vienen datos atraves del canal eco
		case datos := <-eco:
			os.Stdout.Write(datos)
		//si vienen datos atraves del canal done
		case <-done:
			fmt.Println("Se terminó el tiempo")
			os.Exit(0)
		}
	}

}

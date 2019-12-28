/*funciones que no tienes nombre definido
y solo son definas cuando se necesitan*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	cadena := "hola mundo"

	/*la funcion Map del paquete strings recorre una
	cadena obteniendo el valor al que corresponde
	dicha letra en utf8 y recibe una funcion (anonima en este caso) como parametro para procesar esos datos utf8*/
	nuevaCadena := strings.Map(func(r rune) rune {
		return r + 1
	}, cadena)
	fmt.Println(nuevaCadena)
}

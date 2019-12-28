/*en un programa simepre puede haber
errores , como golang es un lenguaje
compilado podemos detectar muchos de ellos
en tiempo de compilacion pero inevitablemente
algunos pueden aparecer solo en tiempo de
ejecucion*/
/*cuando suceden estos errores golang
manda un panic. Nosotros tambien podemos
crear nuestros propios panic por ejemplo
cuando la conexion a la base de datos no
sucede el resto del programa no tiene sentido
y hay que mandar un panic*/
package main

import (
	"fmt"
)

func imprimir() {
	fmt.Println("hola gophers")
	defer func() {
		/*recover recupera lo que sea que haya
		mandado panic*/
		cadena := recover()
		fmt.Println(cadena)
	}()

	panic("error")
}
func main() {
	imprimir()
	fmt.Println("hola main")
}

/*golang permite declarar nuestros propios
tipos en base a otros ya existentes */
package main

import (
	"fmt"
)

/*la ventaja de utilizar tipos personalizados
es que atraves de metodos podemos cambiar el
tipo de fondo*/
type Dinero int

func (d Dinero) String() string {
	return fmt.Sprintf("$%d", d)
}

func main() {
	var sueldo Dinero
	sueldo = 24000
	fmt.Println("sueldo ", sueldo)

	sueldo = 20000
	fmt.Println("sueldo ", sueldo)

	aumento := 11000
	// sueldo += aumento //error de tipos
	sueldo += Dinero(aumento)

	fmt.Println("sueldo + aumento: ", sueldo)
}

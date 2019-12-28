package main

import (
	"fmt"
	"os"
)

/*la palabra clave defer ejecuta si o si
lo que tenga al lado pero al final
de la funcion main aun cuando ha ocurrido
un error*/
func main() {
	/* muy util cuando se aben recursos
	externos del sistema como archivos o
	conexion a base de datos para cerrar
	el archivo o la conexion*/

	//otros ejemplo
	//abrimos un archivo
	file, err := os.Open("texto.txt")

	/*creamos un slice para alacenar lo que
	leamos del archivo*/
	data := make([]byte, 55)
	c, err := file.Read(data)
	//verificamos que no haya ocurrido ningun error
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//imprimimos lo leido
	fmt.Printf("Cantidad de bytes leidos: %d \n %q \n error: %v", c, data, err)

}

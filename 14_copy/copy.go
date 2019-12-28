/*copy nos permite copiar el contenido
de un slice en otro slice*/

package main

import (
	"fmt"
)

func main() {

	// origen == destino
	// origen := []byte{1, 2, 3}
	// destino := []byte{3, 4, 5}
	/* copy recibe dos parametros:
	a donde se copiara, y desde donde*/
	/*en este caso el slice destino ya tiene
	contenido dentro pero es reemplazdo*/
	// copy(destino, origen)
	// fmt.Println(origen, destino)

	/*si el slice origen > destino
	copia desde el principio de origen
	hasta el tamaño de destino*/
	// origen2 := []int{1, 2, 3}
	// destino2 := make([]int, 2)
	// copy(destino2, origen2)
	// fmt.Println(origen2, destino2)

	/* si el tamaño del slice origen es
	menor que el tamaño del slice destino
	la funcion copy reemplazara solo algunos
	indices de destino*/
	origen3 := []int{1, 2, 3}
	destino3 := []int{3, 4, 5, 6, 7}
	copy(destino3, origen3)
	fmt.Println(origen3, destino3)
}

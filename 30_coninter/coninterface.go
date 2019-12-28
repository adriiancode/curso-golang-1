/*nos ayuda a aplicar una funcion a mas de un tipo de dato, en este caso para que
la funcion presetarse se pueda aplicar
no solo a personas si no a moderadores,
administradores,etc*/

package main

import (
	"fmt"
)

type Persona struct {
	Nombre, Email string
	Edad          int
}

func (p Persona) GetName() string {
	return p.Nombre
}

func (p Persona) GetEmail() string {
	return p.Email
}

type Moderador struct {
	Persona
	Foro string
}

func (m Moderador) AbirForo()  {
	fmt.Println("foro cerrado")
}

type Administrador struct {
	Persona
	seccion string
}

type Usuario interface {
	GetName() string
	GetEmail() string
}

/*en golang no hay clases, se trabaja por composicion*/
/*las interfaces nos permiten separar la definicion de la implementacion */
/*las interfaces me permiten compartir funciones para diferentes */

func presentarse(p Usuario) {
	fmt.Println("Nombre: ", p.GetName())
	fmt.Println("Email: ", p.GetEmail())
}

/*si no creo interface me toca repetir la funcion la cantidad de veces que necesite*/
/*las interfaces tambien son un tipo de dato , como las funciones y las estructuras,int,string,etc*/

func main() {
	alejandro := Persona{"Alejandro", "alejo@gmailcom", 23}
	presentarse(alejandro)

	juan := Moderador{Persona{"juan", "juan@gmailcom", 25}, "jugar en español"}

	presentarse(juan)

	maria := Administrador{Persona{"maria", "maria@gmailcom", 32}, "juegos"}

	presentarse(maria)

	/*usando la interface como tipo de dato*/
	var nuevo Usuario

	/*la intrpretacion logica es:
	- una interface da acceso a todos los metodos
	- los metodos estan asociados a un tipo
	-si un tipo contiene a otro que tiene los mismos metodos 
	tambien le comprtira dichos metodos
	*/
	nuevo = alejandro
	fmt.Println("nombre: ",nuevo.GetName())
	fmt.Println("email: ",nuevo.GetEmail())

	nuevo = juan
	fmt.Println("nombre: ",nuevo.GetName())
	fmt.Println("email: ",nuevo.GetEmail())
	/*nuevo es de tipo Usuario no de tipo Moderador
	por lo que desde nuevo no se puede acceder a los
	metodos de Moderador*/
	// fmt.Println("foro: ",nuevo.AbrirForo()) //error
}

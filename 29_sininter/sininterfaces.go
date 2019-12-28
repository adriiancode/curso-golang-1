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

func (m Moderador) AbrirForo() {
	fmt.Println("Abrir un foro")
}

func (m Moderador) CerrarForo() {
	fmt.Println("Cerar Foro")
}

type Administrador struct {
	Persona
	seccion string
}

func (a Administrador) CrearForo() {
	fmt.Println("foro creado")
}

func (a Administrador) EliminarForo() {
	fmt.Println("foro eliminado")
}

/*en golang no hay clases, se trabaja por composicion*/
/*las interfaces nos permiten separar la definicion de la implementacion */

func presentarse(p Persona) {
	fmt.Println("Nombre: ", p.GetName())
	fmt.Println("Email: ", p.GetEmail())
}

/*si no creo interface me toca repetir la funcion la cantidad de veces que necesite*/
func presentarseA(a Administrador) {
	fmt.Println("Nombre: ", a.GetName())
	fmt.Println("Email: ", a.GetEmail())
}

func presentarseM(m Moderador) {
	fmt.Println("Nombre: ", m.GetName())
	fmt.Println("Email: ", m.GetEmail())
}
func main() {
	alejandro := Persona{"Alejandro", "alejo@gmailcom", 23}
	presentarse(alejandro)

	juan := Moderador{Persona{"juan", "juan@gmailcom", 25}, "jugar en español"}

	presentarseM(juan)

	maria := Administrador{Persona{"maria", "maria@gmailcom", 32}, "juegos"}

	presentarseA(maria)

}

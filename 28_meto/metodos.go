package main

import "fmt"

type Rectangulo struct {
	alto, ancho float64
}

type Circulo struct {
	radio float64
}

/*los metodos siempre estan asociados a un tipo*/
func (r Rectangulo) area() float64 {
	return r.ancho * r.alto
}

/*aunque el metodo se llame igual no choca porque
esta asociado a otro tipo, ya que el ambito de los
metodos es el corchete del tipo al que pertenece*/
// func (c Circulo) area() float64 {
// 	return (c.radio * c.radio) * math.Pi
// }

/*cuando se va a sobreescribir los datos es mala
devolver un nuevo dato*/
// func (r Rectangulo) inc(i float64) Rectangulo {
// 	return Rectangulo{
// 		r.ancho * i,
// 		r.alto * i,
// 	}
// }

/*es mejor modificar el metodo para que el
recibidor sea un puntero y modifique el valor
internamente*/
func (r *Rectangulo) inc(i float64) {
	r.ancho *= i
	r.alto *= i
}

func main() {
	r1 := Rectangulo{12, 2}
	r2 := Rectangulo{15, 3}

	// fmt.Println("El area de r1 es :", r1.area())
	// fmt.Println("El area de r2 es :", r2.area())

	// c1 := Circulo{4}
	// c2 := Circulo{10}

	// fmt.Println(" El area de c1 es: ", c1.area())
	// fmt.Println(" El area de c2 es: ", c2.area())

	fmt.Println("El area de r1 es :", r1.area())
	fmt.Println("El area de r2 es :", r2.area())

	//mala practica
	// r1 = r1.inc(10)
	// r2 = r2.inc(10)

	//bunea practica
	r1.inc(10)
	r2.inc(10)
	fmt.Println("El area de r1 es :", r1.area())
	fmt.Println("El area de r2 es :", r2.area())

}

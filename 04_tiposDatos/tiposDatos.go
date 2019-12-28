package main

func main() {
	/*basicos*/
	//string o cadena
	var nombre = "luis"
	var apellido = "ana"

	//los string con comilla invertida
	//respetan la indentacion y espacios
	var textoTalCual string = `
	<!DOCTYPE html>
	<html>
		<head>
			<title>Titulo</title>
		</head>
		<body>
			<h1>encabezado</h1>
		</body>
	</html>
	`
	//entero
	var altura uint8   // de 0 a 255
	var altura2 uint16 // de 0 a 65535
	var altura3 uint32 // de 0 a 4294967295
	var altura4 uint64 // de 0 a 18446744073709551615
	//var altura5 uint //dependiendo de arquitectura tomara uint32 o unit64

	var edad int8   //de -128 a 127
	var edad2 int16 //de -32768 a 32767
	var edad3 int32 //de -2147483648 a 2147483647
	var edad4 int64 //de -9223372036854775808 a 9223372036854775807
	//var edad5 int //dependiendo de arquitectura tomara int32 o int64

	//flotante o decimal

	var precio float32 // 6 digitos despues del punto
	precio = 1.232323
	var precio2 float64 //15 digitos despues del punto
	precio2 = 1.121212121212121

	//booleano
	esCasado := true
	tieneHijos := false

	/*otros*/
	var otro1 byte = 12
	//byte        alias for uint8
	var otro2 rune = 232323232
	//rune        alias for int32
	//interfaces

}

/* en golang no tenemos try ni catch como en python
para manejar errores*/

package main

import (
	"errors"
	"fmt"
)

var (
	NovalidUser       = errors.New("el usuario no es valido")
	RegistProcessUser = errors.New("usuario en proceso de registro")
	DefaultError      = errors.New("algo paso")
)

func baneado(user string) (err error) {
	ban := false
	switch user {
	case "miguel":
		ban = true
	case "carlos":
		ban = false
	case "juan":
		return NovalidUser
	case "pedro":
		return RegistProcessUser
	default:
		return DefaultError
	}

	if !ban {
		fmt.Printf("El usuario %s no esta baneado \n", user)
	} else {
		fmt.Printf("El usuario %s esta baneado \n", user)
	}
	return
}

func CheckError(e error) {
	if e != nil {
		fmt.Println("Error : ", e)
	}
}

func main() {
	err := baneado("pedro")
	CheckError(err)

	err = baneado("carlos")
	CheckError(err)

	err = baneado("juan")
	CheckError(err)

	err = baneado("miguel")
	CheckError(err)

}

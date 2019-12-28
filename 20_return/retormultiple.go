package main

import (
	"fmt"
)

func multi(numero int) (n1, n2, n3 int) {
	n1 = numero * 2
	n2 = numero * 4
	n3 = numero * 6
	return
}
func main() {
	r1, r2, r3 := multi(2)
	fmt.Println(r1, r2, r3)

}

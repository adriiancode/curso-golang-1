package main

import (
	"fmt"
)

func main() {
	//primera forma
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	//segunda forma (while)
	l := 0
	for l < 10 {
		fmt.Println(l)
		l++
	}

}

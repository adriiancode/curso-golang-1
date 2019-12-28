package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	err := ioutil.WriteFile("archi.txt", []byte("prueba de escritura"), 0644)
	if err != nil {
		log.Fatal(err)
	}

	content, err := ioutil.ReadFile("archi.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%q", content)
}

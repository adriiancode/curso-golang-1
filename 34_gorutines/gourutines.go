package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err) //para el programa si ocurre error
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) //imprime el error pero no para el programa
			continue
		}
		// manejarConn(conn) //unica concexion
		go manejarConn(conn) //conexion multiple
	}
}

func manejarConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("12:12:12\n\r"))
		if err != nil {
			return //ejemplo si el cliente se desconecta
		}
		time.Sleep(1 * time.Second)
	}
}

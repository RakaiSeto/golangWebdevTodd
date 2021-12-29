package main

import (
	"fmt"
	"net"
	"log"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Fprintln(conn, "Hello from TCP")
		fmt.Fprintln(conn, "Hey, how are you")
		fmt.Fprintln(conn, "I'm good")
		conn.Close()
	}
}
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Exit with status 1 if name is not provided
	if len(os.Args) < 2 {
		fmt.Print("You must enter your name as argument\n")
		os.Exit(1)
	}
	// server address and port
	host := "127.0.0.1:8000"
	// dial the server
	conn, err := net.Dial("tcp", host)
	// if connection is not successful, exit.
	if err != nil {
		fmt.Println("Cannot connect to server")
		os.Exit(1)
	}
	// send the name of client, as provided in os args, to the server
	name := fmt.Sprintf("%s\n", os.Args[1])
	conn.Write([]byte(name))

	// manage client requests

	// go routine to read client data from its stdin and sending it to the server.
	go func(conn net.Conn) {
		r := bufio.NewReader(os.Stdin)
		for {
			data, _ := r.ReadBytes('\n')
			(conn).Write(data)
		}
	}(conn)

	// go routine for receing data from server and displaying it to the client's terminal
	go func(conn net.Conn) {
		s := bufio.NewReader(conn)
		for {
			content, _ := s.ReadBytes('\n')
			fmt.Printf("%s", string(content))
			//fmt.Printf("Client Message: %s", string(content))
		}
	}(conn)

	// run for infinite, until client closes the connection.
	for {

	}
}

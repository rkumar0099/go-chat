package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	//if len(os.Args) < 2 {
	//	fmt.Print("You must enter your name as argument\n")
	//	os.Exit(1)
	//}
	host := "127.0.0.1:8000"
	conn, err := net.Dial("tcp", host)
	if err != nil {
		fmt.Println("Cannot connect to server")
		os.Exit(1)
	}
	//name := fmt.Sprintf("%s\n", os.Args[1])
	//conn.Write([]byte(name))
	go func(conn net.Conn) {
		r := bufio.NewReader(os.Stdin)
		for {
			data, _ := r.ReadBytes('\n')
			(conn).Write(data)
		}
	}(conn)

	go func(conn net.Conn) {
		s := bufio.NewReader(conn)
		for {
			content, _ := s.ReadBytes('\n')
			fmt.Printf("Client Message: %s", string(content))
		}
	}(conn)
	time.Sleep(1 * time.Minute)
}

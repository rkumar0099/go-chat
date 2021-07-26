package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

type Client struct {
	name string
	conn net.Conn
	id   int
}

func main() {
	// array to store the address of client record
	clients := make([]*Client, 2)
	// channel to know whether any client has been disconnected
	connections := make(chan bool, 2)
	// for the id of client
	count := 0
	// host address is "127.0.0.1 and port is "8000"
	host := "127.0.0.1:8000"

	// run the server
	lis, err := net.Listen("tcp", host)
	if err != nil {
		fmt.Println("[server]: error connecting to tcp socket")
		os.Exit(1)
	}

	fmt.Println("Server listening at port 8000")
	// wait for the connection by a client
	for {
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println("Cannot connect to the client")
			continue
		}
		// new client record
		c := &Client{
			name: "",
			conn: conn,
			id:   count,
		}
		fmt.Printf("Client %d is connected\n", count)
		clients[count] = c
		count++

		if count == 2 {
			for i := range clients {
				// handle each client in seperate go routine
				go handleClient(clients[i], clients, connections)
			}
			break
		}
	}
	// wait for the connections channel to reach a size of 2 before exiting the main function
	for {
		if len(connections) == 2 {
			break
		}
	}
}

func handleClient(client *Client, clients []*Client, connections chan bool) {
	// bufio reader to read the content send by a client to server
	reader := bufio.NewReader(client.conn)
	// read the name of the client provided by the client itself after connection
	name, _ := reader.ReadBytes('\n')
	client.name = strings.TrimSuffix(string(name), "\n")

	// serve data requests of client
	for {
		data, err := reader.ReadBytes('\n')
		// client conn is closed. Exit the function
		if err != nil {
			msg := client.name + " is disconnected\n"
			fmt.Println(msg)
			connections <- false
			broadcast(client, clients, msg)
			break
		}
		// broadcast data to all clients
		if len(data) >= 1 {
			fmt.Printf("Client message: %s", string(data))
			msg := "Message from " + client.name + ": " + string(data)
			broadcast(client, clients, msg)
		}
	}
}

func broadcast(client *Client, clients []*Client, data string) {
	for _, j := range clients {
		if client.id != j.id {
			(j.conn).Write([]byte(data))
		}
	}
}

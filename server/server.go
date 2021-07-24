package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

type Client struct {
	name string
	conn net.Conn
	id   int
}

func main() {
	clients := make([]*Client, 2)
	//conns := make([]net.Conn, 2)
	count := 0
	host := "127.0.0.1:8000"
	lis, err := net.Listen("tcp", host)
	if err != nil {
		fmt.Println("[server]: error connecting to tcp socket")
		os.Exit(1)
	}
	fmt.Println("Server listening at port 8000")

	for {

		conn, err := lis.Accept()
		if err != nil {
			fmt.Println("Cannot connect to the client")
			continue
		}
		//time.Sleep(100 * time.Millisecond)
		count++
		c := &Client{
			name: "",
			conn: conn,
			id:   count,
		}
		fmt.Printf("Client %d is connected\n", count)
		clients[count-1] = c
		if count == 2 {
			for i := range clients {
				go handleClient(clients[i], clients)
			}
			break
		}
	}
	time.Sleep(1 * time.Minute)
}

func handleClient(client *Client, clients []*Client) {
	reader := bufio.NewReader(client.conn)
	name, _ := reader.ReadBytes('\n')
	client.name = string(name)
	for {
		data, _ := reader.ReadBytes('\n')
		fmt.Printf("Client message: %s", string(data))
		//fromS, _ := r.ReadBytes('\n')
		broadcast(client, clients, data)
	}
}

func broadcast(client *Client, clients []*Client, data []byte) {
	message := "Message from " + client.name + ": " + string(data)
	newMsg := []byte(message)
	for _, j := range clients {
		if client.conn != j.conn {
			(j.conn).Write(newMsg)
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"net"
)

var clients = make(map[net.Conn]string)

func handleClient(conn net.Conn) {
	defer conn.Close()
	fmt.Println("New client connected", conn.RemoteAddr().String())

	clients[conn] = conn.RemoteAddr().String()

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Client disconnected", conn.RemoteAddr().String())
			delete(clients, conn)
			return
		}

		fmt.Print("Received: " + message)
		for c := range clients {
			if c != conn {
				c.Write([]byte(message))
			}
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Print("Error starting server: ", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server started on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Print("Error accepting connection: ", err)
			continue
		}
		go handleClient(conn)
	}
}

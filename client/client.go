package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error conencting to serer: ", err)
		return
	}
	defer conn.Close()

	go func() {
		for {
			message, _ := bufio.NewReader(conn).ReadString('\n')
			fmt.Print("Incoming message: ", message)
		}
	}()

	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		conn.Write([]byte(text + "\n"))
	}
}

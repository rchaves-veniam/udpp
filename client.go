package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")

	defer conn.Close()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Input: ")
		text, _ := reader.ReadString('\n')

		// Send to socket
		buf := []byte(text + "\n")
		_, err := conn.Write(buf)
		if err != nil {
			fmt.Println(text, err)
		}

		// Listen for reply
		msg, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Got: " + msg)
	}
}

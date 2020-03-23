package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	fmt.Println("Starting server...")

	// Listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081")

	// Accept connection
	conn, _ := ln.Accept()

	for {
		msg, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Got:", string(msg))
		newMsg := strings.ToUpper(msg)
		conn.Write([]byte(newMsg + "\n"))
	}
}

package udpp

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func Example_serverClientConnection() {
	// Server starts. Client starts.
	// Client sends 'test', server replies 'TEST'
	var serverPort string = "8081"
	var serverIP string = "127.0.0.1"

	// Launch server in background
	go server(serverIP, serverPort)
	// Start client. On exit, client will terminate server.
	client(serverIP, serverPort)

	// Output: Connecting... Success!
	// Server:test
	// Client:TEST
}

func server(host, port string) {
	// Listen on all interfaces
	addr := net.JoinHostPort(host, port)
	ln, _ := net.Listen("tcp", addr)

	// Accept connection
	conn, _ := ln.Accept()
	defer conn.Close()

	msg, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Server:", string(msg))
	if len(msg) < 1 {
		conn.Close()
		return
	}
	newMsg := strings.ToUpper(msg)
	conn.Write([]byte(newMsg + "\n"))
}

func client(host, port string) {
	addr := net.JoinHostPort(host, port)
	fmt.Print("Connecting... ")
	conn, err := net.Dial("tcp", addr)
	for err != nil {
		conn, err = net.Dial("tcp", addr)
	}
	fmt.Println("Success!")

	defer conn.Close()

	text := "test"
	buf := []byte(text + "\n")
	_, err = conn.Write(buf)
	msg, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println("Client:" + msg)
}

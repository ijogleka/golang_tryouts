package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

// Simple TCP Server. It
// Reads for any bytes coming in, on a specific port
// On receiving bytes, it prints them out delimited by a newline character by default
// It sends some data back out to the client.
//
// How to test?
// 1. Start the server by running `go run main.go`
// 2. On a different terminal, run `telnet localhost 8080`
//    and play around for 10 seconds. type some text in, and so on
func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		// Process the data received on the connection in a separate thread
		// a.k.a. goroutine
		// So that the server can continue listening for newer connections
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	// Process the connection for 10 seconds and then just close it pronto
	// Noone should be able to access my precious resources for more than 10
	// seconds...in one go
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		fmt.Println("CONNECTION TIMED OUT")
	}

	// Get a scanner to scan the bytes received on the input connection
	// as Text
	scanner := bufio.NewScanner(conn) // Since `conn` supports the `Reader` interface
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say %s\n", ln)
	}
	defer conn.Close()

	// Check if we get to this line
	fmt.Println("Did we reach this line?")
}

package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

// Simple TCP Server. It
// 1. Listens for any connection requests coming in on the port 8080
// 2. On receiving a connection request, it
// 3. Accepts it and writes strings to the connection (i.e. sending it back to the client)
// 4. Closes the connection once the writes have been done
// 5. Keeps listening to more requests coming in.
//
// NOTES:
// 1. It's all about the interfaces here
//    The `conn` supports the Writer Interface and hence can be passed into all the
//    WriteString, ...print... functions
// 2. Cool ai'nt it?
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

		io.WriteString(conn, "\nHello from tcp server\n")
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintf(conn, "%v", "Well, I hope!")

		conn.Close()
	}
}

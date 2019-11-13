package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println("Could not dial in")
	}
	defer conn.Close()

	// fmt.Fprintln(conn, "Trying to dial in..")
	fmt.Fprintln(conn, "Trying to dial in")
}

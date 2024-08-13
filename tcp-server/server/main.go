package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
        fmt.Println("Error:", err)
        return
    }

	// stop listening, any blocked Accept Accept operations will be unblocked and return errors.
	defer listener.Close()

	for {
		// This function blocks until a client connects.
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		// Conn is a generic stream-oriented network connection. Multiple goroutines may invoke methods on a Conn simultaneously.
		// we can read from and write data to a connection

		// After accepting a client connection, you’ll typically spawn a new goroutine to handle that connection.
		// this allows handling multiple clients concurrently
		go handleClient(conn)
	}

}

func handleClient(conn net.Conn) {
	// When the communication is done, you should close the connection using conn.Close().
	defer conn.Close()

	// Read and process data from the client
    // ...
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
        fmt.Println("Error:", err)
        return
    }

	fmt.Printf("Received: %s\n", buffer[:n])

    // Write data back to the client
    // ...

}
package main

import (
	"fmt"
	"net"
	"sync"
)

var errorCount int

func performConnection() {
	// connect to server by specifying network and IP + port
	conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
		errorCount++
        fmt.Println("Error:", err)
        return
    }
    defer conn.Close()

    // Send data to the server
    // ...
	data := []byte("Hello from client!")
	_, err = conn.Write(data)
	if err != nil {
		errorCount++
        fmt.Println("Error:", err)
        return
    }

    // Read and process data from the server
    // ...

}

func main() {
	var wg sync.WaitGroup

	for i:=0; i<1000; i++ {
		wg.Add(1)
		go func() {
            defer wg.Done()
            performConnection()
        }()
	}

	wg.Wait()

	fmt.Println("Error count:", errorCount)
}
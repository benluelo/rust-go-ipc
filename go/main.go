package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Create a Unix domain socket and listen for incoming connections.
	socket, err := net.Listen("unix", "/tmp/ipc-test.sock")
	if err != nil {
		log.Fatal(err)
	}

	// Cleanup the sockfile.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Remove("/tmp/ipc-test.sock")
		os.Exit(1)
	}()

	for {
		// Accept an incoming connection.
		conn, err := socket.Accept()
		if err != nil {
			log.Fatal(err)
		}

		log.Println("recieved connection at ", conn)

		// Handle the connection in a separate goroutine.
		go handle_connection(conn)
	}
}

func handle_connection(conn net.Conn) {
	defer conn.Close()
	// Create a buffer for incoming data.
	buf := make([]byte, 4096)

	// Read data from the connection.
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	// Echo the data back to the connection.
	_, err = conn.Write(buf[:n])
	if err != nil {
		log.Fatal(err)
	}
}
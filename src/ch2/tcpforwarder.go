package main
// TCP forwarder

import (
	"log"
	"net"
	"io"
	// "fmt"
)

func handleconn(conn net.Conn) {
	// Make a buffer
	dataBuffer := make([]byte, 512)
	fwdAddress := "127.0.0.1:9002"
	tcpAddr, _ := net.ResolveTCPAddr("tcp", fwdAddress)

	// Forward data to a different port
	log.Println("Forwarding data to 127.0.0.1:9002")
	fwdConn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Println("Error forwarding, closing connection. ", err)
		conn.Close()
		return
	}

	for {
		recvLen, err := conn.Read(dataBuffer[0:])
		if err == io.EOF {
			log.Println("Client disconnected")
			break
		}

		if err != nil {
			log.Println("Unexpected Error reading data from connection")
			break
		}

		log.Printf("Recieved %d bytes: %s\n", recvLen, string(dataBuffer))

		fwdConn.Write(dataBuffer)
		if err != nil {
			log.Println("Error forwarding data")
		}
		log.Println("Data Forwarded")
	}
	fwdConn.Close()
}

func main() { 
// bind to TCP port 9001
listener, err := net.Listen("tcp", ":9001")

if err !=nil {
	log.Fatalln("Unable to bind to port.")
}

log.Println("Listening on 0.0.0.0:9001")
for {
	// Wait for connection.
	conn, err := listener.Accept()
	if err != nil {
		log.Fatalln("Unable to accept connection.")
	}
	log.Println("Recieved Connection from: ", conn.RemoteAddr().String())
	// Handle the connection 
	go handleconn(conn)
}
}
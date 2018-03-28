package mychat

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

// Peer defines common methods for Client and Server
type Peer interface {
	Start()
}

// ChatServer defines a server for handling chat messages
type ChatServer struct {
	IPAddress string
	Port      int
}

// ChatClient defines a client for sending/recieving messages
type ChatClient struct {
	IPAddress string
	Port      int
}

// Start method starts the server
func (p *ChatServer) Start() {
	fmt.Println("Starting Chat Server on address", p.IPAddress, "and port", p.Port)
	address := p.IPAddress + ":" + strconv.Itoa(p.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("Error :", err)
	}
	fmt.Println("Listening for connections")
	conn, connErr := listener.Accept()
	if connErr != nil {
		log.Fatal("Error :", connErr)
	}
	go ReadFromReaderToStdOut(conn)
	go ReadFromStdOutToWriter(conn)
}

// Start method starts the client
func (p *ChatClient) Start() {
	fmt.Println("Starting Chat Client on address", p.IPAddress, "and port", p.Port)
	address := p.IPAddress + ":" + strconv.Itoa(p.Port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal("Error :", err)
	}
	go ReadFromReaderToStdOut(conn)
	go ReadFromStdOutToWriter(conn)
}

// ReadFromReaderToStdOut reads from given reader and prints to output
func ReadFromReaderToStdOut(reader io.Reader) {
	for {
		fmt.Print(GetMessageFromReader(reader))
	}
}

// ReadFromStdOutToWriter reads from stdout and prints to given writer
func ReadFromStdOutToWriter(writer io.Writer) {
	for {
		fmt.Fprint(writer, GetMessageFromReader(os.Stdin))
	}
}

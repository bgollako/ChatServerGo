package main

import (
	"flag"
	"learninggo/mychat"
	"log"
	"sync"
)

func main() {
	var port int
	var isListening bool
	var ipAddress string
	flag.BoolVar(&isListening, "listen", false, "Listens on the given ip address")
	flag.StringVar(&ipAddress, "ipaddress", "", "IP Address for the server to bind to")
	flag.IntVar(&port, "port", 8080, "Port to listen on")
	flag.Parse()
	var wg sync.WaitGroup
	if isListening {
		if ipAddress == "" {
			log.Fatalf("Did not enter IP Address")
		}
		wg.Add(2)
		server := &mychat.ChatServer{IPAddress: ipAddress, Port: port}
		server.Start()
	} else {
		wg.Add(2)
		client := &mychat.ChatClient{IPAddress: ipAddress, Port: port}
		client.Start()
	}
	wg.Wait()
}

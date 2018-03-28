package mychat

import (
	"bufio"
	"io"
	"log"
)

// GetMessageFromReader reads data from the reader
func GetMessageFromReader(reader io.Reader) string {
	buffReader := bufio.NewReader(reader)
	message, err := buffReader.ReadString('\n')
	if err != nil {
		log.Fatal("Error :", err)
	}
	return message
}

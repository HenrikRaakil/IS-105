package main

import (
	"crypto/rand"
	"fmt"
	"net"
)

func main() {
	// Genererer mock data
	payload := make([]byte, 1<<24) // Allokerer 2^24 = 16777216 bytes
	_, err := rand.Read(payload)
	if err != nil {
		fmt.Println(err)
	}

	listener, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println(err)
	}

	//go func() {
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn.Close()
		fmt.Println("The server is waiting")
		_, err = conn.Write(payload) // den første delen brukes til å sende x antall bytes. "_" betyr at man ignorerer det som vanligvis står der
		if err != nil {
			fmt.Println(err)
		}
	}
}
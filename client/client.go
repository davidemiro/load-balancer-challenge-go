package client

import (
	"log"
	"net"
)

func StartClient(name string, addr string) {

	conn, err := net.Dial("tcp", addr)

	log.SetPrefix(name + " ")

	log.Println("Start Client")
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	conn.Write([]byte("I am a client"))

	buffer := make([]byte, 1024)
	conn.Read(buffer)
	message := string(buffer)

	log.Println(message)

}

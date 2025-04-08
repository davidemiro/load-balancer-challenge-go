package client

import (
	"fmt"
	"net"
)

func StartClient() {
	conn, err := net.Dial("tcp", "127.0.0.1:8082")

	fmt.Println("Start Client")
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	buffer := make([]byte, 1024)

	conn.Read(buffer)
	message := string(buffer)

	fmt.Println(message)

}

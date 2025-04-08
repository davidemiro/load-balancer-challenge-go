package server

import (
	"fmt"
	"net"
)

type Server struct {
	handler *HttpHandler
	name    string
	ip      string
	port    string
}

func (server *Server) NewServer(name string, ip string, port string) {
	server.name = name
	server.ip = ip
	server.port = port

}

func (server *Server) Start() {

	listener, err := net.Listen("tcp", server.ip+":"+server.port)
	if err != nil {
		fmt.Println("[ERROR] starting server: " + err.Error())
	}

	defer listener.Close()

	fmt.Println("[LISTENING] on port " + server.port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("[ERROR] Accepting connection:", err)
			continue
		}

		conn.Write([]byte("Hello world!\n I am " + server.name))

	}

}

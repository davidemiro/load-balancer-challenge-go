package server

import (
	"log"
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
	log.SetPrefix(server.name + " ")

}

func (server *Server) Start() {

	listener, err := net.Listen("tcp", server.ip+":"+server.port)
	if err != nil {
		log.Println("[ERROR] starting server: " + err.Error())
	}

	defer listener.Close()

	log.Println("[LISTENING] on port " + server.port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("[ERROR] Accepting connection:", err)
			continue
		}

		buffer := make([]byte, 1024)
		conn.Read(buffer)

		log.Println("[MESSAGE]: " + string(buffer))

		conn.Write([]byte("Hello world! I am " + server.name))

	}

}

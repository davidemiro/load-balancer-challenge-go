package server

import (
	"fmt"
	"log"
	"net/http"
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
	server.handler = new(HttpHandler)
	server.handler.NewHttpHandler(server.name)
	log.Printf("Starting server %s at port %s and address %s\n", server.name, server.port, server.ip)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", server.ip, server.port), server.handler); err != nil {
		log.Fatalln(err)
	}

}

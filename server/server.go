package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	handler *HttpHandler
	name    string
	ip      string
	port    string
}

func (server *Server) NewServer(name string, ip string, port string) *Server {
	return &Server{nil, name, ip, port}
}

func (server *Server) Start() {
	server.handler = &HttpHandler{server.name}
	http.ListenAndServe(fmt.Sprintf("%s:%s", server.ip, server.port), server.handler)

}

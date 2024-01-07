package server

import (
	"fmt"
	"net/http"
)

type HttpHandler struct {
	name string
}

func (httpHandler *HttpHandler) NewHttpHandler(name string) {
	httpHandler.name = name
}

func (httpHandler *HttpHandler) GetName() string {
	return httpHandler.name
}

func (httpHandler *HttpHandler) SetName(newName string) {
	httpHandler.name = newName
}

// implement `ServeHTTP` method on `HttpHandler` struct
func (httpHandler *HttpHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// create response binary data
	fmt.Println(httpHandler.name)
	fmt.Fprintf(w, "Hello World!\nThe server %s answer\n", httpHandler.name)
}

package server

import (
	"fmt"
	"net/http"
)

type HttpHandler struct {
	name string
}

func (httpHandler *HttpHandler) GetName() string {
	return httpHandler.name
}

func (HttpHandler *HttpHandler) SetName(newName string) {
	HttpHandler.name = newName
}

// implement `ServeHTTP` method on `HttpHandler` struct
func (h HttpHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// create response binary data
	fmt.Fprintf(w, "Hello World!\nThe server %s answer\n", h.name)
}

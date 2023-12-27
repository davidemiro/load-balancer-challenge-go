package LoadBalancer

import (
	"LoadBalancer/server"
	"net/http"
)

func main() {

	//create a new handler
	handler := server.HttpHandler{}

	//listen and serve
	http.ListenAndServe(":9000", handler)
}

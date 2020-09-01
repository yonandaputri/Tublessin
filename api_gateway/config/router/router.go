package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var serverHost, serverPort string

// CreateRouter for creating new Route
func CreateRouter() *mux.Router {
	return mux.NewRouter()
}

// StartServer routing
func StartServer(r *mux.Router) {
	configServer()
	log.Println("Server Start at http://" + serverHost + ":" + serverPort)
	http.ListenAndServe(serverHost+":"+serverPort, r)
}

// Config server Host and Port
func configServer() {
	serverHost = "localhost"
	serverPort = "8080"
}

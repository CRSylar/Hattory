package main

import (
	"github.com/CRSylar/hattory/cmd"
	"github.com/CRSylar/hattory/handlers"
)

func main() {

	// Create a new ServeMux - Request Handler
	h := handlers.CreateRequestsHandler()
	handlers.RegisterRoutes()
	handlers.RegisterApiRoutes()
	// Create a new Server
	s := cmd.NewServer(h)
	// Start the server
	cmd.StartServer(s)
}

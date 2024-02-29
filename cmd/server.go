package cmd

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
)

func NewServer(handler http.Handler) *http.Server {
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	s := http.Server{
		Addr:    ":"+port,
		Handler: handler,
	}
	return &s
}

func StartServer(s *http.Server) {
	fmt.Printf("Server starting on port %s...\n", s.Addr)
	fmt.Println("Press Ctrl+C to stop the server")
	for err := s.ListenAndServe(); err != nil; {
		newPort, _ := strconv.Atoi(s.Addr[1:])
		s.Addr = ":" + fmt.Sprint(newPort+1)
		slog.Warn(fmt.Sprintf("Port %d is already in use, trying %s...\n", newPort, s.Addr))
	}
	
	s.Close()
}
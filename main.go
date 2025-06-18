package main

import (
	"fmt"
	"log"
	"net"
)

type Config struct {
	ListenAddr string
}

type Server struct {
	Config
	ln net.Listener
}

func NewServer(cfg Config) *Server {

	if cfg.ListenAddr == "" {
		cfg.ListenAddr = ":8080"
	}

	return &Server{Config: cfg}
}

func (s *Server) StartServer() error {

	var err error
	s.ln, err = net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return fmt.Errorf("failed to listen on port %s: %w", s.ListenAddr, err)
	}
	defer s.Close() // ensure listener is closed when start exits.
	// Defer keyword schedules call to be executed just before the function it is in returns

	log.Printf("TCP Server Listening on %s\n", s.ListenAddr)

	//infinite loop to continuously accept new connections
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				log.Println("Listener timed out, shutting down.")
				break
			}
			log.Printf("Error accepting connection: %v\n", err)
			continue
		}
		go s.handleConnection(conn) // go keyword initiates a new goroutine, which runs concurrently
	}
	return nil
}

func (s *Server) Close() error {
	return nil
}

func (s *Server) handleConnection(conn net.Conn) error {
	return nil
}

package server

import (
	"fmt"
	"net"
	"ooxymoron/pkg/utils"
)

type Options struct {
	ListenAddr   string
	UpstreamAddr string
}

type Server struct {
	opts Options
}

func NewServer(opts Options) *Server {
	return &Server{opts: opts}
}

func (s *Server) Run() error {
	ln, err := net.Listen("tcp", s.opts.ListenAddr)

	if err != nil {
		return fmt.Errorf("Failed to listen at address: %s due to error: %w", s.opts.ListenAddr, err)
	}

	defer ln.Close()
	
	fmt.Printf("Ooxymoron listening on %s -> %s\n", s.opts.ListenAddr, s.opts.UpstreamAddr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			return fmt.Errorf("Failed to accept connection: %w", err)
		}

		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 65535)

	for {
		readBytes, err := conn.Read(buffer)

		if err != nil {
			fmt.Printf("Error reading data: %s\n", err)
			return
		}

		utils.CheckProtocol(buffer[:readBytes])

		fmt.Printf("Received data: %s\n", buffer[:readBytes])
	}
}

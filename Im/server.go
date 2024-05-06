package main

import (
	"fmt"
	"net"
)

type Server struct {
	ip   string
	port int
}

func NewServer(ip string, port int) *Server {
	return &Server{
		ip:   ip,
		port: port,
	}
}

func (s *Server) Handler(conn net.Conn) {
	fmt.Println("Connection successful!")
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.ip, s.port))
	fmt.Println("Server listening...:")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("conn err:", err)
			return
		}

		go s.Handler(conn)
	}
}

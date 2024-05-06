package main

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	ip   string
	port int

	// 在线用户的列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	// 消息广播的message
	Message chan string
}

func NewServer(ip string, port int) *Server {
	return &Server{
		ip:        ip,
		port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
}

// 监听Message广播消息Channel，一旦有消息发送给全部的在线用户
func (s *Server) ListenMessage() {
	for {
		msg := <-s.Message

		s.mapLock.Lock()
		for _, v := range s.OnlineMap {
			v.C <- msg
		}
		s.mapLock.Unlock()
	}
}

// 广播消息
func (s *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + ":" + msg
	s.Message <- sendMsg
}

func (s *Server) Handler(conn net.Conn) {
	fmt.Println("Connection successful!")

	user := NewUser(conn)

	// 用户上线，将用户加到OnlineMap中
	s.mapLock.Lock()
	s.OnlineMap[user.Name] = user
	s.mapLock.Unlock()

	// 广播当前用户上线消息
	s.BroadCast(user, "已上线")

	select {}
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.ip, s.port))
	fmt.Println("Server listening...")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listener.Close()

	// 启动监听Message的goroutine
	go s.ListenMessage()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("conn err:", err)
			return
		}

		go s.Handler(conn)
	}
}

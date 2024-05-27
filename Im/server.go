package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
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
	fmt.Printf("%s Connection successful!\n", conn.RemoteAddr().String())

	user := NewUser(conn, s)

	user.Online()

	isAlive := make(chan struct{})

	// 接收客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}

			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err:", err)
				return
			}

			// 提取用户的消息（去除"\n")
			msg := string(buf[:n-1])

			// 进行消息处理
			user.DoMessage(msg)

			isAlive <- struct{}{}
		}
	}()

	for {
		select {
		case <-isAlive:
			// 更新下方的定时器

		case <-time.After(60 * time.Second):
			// 超时后将用户剔除下线

			user.SendMsg("你被踢了")

			close(user.C)

			// 关闭连接
			conn.Close()

			return
		}
	}
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

package main

import "net"

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn
}

// 创建一个用户
func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,
	}

	// 启动监听当前 User Channel消息的goroutine
	go user.ListenMsg()

	return user
}

func (u *User) ListenMsg() {
	for {
		msg := <-u.C

		u.conn.Write([]byte(msg + "\n"))
	}
}

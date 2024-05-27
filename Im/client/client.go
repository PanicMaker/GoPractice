package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn

	flag int
}

func NewClient(ip string, port int) (*Client, error) {
	client := &Client{
		ServerIp:   ip,
		ServerPort: port,
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return nil, err
	}

	client.conn = conn

	return client, nil
}

func (c *Client) DealResponse() {
	io.Copy(os.Stdout, c.conn)
}

func (c *Client) menu() bool {
	var flag int

	fmt.Println("1.Public Chat")
	fmt.Println("2.Private Chat")
	fmt.Println("3.Rename")
	fmt.Println("0.Quit")

	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		c.flag = flag
		return true
	} else {
		fmt.Println("Please input right number")
		return false
	}
}

func (c *Client) Run() {
	for c.flag != 0 {
		for c.menu() != true {

		}

		switch c.flag {
		case 1:
			break
		case 2:
			break
		case 3:
			c.UpdateName()
			break
		}
	}
}

func (c *Client) UpdateName() bool {
	fmt.Println("Please input username")
	fmt.Scanln(&c.Name)

	sendMsg := "rename|" + c.Name + "\n"
	_, err := c.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn Write err:", err)
		return false
	}
	return true
}

var (
	serverIp   string
	serverPort int
)

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器IP地址")
	flag.IntVar(&serverPort, "port", 123, "设置服务器端口")
}

func main() {
	flag.Parse()

	client, err := NewClient(serverIp, serverPort)
	if err != nil {
		fmt.Println("Connect Server Error...")
		return
	}

	go client.DealResponse()

	fmt.Println("Connect Server Success...")

	client.Run()
}

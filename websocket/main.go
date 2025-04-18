package main

import (
	"context"
	"encoding/json" // 添加 json 包导入
	"fmt"           // 添加 fmt 包导入
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/mailru/easygo/netpoll"
)

// 消息类型
type Message struct {
	Type    string `json:"type"`
	Channel string `json:"channel,omitempty"`
	Content string `json:"content"`
	Sender  string `json:"sender"`
}

// 客户端连接
type Client struct {
	id        string
	conn      net.Conn
	send      chan []byte
	channels  map[string]bool
	createdAt time.Time
}

// 频道管理器
type ChannelManager struct {
	clients    map[string]map[*Client]bool // 按频道归类的客户端
	broadcast  chan Message
	register   chan *Client
	unregister chan *Client
	join       chan struct {
		client  *Client
		channel string
	}
	leave chan struct {
		client  *Client
		channel string
	}
	mu sync.RWMutex
}

func NewChannelManager() *ChannelManager {
	return &ChannelManager{
		clients:    make(map[string]map[*Client]bool),
		broadcast:  make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		join: make(chan struct {
			client  *Client
			channel string
		}),
		leave: make(chan struct {
			client  *Client
			channel string
		}),
	}
}

func (cm *ChannelManager) Run() {
	for {
		select {
		case client := <-cm.register:
			// 注册新客户端
			log.Printf("新客户端连接: %s", client.id)

		case client := <-cm.unregister:
			// 注销客户端，从所有频道移除
			cm.mu.Lock()
			for channel := range client.channels {
				delete(cm.clients[channel], client)
			}
			cm.mu.Unlock()
			close(client.send)
			log.Printf("客户端断开连接: %s", client.id)

		case join := <-cm.join:
			// 客户端加入频道
			cm.mu.Lock()
			if _, ok := cm.clients[join.channel]; !ok {
				cm.clients[join.channel] = make(map[*Client]bool)
			}
			cm.clients[join.channel][join.client] = true
			join.client.channels[join.channel] = true
			cm.mu.Unlock()
			log.Printf("客户端 %s 加入频道: %s", join.client.id, join.channel)

		case leave := <-cm.leave:
			// 客户端离开频道
			cm.mu.Lock()
			if _, ok := cm.clients[leave.channel]; ok {
				delete(cm.clients[leave.channel], leave.client)
				delete(leave.client.channels, leave.channel)
			}
			cm.mu.Unlock()
			log.Printf("客户端 %s 离开频道: %s", leave.client.id, leave.channel)

		case message := <-cm.broadcast:
			// 向特定频道广播消息
			data, err := json.Marshal(message) // 引入 json 包
			if err != nil {
				log.Printf("消息序列化错误: %v", err)
				continue
			}

			cm.mu.RLock()
			clients := cm.clients[message.Channel]
			cm.mu.RUnlock()

			// 向频道内所有客户端发送消息
			for client := range clients {
				select {
				case client.send <- data:
					// 成功将消息放入客户端发送队列
				default:
					// 客户端发送队列已满，移除该客户端
					cm.mu.Lock()
					delete(cm.clients[message.Channel], client)
					delete(client.channels, message.Channel)
					close(client.send)
					cm.mu.Unlock()
				}
			}
		}
	}
}

func main() {
	// 创建epoll实例
	poller, err := netpoll.New(&netpoll.Config{})
	if err != nil {
		log.Fatalf("无法创建netpoll实例: %v", err)
	}

	// 创建频道管理器
	channelManager := NewChannelManager()
	go channelManager.Run()

	// 设置HTTP处理
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		// 升级HTTP连接为WebSocket
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			log.Printf("WebSocket升级错误: %v", err)
			return
		}

		// 创建新客户端
		client := &Client{
			id:        generateID(), // 引入 fmt 包
			conn:      conn,
			send:      make(chan []byte, 256),
			channels:  make(map[string]bool),
			createdAt: time.Now(),
		}

		// 注册客户端
		channelManager.register <- client

		// 设置epoll处理
		desc := netpoll.Must(netpoll.HandleRead(conn))
		poller.Start(desc, func(ev netpoll.Event) {
			if ev&(netpoll.EventReadHup|netpoll.EventHup) != 0 {
				// 连接关闭
				poller.Stop(desc)
				channelManager.unregister <- client
				conn.Close()
				return
			}

			// 处理读取事件
			go func() {
				msg, op, err := wsutil.ReadClientData(conn)
				if err != nil {
					log.Printf("读取错误: %v", err)
					poller.Stop(desc)
					channelManager.unregister <- client
					conn.Close()
					return
				}

				// 解析消息
				var message Message
				if err := json.Unmarshal(msg, &message); err != nil { // 引入 json 包
					log.Printf("消息解析错误: %v", err)
					return
				}

				// 处理不同类型的消息
				switch message.Type {
				case "join":
					channelManager.join <- struct {
						client  *Client
						channel string
					}{client, message.Channel}

				case "leave":
					channelManager.leave <- struct {
						client  *Client
						channel string
					}{client, message.Channel}

				case "message":
					message.Sender = client.id
					channelManager.broadcast <- message
				}

				// 发送确认消息
				ack := Message{
					Type:    "ack",
					Content: "收到消息",
				}
				ackData, _ := json.Marshal(ack) // 引入 json 包
				err = wsutil.WriteServerMessage(conn, op, ackData)
				if err != nil {
					log.Printf("写入错误: %v", err)
				}
			}()
		})

		// 启动发送goroutine
		go func() {
			for message := range client.send {
				err := wsutil.WriteServerMessage(conn, ws.OpText, message)
				if err != nil {
					log.Printf("发送错误: %v", err)
					conn.Close()
					break
				}
			}
		}()
	})

	// 创建HTTP服务器
	server := &http.Server{
		Addr: ":8080",
	}

	// 启动服务器
	go func() {
		log.Println("启动WebSocket服务器在 :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe错误: %v", err)
		}
	}()

	// 等待终止信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("关闭服务器...")

	// 创建超时上下文
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// 优雅关闭
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("服务器强制关闭: %v", err)
	}

	log.Println("服务器优雅退出")
}

func generateID() string {
	// 生成唯一ID的实现
	return fmt.Sprintf("%d", time.Now().UnixNano()) // 引入 fmt 包
}

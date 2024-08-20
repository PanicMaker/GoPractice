package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

// RateLimitMiddleware 创建一个速率限制中间件
func RateLimitMiddleware(limit int, interval time.Duration) gin.HandlerFunc {
	// 创建一个带缓冲的通道，缓冲大小为limit
	ticker := time.NewTicker(interval)
	tokens := make(chan struct{}, limit)

	// 启动一个协程定期往通道中添加令牌
	go func() {
		for {
			select {
			case <-ticker.C:
				// 每个时间间隔向通道中填充令牌，直到达到通道的缓冲大小
				for i := 0; i < limit; i++ {
					select {
					case tokens <- struct{}{}:
					default:
					}
				}
			}
		}
	}()

	// 返回Gin中间件
	return func(c *gin.Context) {
		select {
		case <-tokens: // 从通道中取出一个令牌，表示允许请求
			c.Next() // 处理下一个中间件或请求
		default: // 如果通道为空，表示达到速率限制
			c.AbortWithStatus(429) // 429 Too Many Requests
		}
	}
}

func main() {
	r := gin.Default()

	// 每秒最多10个请求
	r.Use(RateLimitMiddleware(10, time.Second))

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	r.Run(":8080")
}

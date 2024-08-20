package main

import (
	"github.com/gin-gonic/gin"
	"sync"
	"sync/atomic"
	"time"
)

func RateLimitLock(qps int) gin.HandlerFunc {
	var count int
	var last = time.Now()
	var mu sync.Mutex

	TimeWindow := 1000 * time.Millisecond

	return func(c *gin.Context) {
		now := time.Now()

		// 这里只对临界区加锁，临界区仅限于更新共享变量的部分
		mu.Lock()
		// 如果时间窗口已经过去，重置计数器和时间窗口
		if now.Sub(last) > TimeWindow {
			count = 0
			last = now
		}

		count++
		currentCount := count
		mu.Unlock()

		// 在锁外执行逻辑，只有当请求数超过限额时才会返回429
		if currentCount <= qps {
			c.Next()
		} else {
			c.AbortWithStatus(429) // 429 Too Many Requests
		}
	}
}

func RateLimitAtomic(qps int) gin.HandlerFunc {
	var count int32 // 使用int32以便于使用原子操作
	var last int64  // 使用Unix时间戳（纳秒）

	TimeWindow := int64(time.Second)

	// 将 last 初始化为当前时间的 Unix 纳秒时间戳
	atomic.StoreInt64(&last, time.Now().UnixNano())

	return func(c *gin.Context) {
		now := time.Now().UnixNano()

		// 使用原子操作获取和更新 last 时间戳
		lastTime := atomic.LoadInt64(&last)
		if now-lastTime >= TimeWindow {
			// 如果时间窗口已过，重置计数器，并更新 last 为当前时间
			atomic.StoreInt64(&last, now)
			atomic.StoreInt32(&count, 1)
		} else {
			// 增加计数器
			currentCount := atomic.AddInt32(&count, 1)

			// 如果请求数超过限制，则返回429
			if currentCount > int32(qps) {
				c.AbortWithStatus(429) // 429 Too Many Requests
				return
			}
		}

		// 继续处理下一个中间件或请求
		c.Next()
	}
}

func main() {
	r := gin.Default()

	r.Use(RateLimitAtomic(5))

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	r.Run(":8080")
}

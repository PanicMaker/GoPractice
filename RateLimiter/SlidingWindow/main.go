package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"sync"
	"time"
)

// SlidingWindowLimiter 滑动窗口限流器
type SlidingWindowLimiter struct {
	limit        int           // 窗口请求上限
	window       int64         // 窗口时间大小
	smallWindow  int64         // 小窗口时间大小
	smallWindows int64         // 小窗口数量
	counters     map[int64]int // 小窗口计数器
	mu           sync.Mutex    // 避免并发问题
}

// NewSlidingWindowLimiter 创建滑动窗口限流器
func NewSlidingWindowLimiter(limit int, window time.Duration, smallWindow time.Duration) (*SlidingWindowLimiter, error) {
	// 窗口时间必须能够被小窗口时间整除
	if window%smallWindow != 0 {
		return nil, errors.New("window cannot be split by integers")
	}

	return &SlidingWindowLimiter{
		limit:        limit,
		window:       int64(window),
		smallWindow:  int64(smallWindow),
		smallWindows: int64(window / smallWindow),
		counters:     make(map[int64]int),
	}, nil
}

func (l *SlidingWindowLimiter) TryAcquire() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	// 获取当前小窗口值
	currentSmallWindow := time.Now().UnixNano() / l.smallWindow * l.smallWindow
	// 获取起始小窗口值
	startSmallWindow := currentSmallWindow - l.smallWindows*(l.smallWindows-1)

	// 计算当前窗口的请求总数
	var count int
	for smallWindow, counter := range l.counters {
		if smallWindow < startSmallWindow {
			delete(l.counters, smallWindow)
		} else {
			count += counter
		}
	}

	// 若到达窗口请求上限，请求失败
	if count >= l.limit {
		return false
	}

	// 若没到窗口请求上限，当前小窗口计数器+1，请求成功
	l.counters[currentSmallWindow]++
	return true
}

func RateLimit() gin.HandlerFunc {

	limiter, _ := NewSlidingWindowLimiter(100, 1*time.Minute, 10*time.Second)

	return func(c *gin.Context) {

		if limiter.TryAcquire() {
			c.Next()
		} else {
			c.AbortWithStatus(429) // 429 Too Many Requests
		}
	}
}

func main() {
	r := gin.Default()

	r.Use(RateLimit())

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	r.Run(":8080")
}

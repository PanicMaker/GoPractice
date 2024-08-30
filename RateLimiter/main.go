package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

func RateLimit(limiter RateLimiter) gin.HandlerFunc {

	return func(c *gin.Context) {

		if ok, _ := limiter.TryAcquire(); ok {
			c.Next()
		} else {
			c.AbortWithStatus(429) // 429 Too Many Requests
		}
	}
}

func main() {
	r := gin.Default()

	//limiter := NewLeakyBucketLimiter(100, 2)

	second := NewSlidingLogLimiterStrategy(1, 1*time.Second)
	minute := NewSlidingLogLimiterStrategy(10, 1*time.Minute)

	limiter, _ := NewSlidingLogLimiter(time.Second, second, minute)

	r.Use(RateLimit(limiter))

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	r.Run(":8080")
}

package main

import (
	"github.com/gin-gonic/gin"
)

func RateLimit(limiter RateLimiter) gin.HandlerFunc {

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

	limiter := NewLeakyBucketLimiter(100, 2)

	r.Use(RateLimit(limiter))

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	r.Run(":8080")
}

package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
)

var limiter = rate.NewLimiter(10, 1)

func main() {
	r := gin.Default()

	r.Use(RateLimit())

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "Ping!")
	})

	r.GET("/wait", func(c *gin.Context) {

		err := limiter.Wait(c)
		if err != nil {
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": "OK!",
		})
	})

	r.GET("/allow", func(c *gin.Context) {

		if limiter.Allow() {
			c.JSON(http.StatusOK, gin.H{
				"data": "OK!",
			})
		} else {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"data": "TooManyRequests",
			})
		}
	})

	r.Run(":8080")
}

func RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 如果取不到令牌就中断本次请求返回 rate limit...
		if !limiter.Allow() {
			c.String(http.StatusTooManyRequests, "rate limit...")
			c.Abort()
			return
		}
		c.Next()
	}
}

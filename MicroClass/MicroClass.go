package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	g.Use(Cors())

	g.POST("/A", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"Url":  c.Request.URL.String(),
			"Data": "你正在请求接口 A",
		})
	})

	g.POST("/B", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Url":  c.Request.URL.String(),
			"Data": "你正在请求接口 B",
		})
	})

	g.Run(":8080")
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, issn")
			c.Header("Access-Control-Expose-headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

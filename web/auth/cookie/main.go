package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Cookie("abc"); err == nil {
			if cookie == "123" {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusUnauthorized, gin.H{
			"msg":  "401",
			"code": http.StatusUnauthorized,
		})
		c.Abort()
		return
	}
}

func main() {
	r := gin.Default()

	r.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("key_cookie")

		if err != nil {
			cookie = "NotSet"
			// 给客户端设置cookie
			//  maxAge int, 单位为秒
			// path,cookie所在目录
			// domain string,域名
			//   secure 是否智能通过https访问
			// httpOnly bool  是否允许别人通过js获取自己的cookie
			c.SetCookie("key_cookie", "value_cookie", 60, "/",
				"localhost", false, true)
		}
		fmt.Printf("cookie的值是： %s\n", cookie)
	})

	r.GET("/login", func(c *gin.Context) {
		c.SetCookie("abc", "123", 60, "/", "localhost", false, true)
		c.String(http.StatusOK, "Login success!")
	})

	r.GET("/home", AuthMiddleWare(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "home",
			"code": http.StatusOK,
			"msg":  "成功",
		})
	})

	r.Run(":9999")
}

package main

import "github.com/gin-gonic/gin"

func MaxAllowed(n int) gin.HandlerFunc {
	reqChan := make(chan struct{}, n)
	acquire := func() { reqChan <- struct{}{} }
	release := func() { <-reqChan }

	return func(c *gin.Context) {
		acquire()
		defer release()
		c.Next()
	}
}

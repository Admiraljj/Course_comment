package controllers

import "github.com/gin-gonic/gin"

func Helloworld(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})
}

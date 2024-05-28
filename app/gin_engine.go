package app

import "github.com/gin-gonic/gin"

func NewGinEngine() *gin.Engine {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})
	return r
}

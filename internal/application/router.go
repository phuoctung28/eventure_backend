package application

import "github.com/gin-gonic/gin"

type Router struct {
}

func (r *Router) With(engine *gin.Engine) {
}

func NewRouter() *Router {
	router := &Router{}
	return router
}

package configuration

import (
	"eventure_backend/internal/application"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine    *gin.Engine
	apiRouter *application.Router
}

func (s *Server) Start() {
	s.apiRouter.With(s.engine)
	s.engine.Run()
}

func NewServer(engine *gin.Engine, apiRouter *application.Router) *Server {
	return &Server{
		engine:    engine,
		apiRouter: apiRouter,
	}
}

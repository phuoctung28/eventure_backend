package app

import (
	"eventure_backend/internal/api"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine    *gin.Engine
	apiRouter *api.Router
}

func (s *Server) Start() {
	s.apiRouter.With(s.engine)
	s.engine.Run()
}

func NewServer(engine *gin.Engine, apiRouter *api.Router) *Server {
	return &Server{
		engine:    engine,
		apiRouter: apiRouter,
	}
}

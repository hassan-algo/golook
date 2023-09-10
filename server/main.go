package server

import (
	"github.com/gin-gonic/gin"
	"github.com/hassan-algo/golook/page"
)

type Server struct {
	r *gin.Engine
}

func CreateNewServer() *Server {
	return &Server{
		r: gin.Default(),
	}
}

func (s *Server) CreatePage(path string, p *page.Page) {
	s.r.GET(path, p.PageHandler)
}

func (s *Server) Start(port string) {
	s.r.Run(port)
}

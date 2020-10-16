package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Server interface {
	Route(method, path string, handlerFunc ...gin.HandlerFunc)
	Run(port uint16) error
}

type server struct {
	http *gin.Engine
}

func (s *server) Route(method, path string, handlerFunc ...gin.HandlerFunc) {
	s.http.Handle(method, path, handlerFunc...)
}

func (s *server) Run(port uint16) error {
	return s.http.Run(fmt.Sprintf(":%d", port))
}

func NewServer(http *gin.Engine) Server {
	return &server{
		http: http,
	}
}

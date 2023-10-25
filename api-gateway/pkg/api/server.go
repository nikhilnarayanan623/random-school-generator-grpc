package api

import (
	"github.com/gin-gonic/gin"
	"github.com/nikhilnarayanan623/random-school-generator-grpc/api-gateway/pkg/api/handler/interfaces"
	"github.com/nikhilnarayanan623/random-school-generator-grpc/api-gateway/pkg/config"
)

type Server struct {
	port   string
	engine *gin.Engine
}

func NewServerHTTP(cfg config.Config, schoolHandler interfaces.SchoolHandler) *Server {

	engine := gin.New()

	engine.Use(gin.Logger())
	engine.GET("school", schoolHandler.GetOne)

	return &Server{
		engine: engine,
		port:   cfg.ApiPort,
	}
}

func (c *Server) Start() {
	c.engine.Run((":" + c.port))
}

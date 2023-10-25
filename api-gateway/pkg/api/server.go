package api

import (
	"api-gateway/pkg/api/handler/interfaces"
	"api-gateway/pkg/config"

	// "net/http"
	// _ "net/http/pprof"

	"github.com/gin-gonic/gin"
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

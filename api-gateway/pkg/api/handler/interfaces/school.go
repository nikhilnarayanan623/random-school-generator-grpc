package interfaces

import "github.com/gin-gonic/gin"

type SchoolHandler interface {
	GetOne(ctx *gin.Context)
}

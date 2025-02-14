package routes

import (
	"app/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/sort", handler.SortHandler)
	return r
}

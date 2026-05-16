package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	systemHandler "github.com/noctivs/smart-community-server/internal/system/handler"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	api := r.Group("/api/v1")

	{
		api.GET("/health", systemHandler.Health)
	}

	return r
}

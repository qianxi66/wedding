package routes

import (
	"github.com/changwei4869/wedding/middleware"
	"github.com/changwei4869/wedding/modules/health"
	"github.com/changwei4869/wedding/utils"
	_ "github.com/changwei4869/wedding/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api")
	{
		v1.GET("/health", health.HealthCheck)
	}

	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	_ = r.Run(utils.HttpPort)

}

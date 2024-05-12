package modules

import (
	_ "github.com/changwei4869/wedding/docs"
	"github.com/changwei4869/wedding/middleware"
	"github.com/changwei4869/wedding/modules/file"
	"github.com/changwei4869/wedding/modules/health"
	"github.com/changwei4869/wedding/modules/tag"
	"github.com/changwei4869/wedding/utils"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api", middleware.Cors())
	{
		// health
		api.POST("/health", health.HealthCheck)
		// tag
		api.GET("/tag/:id", tag.GetTagById)
		// file
		api.POST("/file/upload", file.UploadFile)
	}

	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	

	_ = r.Run(utils.HttpPort)

}

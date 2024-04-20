package routes

import (
	"github.com/changwei4869/wedding/middleware"
	"github.com/changwei4869/wedding/utils"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func createMyRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	return p
}

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	// 设置信任网络 []string
	// nil 为不计算，避免性能消耗，上线应当设置
	_ = r.SetTrustedProxies(nil)

	r.HTMLRender = createMyRender()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	_ = r.Run(utils.HttpPort)

}

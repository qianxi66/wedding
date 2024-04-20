package main

import (
	"github.com/changwei4869/wedding/model"
	"github.com/changwei4869/wedding/routes"
)

func main() {
	// 引用数据库
	model.InitDb()
	// 引入路由组件
	routes.InitRouter()

}

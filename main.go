package main

import (
	"github.com/gin-gonic/gin"
	"gin-study/tool"
	"gin-study/controller"
)

func main() {
	cfg, err := tool.ParseConfig("./config/app.json")
	if err != nil{
		panic(err.Error())
	}

	app := gin.Default()
	registerRouter(app)
	app.Run(cfg.AppHost)


}


// 路由设置
func registerRouter(router *gin.Engine)  {
	new(controller.HelloController).Router(router)
}

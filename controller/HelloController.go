package controller

import (
	"github.com/gin-gonic/gin"
	"gin-study/model"
	"time"
	"gin-study/dao"
	"gin-study/tool"
)

type HelloController struct {

}

func (hello *HelloController) Router(engine *gin.Engine) {
	engine.GET("/hello", hello.Hello)
}

func (hello *HelloController) Hello(context *gin.Context)  {

	userInfo := model.UserInfo{
		Name:"tom",
		Password: "123456",
		CreateTime: time.Now().Unix(),
	}
	user := dao.MemberDao{tool.DbEngine}
	result := user.InsertUserInfo(userInfo)
	if result > 0 {
		context.JSON(200,map[string]interface{}{
			"message": "插入成功",
		})
		return
	}
	context.JSON(200,map[string]interface{}{
		"message": "插入失败",
	})

}
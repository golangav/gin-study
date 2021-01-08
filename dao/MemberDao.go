package dao

import (
	"gin-study/tool"
	"gin-study/model"
	"fmt"
)

type MemberDao struct {
	*tool.Orm
}

func (md *MemberDao) InsertUserInfo(user model.UserInfo) int64{
	result, err := md.InsertOne(&user)
	if err != nil{
		fmt.Println(err.Error())
	}
	return result
}

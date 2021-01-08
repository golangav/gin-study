package model

type UserInfo struct {
	Id int64 `xorm:"pk autoincr" json:"id"`
	Name string `xorm:"varchat(11)" json:"name"`
	Password string `xorm:"varchat(20)" json:"password"`
	CreateTime int64 `xorm:"bigint" json:"create_time"`
} 
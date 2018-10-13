package entity

//用户性别
const (
	WOMEN="W"
	MAN="N"
	Unknow=""
)
//群聊用户
type Userinfo struct {
	Id         int64     `xorm:"pk autoincr bigint(64)" form:"id" json:"id"`
	Avatar	   string 		`xorm:"varchar(150)" form:"avatar" json:"avatar"`
	Nickname    string	`xorm:"varchar(20)" form:"nickname" json:"nickname"`   // 什么角色
	Sex      string	`xorm:"varchar(2)" form:"sex" json:"sex"`   // 什么角色
	Salt      string	`xorm:"varchar(10)" form:"salt" json:"-"`   // 什么角色
	Token      string	`xorm:"varchar(40)" form:"token" json:"token"`   // 什么角色
}

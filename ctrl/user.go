package ctrl

import (
	"../entity"
	"../service"
	"../util"
	"net/http"
	"strconv"
)
type UserCtrl struct {

}

var userService service.UserService

//简单的登录逻辑
func( ctrl * UserCtrl)Authorization(w http.ResponseWriter, req *http.Request) {
	var user entity.Userinfo
	req.ParseForm()
	user.Id,_=strconv.ParseInt(req.Form.Get("id"),10,64)
	user.Token=req.Form.Get("token")
	//如果这个用的上,那么可以直接

	if user.Id==0 ||  len(user.Token)==0 {
		util.RespFail(w,"缺少参数")
		return
	}

	token,err := userService.Authorization(user.Id,user.Token)
	if err!=nil{
		util.RespFail(w,err.Error())
		return
	}else{
		user.Token=token
		util.RespOk(w,user)
	}
}


//上传头像
func( ctrl * UserCtrl)Upload(w http.ResponseWriter, req *http.Request) {
	userService.AddTask(w,req)
}
package service

import (


	"../entity"


	"errors"
	"net/http"
	"log"
)

type UserService struct {

}


type UploadInfo struct{
	req *http.Request
	writer http.ResponseWriter
}
var queueupload chan UploadInfo =make(chan UploadInfo)
func (service *UserService) Authorization(id int64,token string) (newtoken string,err error){
	var user entity.Userinfo
	DBengin.ID(id).Get(&user)
	if user.Id==0{
		err = errors.New("缺少ID")
		return
	}
	if user.Token!=token{
		err = errors.New("签名不正确")
		return
	}



	_,err=DBengin.ID(id).Cols("token").Update(user)
	return user.Token,err
}


func StartTask(){
	go func(){
		select {
		case task :=<- queueupload:
			go dispatch(task.writer,task.req);
		}
	}()
	log.Printf("开启任务成功")

}


func (service *UserService) AddTask(w http.ResponseWriter,req *http.Request){
	queueupload <- UploadInfo{
		req:req,
		writer:w,
	}
}

func dispatch(w http.ResponseWriter,req *http.Request){
		//这里进行网络操作
}

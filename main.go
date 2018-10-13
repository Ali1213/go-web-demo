package main

import (

"fmt"


	"flag"
	"./util"


	"./service"

	"./ctrl"
	"net/http"
	"html/template"
	"log"
)

var userctrl ctrl.UserCtrl

func main() {
	fpath  := flag.String("cfgfile","app.conf","config file path")
	_,err:=util.Parse(*fpath)
	if err!=nil{
		fmt.Sprintf("error when %s",err.Error())
		return
	}

	logmap,_ := util.GetSec("log")
	service.InitLog(logmap)


	dbmap,_ := util.GetSec("database")
	service.InitDb(dbmap)
	fmt.Println("Starting application...")



	//注意这里需要开启任务
	service.StartTask()
	appmap,err := util.GetSec("app")
	http.HandleFunc("/user/authorization", userctrl.Authorization)
	http.HandleFunc("/user/upload", userctrl.Upload)
	http.HandleFunc("/", index)

	err = http.ListenAndServe(appmap["addr"], nil)
	if err!=nil{
		log.Print(err.Error())
	}
}


func index(w http.ResponseWriter, r *http.Request){
	t, _ :=template.ParseFiles("view/index.html")
	fmt.Println(t.Name())
	t.Execute(w, "Hello world")
}




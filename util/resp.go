package util

import (

	"net/http"

	"encoding/json"

	"fmt"
)
type H struct {
	Code int	`json:"code"`
	Data interface{} `json:"data,omitempty"`
	Msg string `json:"msg,omitempty"`
	Total interface{} `json:"total,omitempty"`

}
func RespJson(w http.ResponseWriter,data interface{}){


		header :=w.Header()
		header.Set("Content-Type","application/json;charset=utf-8")
		w.WriteHeader(http.StatusOK)
		ret,err :=json.Marshal(data)
		if err!=nil{
			fmt.Println(err.Error())
		}

		w.Write(ret)
}

func RespOk(w http.ResponseWriter,data interface{}){

	RespJson(w,H{
		Code:http.StatusOK,
		Data:data,
	})
}


func RespFail(w http.ResponseWriter,msg string){

	RespJson(w,H{
		Code:http.StatusNotFound,
		Msg :msg,
	})
}

func RespOkList(w http.ResponseWriter,data interface{},total interface{}){

	RespJson(w,H{
		Code:http.StatusOK,
		Data:data,
		Total:total,
	})
}
package service

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"

	"fmt"
	"strings"
	"os"
	"log"
	"strconv"
)
var DBengin *xorm.Engine
func InitDb(dbmap map[string]string){

	driverName     := dbmap["driveName"]
	dataSourceName := dbmap["dataSourceName"]
	showsql :=       dbmap["showSql"]!="false"
	maxIdle,_ :=        strconv.Atoi(dbmap["maxIdle"])
	maxOpen,_ :=       strconv.Atoi(dbmap["maxOpen"])
	dbengin , err := xorm.NewEngine(driverName, dataSourceName)

	if err != nil {
		panic("data source init error ==>"+err.Error())
		return
	}
	dbengin.ShowSQL(showsql)
	dbengin.SetMaxIdleConns(maxIdle)
	dbengin.SetMaxOpenConns(maxOpen)

	dbengin.SetConnMaxLifetime(5*time.Second)
	DBengin = dbengin
}

func InitLog(logmap map[string]string){
	logfile := logmap["output"]
	outfile := os.Stdout

	if "stdout"!=strings.ToLower(logfile){
		outfile, err := os.OpenFile(logfile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666) //打开文件，若果文件不存在就创建一个同名文件并打开
		if err != nil || outfile==nil{
			fmt.Println("error when init log file")
			os.Exit(1)
		}
	}
	log.SetOutput(outfile)
	strflags := logmap["flags"]
	flags := 0

	if (strings.Contains(strflags,"Ldate")){
		flags |= log.Ldate
	}
	if (strings.Contains(strflags,"Ltime")){
		flags |= log.Ltime
	}
	if (strings.Contains(strflags,"Lshortfile")){
		flags |= log.Lshortfile

	}
	if (strings.Contains(strflags,"Lmicroseconds")){
		flags |= log.Lmicroseconds

	}
	if (strings.Contains(strflags,"LUTC")){
		flags |= log.LUTC

	}

	if (strings.Contains(strflags,"LstdFlags")){
		flags |= log.LstdFlags

	}

	log.SetFlags(flags) //设置答应日志每一行前的标志信息，这里设置了日期，打印时间，当前go文件的文件名

}

package util

import (
	"github.com/Unknwon/goconfig"


	"strconv"
)

var cfg *goconfig.ConfigFile
var cfgmap map[string]string = make(map[string]string)
var filepath string
func Parse(fpath string)(c map[string]string ,err error){
	cfg, err := goconfig.LoadConfigFile(fpath)
	filepath = fpath
	sec :=cfg.GetSectionList()
	for _,v :=range sec{
		keys := cfg.GetKeyList(v)
		for _,b:= range keys{
			cfgmap[v+"."+b],_ = cfg.GetValue(v,b)
		}
	}
	return c,err
}

func GetAllCfg()(c map[string]string){
	return cfgmap
}

func GetValue(sec,key string)(string){
	return  cfgmap[sec+"."+key]
}

func GetSec(sec string)(map[string]string,error){
	cfg, _ := goconfig.LoadConfigFile(filepath)
	return  cfg.GetSection(sec)
}
func GetBool(sec,key string)(bool){
	return  cfgmap[sec+"."+key] =="true" || cfgmap[sec+"."+key] =="1"
}

func GetInt(sec,key string)(int){
	 a,_:= strconv.Atoi(cfgmap[sec+"."+key])
	return a
}
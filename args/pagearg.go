package args

import "fmt"

type PageArg struct {
	Pagefrom int  `json:"pagefrom" form:"pagefrom"`
	Pagesize int  `json:"pagesize" form:"pagesize"`
	Kword string  `json:"kword" form:"kword"`
	Asc string 	`json:"asc" form:"asc"`
	Desc string 	`json:"desc" form:"desc"`
}

func (p*PageArg) GetPageSize() int{
	if p.Pagesize==0{
		return 100
	}else{
		return p.Pagesize
	}

}
func (p*PageArg) GetPageFrom() int{
	if p.Pagefrom<0{
		return 0
	}else{
		return p.Pagefrom
	}
}

func (p*PageArg) GetOrderBy() string{
	if len(p.Asc)>0{
		return fmt.Sprintf(" %s asc",p.Asc)
	} else if len(p.Desc)>0{
		return fmt.Sprintf(" %s desc",p.Desc)
	}else{
		return ""
	}
}


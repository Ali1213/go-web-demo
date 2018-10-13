package args

type UserArg struct {
	PageArg
	Mobile string `json:"mobile" form:"mobile"`
}


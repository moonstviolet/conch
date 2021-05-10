package proto

type LoginReq struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type SignupReq struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Email    string `json:"email" form:"email"`
	Nickname string `json:"nickname" form:"nickname"`
	Motto    string `json:"motto" form:"motto"`
}

type FindUserReq struct {
	Username string `json:"username" form:"username"`
}

type FindUserResp struct {
	IsValid bool `json:"isValid"`
}

package proto

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
}

type LogoutReq struct {
}

type LogoutResp struct {
}

type SignupReq struct {
}

type SignupResp struct {
}

type FindUserReq struct {
}

type FindUserResp struct {
}

type ProfileReq struct {
}

type ProfileResp struct {
}

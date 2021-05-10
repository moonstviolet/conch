package proto

import "conch/models"

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LogoutReq struct {
}

type LogoutResp struct {
}

type SignupReq struct {
	User models.User
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

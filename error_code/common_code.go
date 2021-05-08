package error_code

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(1, "服务内部错误")
	InvalidParams             = NewError(2, "参数错误")
	NotFound                  = NewError(3, "找不到资源")
	UnauthorizedAuthNotExit   = NewError(4, "鉴权失败,找不到对应的key")
	UnauthorizedTokenError    = NewError(5, "鉴权失败,Token错误")
	UnauthorizedTokenTimeOut  = NewError(6, "鉴权失败,Token超时")
	UnauthorizedTokenGenerate = NewError(7, "鉴权失败,Token生成失败")
	TooManyRequests           = NewError(8, "请求过多")
)

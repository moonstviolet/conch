package error_code

import "net/http"

type RespError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewError(code int, msg string) *RespError {
	return &RespError{Code: code, Msg: msg}
}

func (re *RespError) StatusCode() int {
	switch re.Code {
	case Success.Code:
		return http.StatusOK
	case ServerError.Code:
		return http.StatusInternalServerError
	case InvalidParams.Code:
		return http.StatusBadRequest
	case UnauthorizedAuthNotExit.Code, UnauthorizedTokenError.Code,
		UnauthorizedTokenTimeOut.Code, UnauthorizedTokenGenerate.Code:
		return http.StatusUnauthorized
	case TooManyRequests.Code:
		return http.StatusTooManyRequests
	default:
		return http.StatusInternalServerError
	}
}

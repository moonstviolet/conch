package middleware

import (
	"conch/error_code"
	"errors"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type RespStruct struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

var (
	EmptyStr = ""

	ErrMustPtr              = errors.New("param must be ptr")
	ErrMustPointToStruct    = errors.New("param must point to struct")
	ErrMustHasThreeParam    = errors.New("method must has three input")
	ErrMustFunc             = errors.New("method must be func")
	ErrMustValid            = errors.New("method must be valid")
	ErrMustResponseErrorPtr = errors.New("method ret must be *error_code.ResponseError")
	ErrMustOneOut           = errors.New("method must has one out")

	ResponseErrorType = reflect.TypeOf((*error_code.RespError)(nil))
)

func checkMethod(method interface{}) (mV reflect.Value, reqT, respT reflect.Type, err error) {
	mV = reflect.ValueOf(method)
	if !mV.IsValid() {
		err = ErrMustValid
		return
	}
	mT := mV.Type()
	if mT.Kind() != reflect.Func {
		err = ErrMustFunc
		return
	}
	if mT.NumIn() != 2 {
		err = ErrMustHasThreeParam
		return
	}
	reqT = mT.In(0)
	if reqT.Kind() != reflect.Ptr {
		err = ErrMustPtr
		return
	}
	reqT = reqT.Elem()
	if reqT.Kind() != reflect.Struct {
		err = ErrMustPointToStruct
		return
	}
	respT = mT.In(1)
	if respT.Kind() != reflect.Ptr {
		err = ErrMustPtr
		return
	}
	respT = respT.Elem()
	if respT.Kind() != reflect.Struct {
		err = ErrMustPointToStruct
		return
	}
	if mT.NumOut() != 1 {
		err = ErrMustOneOut
		return
	}
	retT := mT.Out(0)
	if retT != ResponseErrorType {
		err = ErrMustResponseErrorPtr
		return
	}
	return
}

func CreateHandlerFunc(method interface{}) gin.HandlerFunc {
	mV, reqT, respT, err := checkMethod(method)
	if err != nil {
		panic(err)
	}

	return func(c *gin.Context) {
		req := reflect.New(reqT)
		if err := c.ShouldBind(req.Interface()); err != nil {
			responseError := error_code.NewError(1, err.Error())
			c.JSON(http.StatusBadRequest, responseError)
			return
		}

		resp := reflect.New(respT)
		rets := mV.Call([]reflect.Value{req, resp})
		errValue := rets[0]
		if errValue.Interface() != nil {
			err := errValue.Interface().(*error_code.RespError)
			if err != nil {
				c.JSON(http.StatusOK, errValue.Interface())
				return
			}
		}

		c.PureJSON(http.StatusOK, RespStruct{
			Code: 0,
			Data: resp.Interface(),
		})
	}
}

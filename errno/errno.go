package errno

import (
	"fmt"
	"net/http"
)

var _toStatus = map[int]int{
	Success.Code():               http.StatusOK,
	ErrInternalServer.Code():     http.StatusInternalServerError,
	ErrNotFound.Code():           http.StatusNotFound,
	ErrInvalidParam.Code():       http.StatusBadRequest,
	ErrToken.Code():              http.StatusUnauthorized,
	ErrInvalidToken.Code():       http.StatusUnauthorized,
	ErrTokenTimeout.Code():       http.StatusUnauthorized,
	ErrTooManyRequests.Code():    http.StatusTooManyRequests,
	ErrServiceUnavailable.Code(): http.StatusServiceUnavailable,
}

// Error 返回错误码和消息的结构体
type Error struct {
	code    int      `json:"code"`
	msg     string   `json:"msg"`
	details []string `json:"details"`
}

// NewError 实例化
func NewError(code int, msg string) *Error {
	return &Error{code: code, msg: msg}
}

// Error 获取错误信息
func (e *Error) Error() string {
	return fmt.Sprintf("code：%d, msg:：%s", e.Code(), e.Msg())
}

// Code 获取code
func (e *Error) Code() int {
	return e.code
}

// Msg 获取msg
func (e *Error) Msg() string {
	return e.msg
}

// Details 获取details
func (e *Error) Details() []string {
	return e.details
}

// WithDetails 设置details数据
func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.details = []string{}
	for _, d := range details {
		newError.details = append(newError.details, d)
	}

	return &newError
}

// StatusCode 状态码
func (e *Error) StatusCode() int {
	if c, ok := _toStatus[e.code]; ok {
		return c
	}

	return http.StatusBadRequest
}

// Err represents an error
type Err struct {
	Code    int
	Message string
	Err     error
}

// Error 格式化
func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}

// DecodeErr 对错误进行解码，返回错误code和错误提示
func DecodeErr(err error) (int, string) {
	if err == nil {
		return Success.code, Success.msg
	}

	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *Error:
		return typed.code, typed.msg
	default:
	}

	return ErrInternalServer.Code(), err.Error()
}

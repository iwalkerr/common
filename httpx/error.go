package httpx

import (
	"fmt"

	"github.com/pkg/errors"
)

type CodeError struct {
	errCode ErrorType
	errMsg  string
}

// 返回给前端的错误码
func (e *CodeError) getCode() ErrorType {
	return e.errCode
}

// 返回给前端显示端错误信息
func (e *CodeError) getMsg() string {
	return e.errMsg
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode: %d, ErrMsg: %s", e.errCode, e.errMsg)
}

func NewCodeMsg(errCode ErrorType, errMsg string) *CodeError {
	return &CodeError{errCode: errCode, errMsg: errMsg}
}
func NewCode(errCode ErrorType) *CodeError {
	return &CodeError{errCode: errCode, errMsg: mapErrMsg(errCode)}
}

func NewMsg(errMsg string) *CodeError {
	return &CodeError{errCode: SERVER_COMMON_ERROR, errMsg: errMsg}
}

func mapErrMsg(errcode ErrorType) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return message[SERVER_COMMON_ERROR]
	}
}

func WrapMsgf(err *CodeError, args ...any) error {
	switch len(args) {
	case 0:
		return err
	case 1:
		return errors.Wrapf(err, "ErrMsg => %v", args...)
	default:
		return errors.Wrapf(err, "%+v, ErrMsg => %v", args[0], args[1])
	}
}

func Wrapf(errType ErrorType, args ...any) error {
	err := NewCode(errType)
	switch len(args) {
	case 0:
		return err
	case 1:
		return errors.Wrapf(err, "ErrMsg => %v", args...)
	default:
		return errors.Wrapf(err, "%+v, ErrMsg => %v", args[0], args[1])
	}
}

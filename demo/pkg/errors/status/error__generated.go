// This is a generated source file. DO NOT EDIT
// Source: status/error__generated.go

package status

import (
	"github.com/saitofun/qkit/kit/statusx"
)

var _ statusx.Error = (*Error)(nil)

func (v Error) StatusErr() *statusx.StatusErr {
	return &statusx.StatusErr{
		Key:       v.Key(),
		Code:      v.Code(),
		Msg:       v.Msg(),
		CanBeTalk: v.CanBeTalk(),
	}
}

func (v Error) Unwrap() error {
	return v.StatusErr()
}

func (v Error) Error() string {
	return v.StatusErr().Error()
}

func (v Error) StatusCode() int {
	return statusx.StatusCodeFromCode(int(v))
}

func (v Error) Code() int {
	if with, ok := (interface{})(v).(statusx.ServiceCode); ok {
		return with.ServiceCode() + int(v)
	}
	return int(v)

}

func (v Error) Key() string {
	switch v {
	case Unauthorized:
		return "Unauthorized"
	case Conflict:
		return "Conflict"
	case InternalServerError:
		return "InternalServerError"
	case SomeBusinessErrorAs5xx:
		return "SomeBusinessErrorAs5xx"
	}
	return "UNKNOWN"
}

func (v Error) Msg() string {
	switch v {
	case Unauthorized:
		return "Unauthorized"
	case Conflict:
		return "Conflict conflict error"
	case InternalServerError:
		return "InternalServerError 内部错误"
	case SomeBusinessErrorAs5xx:
		return "业务处理错误"
	}
	return "-"
}

func (v Error) CanBeTalk() bool {
	switch v {
	case Unauthorized:
		return true
	case Conflict:
		return false
	case InternalServerError:
		return false
	case SomeBusinessErrorAs5xx:
		return true
	}
	return false
}

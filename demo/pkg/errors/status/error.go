package status

import (
	"net/http"
)

//go:generate toolkit gen status Error
type Error int

func (Error) ServiceCode() int {
	return 999 * 1e3
}

const (
	// InternalServerError 内部错误
	InternalServerError Error = http.StatusInternalServerError*1e6 + iota + 1
	// @errTalk 业务处理错误
	SomeBusinessErrorAs5xx
)

const (
	// @errTalk Unauthorized
	Unauthorized Error = http.StatusUnauthorized*1e6 + iota + 1
)

const (
	// Conflict conflict error
	Conflict Error = http.StatusConflict*1e6 + iota + 1
)

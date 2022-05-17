package dep_enums

import (
	"bytes"
	"database/sql/driver"
	"errors"

	"github.com/sincospro/qkit/gen/enum"
)

var InvalidRuntimeType = errors.New("invalid RuntimeType type")

func ParseRuntimeTypeFromLabelString(s string) (RuntimeType, error) {
	switch s {
	case "":
		return RUNTIME_TYPE_UNKNOWN, nil
	case "静态可执行资源":
		return RUNTIME_TYPE__STATIC, nil
	case "后台服务":
		return RUNTIME_TYPE__DAEMON, nil
	}
	return RUNTIME_TYPE_UNKNOWN, InvalidRuntimeType
}

func (v RuntimeType) String() string {
	switch v {
	case RUNTIME_TYPE_UNKNOWN:
		return ""
	case RUNTIME_TYPE__STATIC:
		return "STATIC"
	case RUNTIME_TYPE__DAEMON:
		return "DAEMON"
	}
	return "UNKNOWN"
}

func ParseRuntimeTypeFromString(s string) (RuntimeType, error) {
	switch s {
	case "":
		return RUNTIME_TYPE_UNKNOWN, nil
	case "STATIC":
		return RUNTIME_TYPE__STATIC, nil
	case "DAEMON":
		return RUNTIME_TYPE__DAEMON, nil
	}
	return RUNTIME_TYPE_UNKNOWN, InvalidRuntimeType
}

func (v RuntimeType) Label() string {
	switch v {
	case RUNTIME_TYPE_UNKNOWN:
		return ""
	case RUNTIME_TYPE__STATIC:
		return "静态可执行资源"
	case RUNTIME_TYPE__DAEMON:
		return "后台服务"
	}
	return "UNKNOWN"
}

func (v RuntimeType) Int() int {
	return int(v)
}

func (RuntimeType) TypeName() string {
	return "github.com/sincospro/qkit/conf/dep/dep_enums.RuntimeType"
}

func (RuntimeType) ConstValues() []enum.IntStringerEnum {
	return []enum.IntStringerEnum{RUNTIME_TYPE__STATIC, RUNTIME_TYPE__DAEMON}
}

func (v RuntimeType) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidRuntimeType
	}
	return []byte(str), nil
}

func (v *RuntimeType) UnmarshalText(data []byte) (err error) {
	*v, err = ParseRuntimeTypeFromString(string(bytes.ToUpper(data)))
	return
}

func (v RuntimeType) Value() (driver.Value, error) {
	offset := 0
	if o, ok := (interface{})(v).(enum.ValueOffset); ok {
		offset = o.Offset()
	}
	return int64(v) + int64(offset), nil
}

func (v *RuntimeType) Scan(src interface{}) error {
	offset := 0
	if o, ok := (interface{})(v).(enum.ValueOffset); ok {
		offset = o.Offset()
	}

	i, err := enum.ScanIntEnumStringer(src, offset)
	if err != nil {
		return err
	}
	*v = RuntimeType(i)
	return nil
}

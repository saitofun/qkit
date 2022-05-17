package dep_enums

import (
	"bytes"
	"database/sql/driver"
	"errors"

	"github.com/sincospro/qkit/gen/enum"
)

var InvalidRuntimeMethod = errors.New("invalid RuntimeMethod type")

func ParseRuntimeMethodFromLabelString(s string) (RuntimeMethod, error) {
	switch s {
	case "":
		return RUNTIME_METHOD_UNKNOWN, nil
	case "SUPERVISOR":
		return RUNTIME_METHOD__SUPERVISOR, nil
	case "DOCKER":
		return RUNTIME_METHOD__DOCKER, nil
	}
	return RUNTIME_METHOD_UNKNOWN, InvalidRuntimeMethod
}

func (v RuntimeMethod) String() string {
	switch v {
	case RUNTIME_METHOD_UNKNOWN:
		return ""
	case RUNTIME_METHOD__SUPERVISOR:
		return "SUPERVISOR"
	case RUNTIME_METHOD__DOCKER:
		return "DOCKER"
	}
	return "UNKNOWN"
}

func ParseRuntimeMethodFromString(s string) (RuntimeMethod, error) {
	switch s {
	case "":
		return RUNTIME_METHOD_UNKNOWN, nil
	case "SUPERVISOR":
		return RUNTIME_METHOD__SUPERVISOR, nil
	case "DOCKER":
		return RUNTIME_METHOD__DOCKER, nil
	}
	return RUNTIME_METHOD_UNKNOWN, InvalidRuntimeMethod
}

func (v RuntimeMethod) Label() string {
	switch v {
	case RUNTIME_METHOD_UNKNOWN:
		return ""
	case RUNTIME_METHOD__SUPERVISOR:
		return "SUPERVISOR"
	case RUNTIME_METHOD__DOCKER:
		return "DOCKER"
	}
	return "UNKNOWN"
}

func (v RuntimeMethod) Int() int {
	return int(v)
}

func (RuntimeMethod) TypeName() string {
	return "github.com/sincospro/qkit/conf/dep/dep_enums.RuntimeMethod"
}

func (RuntimeMethod) ConstValues() []enum.IntStringerEnum {
	return []enum.IntStringerEnum{RUNTIME_METHOD__SUPERVISOR, RUNTIME_METHOD__DOCKER}
}

func (v RuntimeMethod) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidRuntimeMethod
	}
	return []byte(str), nil
}

func (v *RuntimeMethod) UnmarshalText(data []byte) (err error) {
	*v, err = ParseRuntimeMethodFromString(string(bytes.ToUpper(data)))
	return
}

func (v RuntimeMethod) Value() (driver.Value, error) {
	offset := 0
	if o, ok := (interface{})(v).(enum.ValueOffset); ok {
		offset = o.Offset()
	}
	return int64(v) + int64(offset), nil
}

func (v *RuntimeMethod) Scan(src interface{}) error {
	offset := 0
	if o, ok := (interface{})(v).(enum.ValueOffset); ok {
		offset = o.Offset()
	}

	i, err := enum.ScanIntEnumStringer(src, offset)
	if err != nil {
		return err
	}
	*v = RuntimeMethod(i)
	return nil
}

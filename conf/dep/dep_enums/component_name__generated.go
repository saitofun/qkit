package dep_enums

import (
	"bytes"
	"database/sql/driver"
	"errors"

	"github.com/sincospro/qkit/gen/enum"
)

var InvalidComponentName = errors.New("invalid ComponentName type")

func ParseComponentNameFromLabelString(s string) (ComponentName, error) {
	switch s {
	case "":
		return COMPONENT_NAME_UNKNOWN, nil
	case "REDIS":
		return COMPONENT_NAME__REDIS, nil
	}
	return COMPONENT_NAME_UNKNOWN, InvalidComponentName
}

func (v ComponentName) String() string {
	switch v {
	case COMPONENT_NAME_UNKNOWN:
		return ""
	case COMPONENT_NAME__REDIS:
		return "REDIS"
	}
	return "UNKNOWN"
}

func ParseComponentNameFromString(s string) (ComponentName, error) {
	switch s {
	case "":
		return COMPONENT_NAME_UNKNOWN, nil
	case "REDIS":
		return COMPONENT_NAME__REDIS, nil
	}
	return COMPONENT_NAME_UNKNOWN, InvalidComponentName
}

func (v ComponentName) Label() string {
	switch v {
	case COMPONENT_NAME_UNKNOWN:
		return ""
	case COMPONENT_NAME__REDIS:
		return "REDIS"
	}
	return "UNKNOWN"
}

func (v ComponentName) Int() int {
	return int(v)
}

func (ComponentName) TypeName() string {
	return "github.com/sincospro/qkit/conf/dep/dep_enums.ComponentName"
}

func (ComponentName) ConstValues() []enum.IntStringerEnum {
	return []enum.IntStringerEnum{COMPONENT_NAME__REDIS}
}

func (v ComponentName) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidComponentName
	}
	return []byte(str), nil
}

func (v *ComponentName) UnmarshalText(data []byte) (err error) {
	*v, err = ParseComponentNameFromString(string(bytes.ToUpper(data)))
	return
}

func (v ComponentName) Value() (driver.Value, error) {
	offset := 0
	if o, ok := (interface{})(v).(enum.ValueOffset); ok {
		offset = o.Offset()
	}
	return int64(v) + int64(offset), nil
}

func (v *ComponentName) Scan(src interface{}) error {
	offset := 0
	if o, ok := (interface{})(v).(enum.ValueOffset); ok {
		offset = o.Offset()
	}

	i, err := enum.ScanIntEnumStringer(src, offset)
	if err != nil {
		return err
	}
	*v = ComponentName(i)
	return nil
}

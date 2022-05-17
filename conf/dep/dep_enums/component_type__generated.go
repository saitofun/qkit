package dep_enums

import (
	"bytes"
	"database/sql/driver"
	"errors"

	"github.com/sincospro/qkit/gen/enum"
)

var InvalidComponentType = errors.New("invalid ComponentType type")

func ParseComponentTypeFromLabelString(s string) (ComponentType, error) {
	switch s {
	case "":
		return COMPONENT_TYPE_UNKNOWN, nil
	case "必选":
		return COMPONENT_TYPE__REQUIRED, nil
	case "可选":
		return COMPONENT_TYPE__OPTIONAL, nil
	}
	return COMPONENT_TYPE_UNKNOWN, InvalidComponentType
}

func (v ComponentType) String() string {
	switch v {
	case COMPONENT_TYPE_UNKNOWN:
		return ""
	case COMPONENT_TYPE__REQUIRED:
		return "REQUIRED"
	case COMPONENT_TYPE__OPTIONAL:
		return "OPTIONAL"
	}
	return "UNKNOWN"
}

func ParseComponentTypeFromString(s string) (ComponentType, error) {
	switch s {
	case "":
		return COMPONENT_TYPE_UNKNOWN, nil
	case "REQUIRED":
		return COMPONENT_TYPE__REQUIRED, nil
	case "OPTIONAL":
		return COMPONENT_TYPE__OPTIONAL, nil
	}
	return COMPONENT_TYPE_UNKNOWN, InvalidComponentType
}

func (v ComponentType) Label() string {
	switch v {
	case COMPONENT_TYPE_UNKNOWN:
		return ""
	case COMPONENT_TYPE__REQUIRED:
		return "必选"
	case COMPONENT_TYPE__OPTIONAL:
		return "可选"
	}
	return "UNKNOWN"
}

func (v ComponentType) Int() int {
	return int(v)
}

func (ComponentType) TypeName() string {
	return "github.com/sincospro/qkit/conf/dep/dep_enums.ComponentType"
}

func (ComponentType) ConstValues() []enum.IntStringerEnum {
	return []enum.IntStringerEnum{COMPONENT_TYPE__REQUIRED, COMPONENT_TYPE__OPTIONAL}
}

func (v ComponentType) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidComponentType
	}
	return []byte(str), nil
}

func (v *ComponentType) UnmarshalText(data []byte) (err error) {
	*v, err = ParseComponentTypeFromString(string(bytes.ToUpper(data)))
	return
}

func (v ComponentType) Value() (driver.Value, error) {
	offset := 0
	if o, ok := (interface{})(v).(enum.ValueOffset); ok {
		offset = o.Offset()
	}
	return int64(v) + int64(offset), nil
}

func (v *ComponentType) Scan(src interface{}) error {
	offset := 0
	if o, ok := (interface{})(v).(enum.ValueOffset); ok {
		offset = o.Offset()
	}

	i, err := enum.ScanIntEnumStringer(src, offset)
	if err != nil {
		return err
	}
	*v = ComponentType(i)
	return nil
}

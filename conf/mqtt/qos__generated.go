package mqtt

import (
	bytes "bytes"
	database_sql_driver "database/sql/driver"
	errors "errors"

	github_com_go_courier_enumeration "github.com/go-courier/enumeration"
)

var InvalidQOS = errors.New("invalid QOS type")

func ParseQOSFromLabelString(s string) (QOS, error) {
	switch s {
	case "":
		return QOS_UNKNOWN, nil
	case "0":
		return QOS__ONCE, nil
	case "1":
		return QOS__AT_LEAST_ONCE, nil
	case "2":
		return QOS__ONLY_ONCE, nil
	}
	return QOS_UNKNOWN, InvalidQOS
}

func (v QOS) String() string {
	switch v {
	case QOS_UNKNOWN:
		return ""
	case QOS__ONCE:
		return "ONCE"
	case QOS__AT_LEAST_ONCE:
		return "AT_LEAST_ONCE"
	case QOS__ONLY_ONCE:
		return "ONLY_ONCE"
	}
	return "UNKNOWN"
}

func ParseQOSFromString(s string) (QOS, error) {
	switch s {
	case "":
		return QOS_UNKNOWN, nil
	case "ONCE":
		return QOS__ONCE, nil
	case "AT_LEAST_ONCE":
		return QOS__AT_LEAST_ONCE, nil
	case "ONLY_ONCE":
		return QOS__ONLY_ONCE, nil
	}
	return QOS_UNKNOWN, InvalidQOS
}

func (v QOS) Label() string {
	switch v {
	case QOS_UNKNOWN:
		return ""
	case QOS__ONCE:
		return "0"
	case QOS__AT_LEAST_ONCE:
		return "1"
	case QOS__ONLY_ONCE:
		return "2"
	}
	return "UNKNOWN"
}

func (v QOS) Int() int {
	return int(v)
}

func (QOS) TypeName() string {
	return "github.com/saitofun/qkit/conf/mqtt.QOS"
}

func (QOS) ConstValues() []github_com_go_courier_enumeration.IntStringerEnum {
	return []github_com_go_courier_enumeration.IntStringerEnum{QOS__ONCE, QOS__AT_LEAST_ONCE, QOS__ONLY_ONCE}
}

func (v QOS) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidQOS
	}
	return []byte(str), nil
}

func (v *QOS) UnmarshalText(data []byte) (err error) {
	*v, err = ParseQOSFromString(string(bytes.ToUpper(data)))
	return
}

func (v QOS) Value() (database_sql_driver.Value, error) {
	offset := 0
	if o, ok := (interface{})(v).(github_com_go_courier_enumeration.DriverValueOffset); ok {
		offset = o.Offset()
	}
	return int64(v) + int64(offset), nil
}

func (v *QOS) Scan(src interface{}) error {
	offset := 0
	if o, ok := (interface{})(v).(github_com_go_courier_enumeration.DriverValueOffset); ok {
		offset = o.Offset()
	}

	i, err := github_com_go_courier_enumeration.ScanIntEnumStringer(src, offset)
	if err != nil {
		return err
	}
	*v = QOS(i)
	return nil
}

package models

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/saitofun/qkit/base/types"
	"github.com/saitofun/qkit/kit/sqlx"
)

var Demo = sqlx.NewDatabase("demo")

type OperationTimes struct {
	CreatedAt types.Timestamp `db:"f_created_at,default='0'" json:"createdAt"`
	UpdatedAt types.Timestamp `db:"f_updated_at,default='0'" json:"updatedAt"`
}

type PrimaryID struct {
	ID uint64 `db:"f_id,autoincrement" json:"-"`
}

func JSONScan(dbv interface{}, v interface{}) error {
	switch val := dbv.(type) {
	case []byte:
		if len(val) == 0 {
			return nil
		}
		return json.Unmarshal(val, v)
	case string:
		if val == "" {
			return nil
		}
		return json.Unmarshal([]byte(val), v)
	case nil:
		return nil
	default:
		return errors.Errorf("cannot sql.Scan() from `%#v`", v)
	}
}

func JSONValue(v interface{}) (driver.Value, error) {
	if v == nil {
		return "", nil
	}
	if zero, ok := v.(interface{ IsZero() bool }); ok && zero.IsZero() {
		return "", nil
	}
	bytes, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	str := string(bytes)
	if str == "null" {
		str = ""
	}
	return str, nil
}

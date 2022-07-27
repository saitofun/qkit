package models

import (
	"database/sql"
	"database/sql/driver"
)

//go:generate toolkit gen model Model --database Demo
// Model database model demo
// @def primary                     ID
// @def index        I_field_string FieldString
// @def unique_index UI_model_id    ModelID
// @def unique_index UI_union_idx   UnionIndexField1 UnionIndexField2
type Model struct {
	PrimaryID
	RefModel
	BaseModel
	OperationTimes
}

type RefModel struct {
	ModelID uint64 `db:"f_model_id" json:"modelID"`
}

type BaseModel struct {
	FieldString      string      `db:"f_field_string,default=''"       json:"fieldString"`
	FieldJsonContent *EmbedModel `db:"f_field_json_content,default=''" json:"fieldJsonContent"`
	UnionIndexField1 int64       `db:"f_union_idx_f1,default='0'"      json:"-"`
	UnionIndexField2 string      `db:"f_union_idx_f2,default=''"       json:"-"`
}

// EmbedModel implemented driver.Valuer and sql.Scanner and DataType for database store and scan
type EmbedModel struct {
	SomeFieldInt   int     `json:"int,omitempty"`
	SomeFieldFloat float64 `json:"float,omitempty"`
}

var (
	_ driver.Valuer = EmbedModel{}
	_ sql.Scanner   = (*EmbedModel)(nil)
)

func (EmbedModel) DataType(driver string) string { return "TEXT" }

func (v EmbedModel) Value() (driver.Value, error) { return JSONValue(v) }

func (v *EmbedModel) Scan(src interface{}) error { return JSONScan(src, v) }

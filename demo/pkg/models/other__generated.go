// This is a generated source file. DO NOT EDIT
// Source: models/other__generated.go

package models

import (
	"fmt"
	"time"

	"github.com/saitofun/qkit/base/types"
	"github.com/saitofun/qkit/kit/sqlx"
	"github.com/saitofun/qkit/kit/sqlx/builder"
)

var OtherTable *builder.Table

func init() {
	OtherTable = Demo.Register(&Other{})
}

type OtherIterator struct {
}

func (OtherIterator) New() interface{} {
	return &Other{}
}

func (OtherIterator) Resolve(v interface{}) *Other {
	return v.(*Other)
}

func (Other) TableName() string {
	return "t_other"
}

func (Other) TableDesc() []string {
	return []string{
		"Other database model demo",
	}
}

func (Other) Comments() map[string]string {
	return map[string]string{}
}

func (Other) ColDesc() map[string][]string {
	return map[string][]string{}
}

func (Other) ColRel() map[string][]string {
	return map[string][]string{}
}

func (Other) PrimaryKey() []string {
	return []string{
		"ID",
	}
}

func (Other) Indexes() builder.Indexes {
	return builder.Indexes{
		"i_field_string": []string{
			"FieldString",
		},
	}
}

func (m *Other) IndexFieldNames() []string {
	return []string{
		"FieldString",
		"ID",
		"ModelID",
		"UnionIndexField1",
		"UnionIndexField2",
	}
}

func (Other) UniqueIndexes() builder.Indexes {
	return builder.Indexes{
		"ui_model_id": []string{
			"ModelID",
		},
		"ui_union_idx": []string{
			"UnionIndexField1",
			"UnionIndexField2",
		},
	}
}

func (Other) UniqueIndexUiModelId() string {
	return "ui_model_id"
}

func (Other) UniqueIndexUiUnionIdx() string {
	return "ui_union_idx"
}

func (m *Other) ColID() *builder.Column {
	return OtherTable.ColByFieldName(m.FieldID())
}

func (Other) FieldID() string {
	return "ID"
}

func (m *Other) ColModelID() *builder.Column {
	return OtherTable.ColByFieldName(m.FieldModelID())
}

func (Other) FieldModelID() string {
	return "ModelID"
}

func (m *Other) ColFieldString() *builder.Column {
	return OtherTable.ColByFieldName(m.FieldFieldString())
}

func (Other) FieldFieldString() string {
	return "FieldString"
}

func (m *Other) ColFieldJsonContent() *builder.Column {
	return OtherTable.ColByFieldName(m.FieldFieldJsonContent())
}

func (Other) FieldFieldJsonContent() string {
	return "FieldJsonContent"
}

func (m *Other) ColUnionIndexField1() *builder.Column {
	return OtherTable.ColByFieldName(m.FieldUnionIndexField1())
}

func (Other) FieldUnionIndexField1() string {
	return "UnionIndexField1"
}

func (m *Other) ColUnionIndexField2() *builder.Column {
	return OtherTable.ColByFieldName(m.FieldUnionIndexField2())
}

func (Other) FieldUnionIndexField2() string {
	return "UnionIndexField2"
}

func (m *Other) ColCreatedAt() *builder.Column {
	return OtherTable.ColByFieldName(m.FieldCreatedAt())
}

func (Other) FieldCreatedAt() string {
	return "CreatedAt"
}

func (m *Other) ColUpdatedAt() *builder.Column {
	return OtherTable.ColByFieldName(m.FieldUpdatedAt())
}

func (Other) FieldUpdatedAt() string {
	return "UpdatedAt"
}

func (m *Other) CondByValue(db sqlx.DBExecutor) builder.SqlCondition {
	var (
		tbl  = db.T(m)
		fvs  = builder.FieldValueFromStructByNoneZero(m)
		cond = make([]builder.SqlCondition, 0)
	)

	for _, fn := range m.IndexFieldNames() {
		if v, ok := fvs[fn]; ok {
			cond = append(cond, tbl.ColByFieldName(fn).Eq(v))
			delete(fvs, fn)
		}
	}
	if len(cond) == 0 {
		panic(fmt.Errorf("no field for indexes has value"))
	}
	for fn, v := range fvs {
		cond = append(cond, tbl.ColByFieldName(fn).Eq(v))
	}
	return builder.And(cond...)
}

func (m *Other) Create(db sqlx.DBExecutor) error {

	if m.CreatedAt.IsZero() {
		m.CreatedAt.Set(time.Now())
	}

	if m.UpdatedAt.IsZero() {
		m.UpdatedAt.Set(time.Now())
	}

	_, err := db.Exec(sqlx.InsertToDB(db, m, nil))
	return err
}

func (m *Other) List(db sqlx.DBExecutor, cond builder.SqlCondition, adds ...builder.Addition) ([]Other, error) {
	var (
		tbl = db.T(m)
		lst = make([]Other, 0)
	)
	adds = append([]builder.Addition{builder.Where(cond), builder.Comment("Other.List")}, adds...)
	err := db.QueryAndScan(builder.Select(nil).From(tbl, adds...), &lst)
	return lst, err
}

func (m *Other) Count(db sqlx.DBExecutor, cond builder.SqlCondition, adds ...builder.Addition) (cnt int64, err error) {
	tbl := db.T(m)
	adds = append([]builder.Addition{builder.Where(cond), builder.Comment("Other.List")}, adds...)
	err = db.QueryAndScan(builder.Select(builder.Count()).From(tbl, adds...), &cnt)
	return
}

func (m *Other) FetchByID(db sqlx.DBExecutor) error {
	tbl := db.T(m)
	err := db.QueryAndScan(
		builder.Select(nil).
			From(
				tbl,
				builder.Where(
					builder.And(
						tbl.ColByFieldName("ID").Eq(m.ID),
					),
				),
				builder.Comment("Other.FetchByID"),
			),
		m,
	)
	return err
}

func (m *Other) FetchByModelID(db sqlx.DBExecutor) error {
	tbl := db.T(m)
	err := db.QueryAndScan(
		builder.Select(nil).
			From(
				tbl,
				builder.Where(
					builder.And(
						tbl.ColByFieldName("ModelID").Eq(m.ModelID),
					),
				),
				builder.Comment("Other.FetchByModelID"),
			),
		m,
	)
	return err
}

func (m *Other) FetchByUnionIndexField1AndUnionIndexField2(db sqlx.DBExecutor) error {
	tbl := db.T(m)
	err := db.QueryAndScan(
		builder.Select(nil).
			From(
				tbl,
				builder.Where(
					builder.And(
						tbl.ColByFieldName("UnionIndexField1").Eq(m.UnionIndexField1),
						tbl.ColByFieldName("UnionIndexField2").Eq(m.UnionIndexField2),
					),
				),
				builder.Comment("Other.FetchByUnionIndexField1AndUnionIndexField2"),
			),
		m,
	)
	return err
}

func (m *Other) UpdateByIDWithFVs(db sqlx.DBExecutor, fvs builder.FieldValues) error {

	if _, ok := fvs["UpdatedAt"]; !ok {
		fvs["UpdatedAt"] = types.Timestamp{Time: time.Now()}
	}
	tbl := db.T(m)
	res, err := db.Exec(
		builder.Update(tbl).
			Where(
				builder.And(
					tbl.ColByFieldName("ID").Eq(m.ID),
				),
				builder.Comment("Other.UpdateByIDWithFVs"),
			).
			Set(tbl.AssignmentsByFieldValues(fvs)...),
	)
	if err != nil {
		return err
	}
	if affected, _ := res.RowsAffected(); affected == 0 {
		return m.FetchByID(db)
	}
	return nil
}

func (m *Other) UpdateByID(db sqlx.DBExecutor, zeros ...string) error {
	fvs := builder.FieldValueFromStructByNoneZero(m, zeros...)
	return m.UpdateByIDWithFVs(db, fvs)
}

func (m *Other) UpdateByModelIDWithFVs(db sqlx.DBExecutor, fvs builder.FieldValues) error {

	if _, ok := fvs["UpdatedAt"]; !ok {
		fvs["UpdatedAt"] = types.Timestamp{Time: time.Now()}
	}
	tbl := db.T(m)
	res, err := db.Exec(
		builder.Update(tbl).
			Where(
				builder.And(
					tbl.ColByFieldName("ModelID").Eq(m.ModelID),
				),
				builder.Comment("Other.UpdateByModelIDWithFVs"),
			).
			Set(tbl.AssignmentsByFieldValues(fvs)...),
	)
	if err != nil {
		return err
	}
	if affected, _ := res.RowsAffected(); affected == 0 {
		return m.FetchByModelID(db)
	}
	return nil
}

func (m *Other) UpdateByModelID(db sqlx.DBExecutor, zeros ...string) error {
	fvs := builder.FieldValueFromStructByNoneZero(m, zeros...)
	return m.UpdateByModelIDWithFVs(db, fvs)
}

func (m *Other) UpdateByUnionIndexField1AndUnionIndexField2WithFVs(db sqlx.DBExecutor, fvs builder.FieldValues) error {

	if _, ok := fvs["UpdatedAt"]; !ok {
		fvs["UpdatedAt"] = types.Timestamp{Time: time.Now()}
	}
	tbl := db.T(m)
	res, err := db.Exec(
		builder.Update(tbl).
			Where(
				builder.And(
					tbl.ColByFieldName("UnionIndexField1").Eq(m.UnionIndexField1),
					tbl.ColByFieldName("UnionIndexField2").Eq(m.UnionIndexField2),
				),
				builder.Comment("Other.UpdateByUnionIndexField1AndUnionIndexField2WithFVs"),
			).
			Set(tbl.AssignmentsByFieldValues(fvs)...),
	)
	if err != nil {
		return err
	}
	if affected, _ := res.RowsAffected(); affected == 0 {
		return m.FetchByUnionIndexField1AndUnionIndexField2(db)
	}
	return nil
}

func (m *Other) UpdateByUnionIndexField1AndUnionIndexField2(db sqlx.DBExecutor, zeros ...string) error {
	fvs := builder.FieldValueFromStructByNoneZero(m, zeros...)
	return m.UpdateByUnionIndexField1AndUnionIndexField2WithFVs(db, fvs)
}

func (m *Other) Delete(db sqlx.DBExecutor) error {
	_, err := db.Exec(
		builder.Delete().
			From(
				db.T(m),
				builder.Where(m.CondByValue(db)),
				builder.Comment("Other.Delete"),
			),
	)
	return err
}

func (m *Other) DeleteByID(db sqlx.DBExecutor) error {
	tbl := db.T(m)
	_, err := db.Exec(
		builder.Delete().
			From(
				tbl,
				builder.Where(
					builder.And(
						tbl.ColByFieldName("ID").Eq(m.ID),
					),
				),
				builder.Comment("Other.DeleteByID"),
			),
	)
	return err
}

func (m *Other) DeleteByModelID(db sqlx.DBExecutor) error {
	tbl := db.T(m)
	_, err := db.Exec(
		builder.Delete().
			From(
				tbl,
				builder.Where(
					builder.And(
						tbl.ColByFieldName("ModelID").Eq(m.ModelID),
					),
				),
				builder.Comment("Other.DeleteByModelID"),
			),
	)
	return err
}

func (m *Other) DeleteByUnionIndexField1AndUnionIndexField2(db sqlx.DBExecutor) error {
	tbl := db.T(m)
	_, err := db.Exec(
		builder.Delete().
			From(
				tbl,
				builder.Where(
					builder.And(
						tbl.ColByFieldName("UnionIndexField1").Eq(m.UnionIndexField1),
						tbl.ColByFieldName("UnionIndexField2").Eq(m.UnionIndexField2),
					),
				),
				builder.Comment("Other.DeleteByUnionIndexField1AndUnionIndexField2"),
			),
	)
	return err
}

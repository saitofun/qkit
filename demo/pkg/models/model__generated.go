// This is a generated source file. DO NOT EDIT
// Source: models/model__generated.go

package models

import (
	"fmt"
	"time"

	"github.com/saitofun/qkit/base/types"
	"github.com/saitofun/qkit/kit/sqlx"
	"github.com/saitofun/qkit/kit/sqlx/builder"
)

var ModelTable *builder.Table

func init() {
	ModelTable = Demo.Register(&Model{})
}

type ModelIterator struct {
}

func (ModelIterator) New() interface{} {
	return &Model{}
}

func (ModelIterator) Resolve(v interface{}) *Model {
	return v.(*Model)
}

func (Model) TableName() string {
	return "t_model"
}

func (Model) TableDesc() []string {
	return []string{
		"Model database model demo",
	}
}

func (Model) Comments() map[string]string {
	return map[string]string{}
}

func (Model) ColDesc() map[string][]string {
	return map[string][]string{}
}

func (Model) ColRel() map[string][]string {
	return map[string][]string{}
}

func (Model) PrimaryKey() []string {
	return []string{
		"ID",
	}
}

func (Model) Indexes() builder.Indexes {
	return builder.Indexes{
		"i_field_string": []string{
			"FieldString",
		},
	}
}

func (m *Model) IndexFieldNames() []string {
	return []string{
		"FieldString",
		"ID",
		"ModelID",
		"UnionIndexField1",
		"UnionIndexField2",
	}
}

func (Model) UniqueIndexes() builder.Indexes {
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

func (Model) UniqueIndexUiModelId() string {
	return "ui_model_id"
}

func (Model) UniqueIndexUiUnionIdx() string {
	return "ui_union_idx"
}

func (m *Model) ColID() *builder.Column {
	return ModelTable.ColByFieldName(m.FieldID())
}

func (Model) FieldID() string {
	return "ID"
}

func (m *Model) ColModelID() *builder.Column {
	return ModelTable.ColByFieldName(m.FieldModelID())
}

func (Model) FieldModelID() string {
	return "ModelID"
}

func (m *Model) ColFieldString() *builder.Column {
	return ModelTable.ColByFieldName(m.FieldFieldString())
}

func (Model) FieldFieldString() string {
	return "FieldString"
}

func (m *Model) ColFieldJsonContent() *builder.Column {
	return ModelTable.ColByFieldName(m.FieldFieldJsonContent())
}

func (Model) FieldFieldJsonContent() string {
	return "FieldJsonContent"
}

func (m *Model) ColUnionIndexField1() *builder.Column {
	return ModelTable.ColByFieldName(m.FieldUnionIndexField1())
}

func (Model) FieldUnionIndexField1() string {
	return "UnionIndexField1"
}

func (m *Model) ColUnionIndexField2() *builder.Column {
	return ModelTable.ColByFieldName(m.FieldUnionIndexField2())
}

func (Model) FieldUnionIndexField2() string {
	return "UnionIndexField2"
}

func (m *Model) ColCreatedAt() *builder.Column {
	return ModelTable.ColByFieldName(m.FieldCreatedAt())
}

func (Model) FieldCreatedAt() string {
	return "CreatedAt"
}

func (m *Model) ColUpdatedAt() *builder.Column {
	return ModelTable.ColByFieldName(m.FieldUpdatedAt())
}

func (Model) FieldUpdatedAt() string {
	return "UpdatedAt"
}

func (m *Model) CondByValue(db sqlx.DBExecutor) builder.SqlCondition {
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

func (m *Model) Create(db sqlx.DBExecutor) error {

	if m.CreatedAt.IsZero() {
		m.CreatedAt.Set(time.Now())
	}

	if m.UpdatedAt.IsZero() {
		m.UpdatedAt.Set(time.Now())
	}

	_, err := db.Exec(sqlx.InsertToDB(db, m, nil))
	return err
}

func (m *Model) List(db sqlx.DBExecutor, cond builder.SqlCondition, adds ...builder.Addition) ([]Model, error) {
	var (
		tbl = db.T(m)
		lst = make([]Model, 0)
	)
	adds = append([]builder.Addition{builder.Where(cond), builder.Comment("Model.List")}, adds...)
	err := db.QueryAndScan(builder.Select(nil).From(tbl, adds...), &lst)
	return lst, err
}

func (m *Model) Count(db sqlx.DBExecutor, cond builder.SqlCondition, adds ...builder.Addition) (cnt int64, err error) {
	tbl := db.T(m)
	adds = append([]builder.Addition{builder.Where(cond), builder.Comment("Model.List")}, adds...)
	err = db.QueryAndScan(builder.Select(builder.Count()).From(tbl, adds...), &cnt)
	return
}

func (m *Model) FetchByID(db sqlx.DBExecutor) error {
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
				builder.Comment("Model.FetchByID"),
			),
		m,
	)
	return err
}

func (m *Model) FetchByModelID(db sqlx.DBExecutor) error {
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
				builder.Comment("Model.FetchByModelID"),
			),
		m,
	)
	return err
}

func (m *Model) FetchByUnionIndexField1AndUnionIndexField2(db sqlx.DBExecutor) error {
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
				builder.Comment("Model.FetchByUnionIndexField1AndUnionIndexField2"),
			),
		m,
	)
	return err
}

func (m *Model) UpdateByIDWithFVs(db sqlx.DBExecutor, fvs builder.FieldValues) error {

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
				builder.Comment("Model.UpdateByIDWithFVs"),
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

func (m *Model) UpdateByID(db sqlx.DBExecutor, zeros ...string) error {
	fvs := builder.FieldValueFromStructByNoneZero(m, zeros...)
	return m.UpdateByIDWithFVs(db, fvs)
}

func (m *Model) UpdateByModelIDWithFVs(db sqlx.DBExecutor, fvs builder.FieldValues) error {

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
				builder.Comment("Model.UpdateByModelIDWithFVs"),
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

func (m *Model) UpdateByModelID(db sqlx.DBExecutor, zeros ...string) error {
	fvs := builder.FieldValueFromStructByNoneZero(m, zeros...)
	return m.UpdateByModelIDWithFVs(db, fvs)
}

func (m *Model) UpdateByUnionIndexField1AndUnionIndexField2WithFVs(db sqlx.DBExecutor, fvs builder.FieldValues) error {

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
				builder.Comment("Model.UpdateByUnionIndexField1AndUnionIndexField2WithFVs"),
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

func (m *Model) UpdateByUnionIndexField1AndUnionIndexField2(db sqlx.DBExecutor, zeros ...string) error {
	fvs := builder.FieldValueFromStructByNoneZero(m, zeros...)
	return m.UpdateByUnionIndexField1AndUnionIndexField2WithFVs(db, fvs)
}

func (m *Model) Delete(db sqlx.DBExecutor) error {
	_, err := db.Exec(
		builder.Delete().
			From(
				db.T(m),
				builder.Where(m.CondByValue(db)),
				builder.Comment("Model.Delete"),
			),
	)
	return err
}

func (m *Model) DeleteByID(db sqlx.DBExecutor) error {
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
				builder.Comment("Model.DeleteByID"),
			),
	)
	return err
}

func (m *Model) DeleteByModelID(db sqlx.DBExecutor) error {
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
				builder.Comment("Model.DeleteByModelID"),
			),
	)
	return err
}

func (m *Model) DeleteByUnionIndexField1AndUnionIndexField2(db sqlx.DBExecutor) error {
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
				builder.Comment("Model.DeleteByUnionIndexField1AndUnionIndexField2"),
			),
	)
	return err
}

package models

import (
	fmt "fmt"
	time "time"

	github_com_go_courier_sqlx_v2 "github.com/go-courier/sqlx/v2"
	github_com_go_courier_sqlx_v2_builder "github.com/go-courier/sqlx/v2/builder"
	github_com_go_courier_sqlx_v2_datatypes "github.com/go-courier/sqlx/v2/datatypes"
)

func (Model) PrimaryKey() []string {
	return []string{
		"ID",
	}
}

func (Model) Indexes() github_com_go_courier_sqlx_v2_builder.Indexes {
	return github_com_go_courier_sqlx_v2_builder.Indexes{
		"i_field_string": []string{
			"FieldString",
		},
	}
}

func (Model) UniqueIndexUIModelID() string {
	return "ui_model_id"
}

func (Model) UniqueIndexUIUnionIdx() string {
	return "ui_union_idx"
}

func (Model) UniqueIndexes() github_com_go_courier_sqlx_v2_builder.Indexes {
	return github_com_go_courier_sqlx_v2_builder.Indexes{
		"ui_model_id": []string{
			"ModelID",
		},
		"ui_union_idx": []string{
			"UnionIndexField1",
			"UnionIndexField2",
		},
	}
}

func (Model) Comments() map[string]string {
	return map[string]string{}
}

var ModelTable *github_com_go_courier_sqlx_v2_builder.Table

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

func (Model) TableDescription() []string {
	return []string{
		"Model database model demo",
	}
}

func (Model) ColDescriptions() map[string][]string {
	return map[string][]string{}
}

func (Model) FieldKeyID() string {
	return "ID"
}

func (m *Model) FieldID() *github_com_go_courier_sqlx_v2_builder.Column {
	return ModelTable.F(m.FieldKeyID())
}

func (Model) FieldKeyModelID() string {
	return "ModelID"
}

func (m *Model) FieldModelID() *github_com_go_courier_sqlx_v2_builder.Column {
	return ModelTable.F(m.FieldKeyModelID())
}

func (Model) FieldKeyFieldString() string {
	return "FieldString"
}

func (m *Model) FieldFieldString() *github_com_go_courier_sqlx_v2_builder.Column {
	return ModelTable.F(m.FieldKeyFieldString())
}

func (Model) FieldKeyFieldJsonContent() string {
	return "FieldJsonContent"
}

func (m *Model) FieldFieldJsonContent() *github_com_go_courier_sqlx_v2_builder.Column {
	return ModelTable.F(m.FieldKeyFieldJsonContent())
}

func (Model) FieldKeyUnionIndexField1() string {
	return "UnionIndexField1"
}

func (m *Model) FieldUnionIndexField1() *github_com_go_courier_sqlx_v2_builder.Column {
	return ModelTable.F(m.FieldKeyUnionIndexField1())
}

func (Model) FieldKeyUnionIndexField2() string {
	return "UnionIndexField2"
}

func (m *Model) FieldUnionIndexField2() *github_com_go_courier_sqlx_v2_builder.Column {
	return ModelTable.F(m.FieldKeyUnionIndexField2())
}

func (Model) FieldKeyCreatedAt() string {
	return "CreatedAt"
}

func (m *Model) FieldCreatedAt() *github_com_go_courier_sqlx_v2_builder.Column {
	return ModelTable.F(m.FieldKeyCreatedAt())
}

func (Model) FieldKeyUpdatedAt() string {
	return "UpdatedAt"
}

func (m *Model) FieldUpdatedAt() *github_com_go_courier_sqlx_v2_builder.Column {
	return ModelTable.F(m.FieldKeyUpdatedAt())
}

func (Model) ColRelations() map[string][]string {
	return map[string][]string{}
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

func (m *Model) ConditionByStruct(db github_com_go_courier_sqlx_v2.DBExecutor) github_com_go_courier_sqlx_v2_builder.SqlCondition {
	table := db.T(m)
	fieldValues := github_com_go_courier_sqlx_v2_builder.FieldValuesFromStructByNonZero(m)

	conditions := make([]github_com_go_courier_sqlx_v2_builder.SqlCondition, 0)

	for _, fieldName := range m.IndexFieldNames() {
		if v, exists := fieldValues[fieldName]; exists {
			conditions = append(conditions, table.F(fieldName).Eq(v))
			delete(fieldValues, fieldName)
		}
	}

	if len(conditions) == 0 {
		panic(fmt.Errorf("at least one of field for indexes has value"))
	}

	for fieldName, v := range fieldValues {
		conditions = append(conditions, table.F(fieldName).Eq(v))
	}

	condition := github_com_go_courier_sqlx_v2_builder.And(conditions...)

	return condition
}

func (m *Model) Create(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	if m.CreatedAt.IsZero() {
		m.CreatedAt = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	_, err := db.ExecExpr(github_com_go_courier_sqlx_v2.InsertToDB(db, m, nil))
	return err

}

func (m *Model) CreateOnDuplicateWithUpdateFields(db github_com_go_courier_sqlx_v2.DBExecutor, updateFields []string) error {

	if len(updateFields) == 0 {
		panic(fmt.Errorf("must have update fields"))
	}

	if m.CreatedAt.IsZero() {
		m.CreatedAt = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	fieldValues := github_com_go_courier_sqlx_v2_builder.FieldValuesFromStructByNonZero(m, updateFields...)

	delete(fieldValues, "ID")

	table := db.T(m)

	cols, vals := table.ColumnsAndValuesByFieldValues(fieldValues)

	fields := make(map[string]bool, len(updateFields))
	for _, field := range updateFields {
		fields[field] = true
	}

	for _, fieldNames := range m.UniqueIndexes() {
		for _, field := range fieldNames {
			delete(fields, field)
		}
	}

	if len(fields) == 0 {
		panic(fmt.Errorf("no fields for updates"))
	}

	for field := range fieldValues {
		if !fields[field] {
			delete(fieldValues, field)
		}
	}

	additions := github_com_go_courier_sqlx_v2_builder.Additions{}

	switch db.Dialect().DriverName() {
	case "mysql":
		additions = append(additions, github_com_go_courier_sqlx_v2_builder.OnDuplicateKeyUpdate(table.AssignmentsByFieldValues(fieldValues)...))
	case "postgres":
		indexes := m.UniqueIndexes()
		fields := make([]string, 0)
		for _, fs := range indexes {
			fields = append(fields, fs...)
		}
		indexFields, _ := db.T(m).Fields(fields...)

		additions = append(additions,
			github_com_go_courier_sqlx_v2_builder.OnConflict(indexFields).
				DoUpdateSet(table.AssignmentsByFieldValues(fieldValues)...))
	}

	additions = append(additions, github_com_go_courier_sqlx_v2_builder.Comment("User.CreateOnDuplicateWithUpdateFields"))

	expr := github_com_go_courier_sqlx_v2_builder.Insert().Into(table, additions...).Values(cols, vals...)

	_, err := db.ExecExpr(expr)
	return err

}

func (m *Model) DeleteByStruct(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	_, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Delete().
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(m.ConditionByStruct(db)),
				github_com_go_courier_sqlx_v2_builder.Comment("Model.DeleteByStruct"),
			),
	)

	return err
}

func (m *Model) FetchByID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("ID").Eq(m.ID),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("Model.FetchByID"),
			),
		m,
	)

	return err
}

func (m *Model) UpdateByIDWithMap(db github_com_go_courier_sqlx_v2.DBExecutor, fieldValues github_com_go_courier_sqlx_v2_builder.FieldValues) error {

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	table := db.T(m)

	result, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Update(db.T(m)).
			Where(
				github_com_go_courier_sqlx_v2_builder.And(
					table.F("ID").Eq(m.ID),
				),
				github_com_go_courier_sqlx_v2_builder.Comment("Model.UpdateByIDWithMap"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return m.FetchByID(db)
	}

	return nil

}

func (m *Model) UpdateByIDWithStruct(db github_com_go_courier_sqlx_v2.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_go_courier_sqlx_v2_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByIDWithMap(db, fieldValues)

}

func (m *Model) FetchByIDForUpdate(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("ID").Eq(m.ID),
				)),
				github_com_go_courier_sqlx_v2_builder.ForUpdate(),
				github_com_go_courier_sqlx_v2_builder.Comment("Model.FetchByIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *Model) DeleteByID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Delete().
			From(db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("ID").Eq(m.ID),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("Model.DeleteByID"),
			))

	return err
}

func (m *Model) FetchByModelID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("ModelID").Eq(m.ModelID),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("Model.FetchByModelID"),
			),
		m,
	)

	return err
}

func (m *Model) UpdateByModelIDWithMap(db github_com_go_courier_sqlx_v2.DBExecutor, fieldValues github_com_go_courier_sqlx_v2_builder.FieldValues) error {

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	table := db.T(m)

	result, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Update(db.T(m)).
			Where(
				github_com_go_courier_sqlx_v2_builder.And(
					table.F("ModelID").Eq(m.ModelID),
				),
				github_com_go_courier_sqlx_v2_builder.Comment("Model.UpdateByModelIDWithMap"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return m.FetchByModelID(db)
	}

	return nil

}

func (m *Model) UpdateByModelIDWithStruct(db github_com_go_courier_sqlx_v2.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_go_courier_sqlx_v2_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByModelIDWithMap(db, fieldValues)

}

func (m *Model) FetchByModelIDForUpdate(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("ModelID").Eq(m.ModelID),
				)),
				github_com_go_courier_sqlx_v2_builder.ForUpdate(),
				github_com_go_courier_sqlx_v2_builder.Comment("Model.FetchByModelIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *Model) DeleteByModelID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Delete().
			From(db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("ModelID").Eq(m.ModelID),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("Model.DeleteByModelID"),
			))

	return err
}

func (m *Model) FetchByUnionIndexField1AndUnionIndexField2(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("UnionIndexField1").Eq(m.UnionIndexField1),
					table.F("UnionIndexField2").Eq(m.UnionIndexField2),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("Model.FetchByUnionIndexField1AndUnionIndexField2"),
			),
		m,
	)

	return err
}

func (m *Model) UpdateByUnionIndexField1AndUnionIndexField2WithMap(db github_com_go_courier_sqlx_v2.DBExecutor, fieldValues github_com_go_courier_sqlx_v2_builder.FieldValues) error {

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	table := db.T(m)

	result, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Update(db.T(m)).
			Where(
				github_com_go_courier_sqlx_v2_builder.And(
					table.F("UnionIndexField1").Eq(m.UnionIndexField1),
					table.F("UnionIndexField2").Eq(m.UnionIndexField2),
				),
				github_com_go_courier_sqlx_v2_builder.Comment("Model.UpdateByUnionIndexField1AndUnionIndexField2WithMap"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return m.FetchByUnionIndexField1AndUnionIndexField2(db)
	}

	return nil

}

func (m *Model) UpdateByUnionIndexField1AndUnionIndexField2WithStruct(db github_com_go_courier_sqlx_v2.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_go_courier_sqlx_v2_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByUnionIndexField1AndUnionIndexField2WithMap(db, fieldValues)

}

func (m *Model) FetchByUnionIndexField1AndUnionIndexField2ForUpdate(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("UnionIndexField1").Eq(m.UnionIndexField1),
					table.F("UnionIndexField2").Eq(m.UnionIndexField2),
				)),
				github_com_go_courier_sqlx_v2_builder.ForUpdate(),
				github_com_go_courier_sqlx_v2_builder.Comment("Model.FetchByUnionIndexField1AndUnionIndexField2ForUpdate"),
			),
		m,
	)

	return err
}

func (m *Model) DeleteByUnionIndexField1AndUnionIndexField2(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Delete().
			From(db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("UnionIndexField1").Eq(m.UnionIndexField1),
					table.F("UnionIndexField2").Eq(m.UnionIndexField2),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("Model.DeleteByUnionIndexField1AndUnionIndexField2"),
			))

	return err
}

func (m *Model) List(db github_com_go_courier_sqlx_v2.DBExecutor, condition github_com_go_courier_sqlx_v2_builder.SqlCondition, additions ...github_com_go_courier_sqlx_v2_builder.Addition) ([]Model, error) {

	list := make([]Model, 0)

	table := db.T(m)
	_ = table

	finalAdditions := []github_com_go_courier_sqlx_v2_builder.Addition{
		github_com_go_courier_sqlx_v2_builder.Where(condition),
		github_com_go_courier_sqlx_v2_builder.Comment("Model.List"),
	}

	if len(additions) > 0 {
		finalAdditions = append(finalAdditions, additions...)
	}

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(db.T(m), finalAdditions...),
		&list,
	)

	return list, err

}

func (m *Model) Count(db github_com_go_courier_sqlx_v2.DBExecutor, condition github_com_go_courier_sqlx_v2_builder.SqlCondition, additions ...github_com_go_courier_sqlx_v2_builder.Addition) (int, error) {

	count := -1

	table := db.T(m)
	_ = table

	finalAdditions := []github_com_go_courier_sqlx_v2_builder.Addition{
		github_com_go_courier_sqlx_v2_builder.Where(condition),
		github_com_go_courier_sqlx_v2_builder.Comment("Model.Count"),
	}

	if len(additions) > 0 {
		finalAdditions = append(finalAdditions, additions...)
	}

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(
			github_com_go_courier_sqlx_v2_builder.Count(),
		).
			From(db.T(m), finalAdditions...),
		&count,
	)

	return count, err

}

func (m *Model) BatchFetchByFieldStringList(db github_com_go_courier_sqlx_v2.DBExecutor, values []string) ([]Model, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("FieldString").In(values)

	return m.List(db, condition)

}

func (m *Model) BatchFetchByIDList(db github_com_go_courier_sqlx_v2.DBExecutor, values []uint64) ([]Model, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("ID").In(values)

	return m.List(db, condition)

}

func (m *Model) BatchFetchByModelIDList(db github_com_go_courier_sqlx_v2.DBExecutor, values []uint64) ([]Model, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("ModelID").In(values)

	return m.List(db, condition)

}

func (m *Model) BatchFetchByUnionIndexField1List(db github_com_go_courier_sqlx_v2.DBExecutor, values []int64) ([]Model, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("UnionIndexField1").In(values)

	return m.List(db, condition)

}

func (m *Model) BatchFetchByUnionIndexField2List(db github_com_go_courier_sqlx_v2.DBExecutor, values []string) ([]Model, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("UnionIndexField2").In(values)

	return m.List(db, condition)

}

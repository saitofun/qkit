package models

//go:generate toolkit gen model Other --database Demo
// Other database model demo
// @def primary                     ID
// @def index        I_field_string FieldString
// @def unique_index UI_model_id    ModelID
// @def unique_index UI_union_idx   UnionIndexField1 UnionIndexField2
type Other struct {
	PrimaryID
	RefModel
	BaseModel
	OperationTimes
}

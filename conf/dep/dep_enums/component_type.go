package dep_enums

//go:generate tools gen enum ComponentType
// ComponentType 组件类型
type ComponentType uint8

const (
	COMPONENT_TYPE_UNKNOWN   ComponentType = iota
	COMPONENT_TYPE__REQUIRED               // 必选
	COMPONENT_TYPE__OPTIONAL               // 可选
)

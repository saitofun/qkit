package dep_enums

//go:generate tools gen enum ComponentName
type ComponentName uint8

const (
	COMPONENT_NAME_UNKNOWN ComponentName = iota
	COMPONENT_NAME__REDIS
)

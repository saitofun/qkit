package dep_enums

//go:generate tools gen enum RuntimeMethod
type RuntimeMethod uint8

const (
	RUNTIME_METHOD_UNKNOWN RuntimeMethod = iota
	RUNTIME_METHOD__SUPERVISOR
	RUNTIME_METHOD__DOCKER
)

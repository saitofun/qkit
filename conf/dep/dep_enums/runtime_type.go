package dep_enums

//go:generate tools gen enum RuntimeType
type RuntimeType uint8

const (
	RUNTIME_TYPE_UNKNOWN RuntimeType = iota
	RUNTIME_TYPE__STATIC             // 静态可执行资源
	RUNTIME_TYPE__DAEMON             // 后台服务
)

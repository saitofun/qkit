package example

type Schema int

const (
	SCHEMA_UNKNOWN Schema = iota
	SCHEMA__HTTP          // http
	SCHEMA__HTTPS         // https
	_
	_1
	_2
	SCHEMA__TCP
	SCHEMA__UDP
	SCHEMA__QUIC
)

func (Schema) Offset() int { return 1 }

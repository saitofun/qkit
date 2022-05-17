package types

import (
	"encoding"
	"reflect"
	"time"
)

type (
	SecurityString interface{ SecurityString() string }
	String         interface{ String() string }

	DefaultSetter   interface{ SetDefault() }
	Initializer     interface{ Init() }
	InitializerWith interface{ Init(interface{}) }

	Span interface {
		// Duration returns common duration value as `time.Duration`
		Duration() time.Duration
		// Int returns an integer value, such as int64(Span.Duration())
		Int() int64
		// String returns a duration formatted string, commonly call Span.Duration().String()
		// eg: Second(1).String() returns `"1s"`
		String() string
		// Literal returns a string presents an integer value regardless of time unit
		// eg: Second(1).Literal() returns `"1"`
		Literal() string
	}

	TextMarshaler   = encoding.TextMarshaler
	TextUnmarshaler = encoding.TextUnmarshaler
)

var (
	RTypeString          = reflect.TypeOf((*String)(nil)).Elem()
	RTypeSecurityString  = reflect.TypeOf((*SecurityString)(nil)).Elem()
	RTypeTextMarshaler   = reflect.TypeOf((*TextMarshaler)(nil)).Elem()
	RTypeTextUnmarshaler = reflect.TypeOf((*TextUnmarshaler)(nil)).Elem()
	RTypeDefaultSetter   = reflect.TypeOf((*DefaultSetter)(nil)).Elem()
	RTypeInitializer     = reflect.TypeOf((*Initializer)(nil)).Elem()
	RTypeInitializerWith = reflect.TypeOf((*InitializerWith)(nil)).Elem()
)

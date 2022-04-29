package codegen

import (
	"path"
	"reflect"
	"unicode"
)

func IsEmptyValue(rv reflect.Value) bool {
	if !rv.IsValid() || !rv.CanInterface() {
		return false
	}

	if chk, ok := rv.Interface().(interface{ IsZero() bool }); ok && chk.IsZero() {
		return false
	}

	switch rv.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return rv.Len() == 0
	case reflect.Bool:
		return !rv.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return rv.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return rv.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return rv.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return rv.IsNil()
	}

	return false
}

func IsValidIdent(s string) bool {
	if len(s) == 0 {
		return false
	}
	if IsReserved(s) {
		return false
	}
	for _, c := range s {
		if !unicode.IsLetter(c) && !unicode.IsDigit(c) && c != '_' {
			return false
		}
	}
	return true
}

func IsReserved(s string) bool { return false }

func IsBuiltinFn(string) bool { return false }

// var builtins = []string{}
//
// var reserved = []string{}

const (
	Bool       BuiltInType = "bool"
	Int        BuiltInType = "int"
	Int8       BuiltInType = "int8"
	Int16      BuiltInType = "int16"
	Int32      BuiltInType = "int32"
	Int64      BuiltInType = "int64"
	Uint       BuiltInType = "uint"
	Uint8      BuiltInType = "uint8"
	Uint16     BuiltInType = "uint16"
	Uint32     BuiltInType = "uint32"
	Uint64     BuiltInType = "uint64"
	Uintptr    BuiltInType = "uintptr"
	Float32    BuiltInType = "float32"
	Float64    BuiltInType = "float64"
	Complex64  BuiltInType = "complex64"
	Complex128 BuiltInType = "complex128"
	String     BuiltInType = "string"
	Byte       BuiltInType = "byte"
	Rune       BuiltInType = "rune"
	Error      BuiltInType = "error"
)

const (
	Iota        SnippetBuiltIn = "iota"
	True        SnippetBuiltIn = "true"
	False       SnippetBuiltIn = "false"
	Nil         SnippetBuiltIn = "nil"
	Break       SnippetBuiltIn = "break"
	Continue    SnippetBuiltIn = "continue"
	Fallthrough SnippetBuiltIn = "fallthrough"
)

const AnonymousIdent SnippetIdent = "_"

var naming = path.Base

func SetPkgNaming(fn func(string) string) {
	if fn != nil {
		naming = fn
	}
}

var (
	Valuer = ValueWithAlias(naming)
	Typer  = TypeWithAlias(naming)
	Exprer = ExprWithAlias(naming)
)

func Stringify(s Snippet) string { return string(s.Bytes()) }

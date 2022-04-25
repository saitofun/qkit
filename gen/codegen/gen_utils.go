package codegen

import (
	"reflect"
	"unicode"
)

func IsEmptyValue(rv reflect.Value) bool { return false }

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

const Anonymous SnippetIdent = "_"

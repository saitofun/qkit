package errors

import (
	"bytes"
	"fmt"
)

func NewErrorSet(root string) *ErrorSet {
	return &ErrorSet{root: root}
}

type ErrorSet struct {
	root   string
	errors []*FieldError
}

func (set *ErrorSet) AddErr(err error, keyPathNodes ...interface{}) {
	if err == nil {
		return
	}
	set.errors = append(set.errors, &FieldError{KeyPath(keyPathNodes), err})
}

func (set *ErrorSet) Each(cb func(*FieldError)) {
	for _, fe := range set.errors {
		cb(fe)
	}
}

func (set *ErrorSet) Flatten() *ErrorSet {
	ret := NewErrorSet(set.root)

	set.Each(
		func(f *FieldError) {
			if sub, ok := f.Error.(*ErrorSet); ok {
				sub.Flatten().Each(
					func(ff *FieldError) {
						ret.AddErr(ff.Error, append(f.Field, ff.Field...)...)
					},
				)
			} else {
				ret.AddErr(f.Error, f.Field...)
			}
		},
	)

	return ret
}
func (set *ErrorSet) Len() int { return len(set.Flatten().errors) }

func (set *ErrorSet) Err() error {
	if len(set.errors) == 0 {
		return nil
	}
	return set
}

func (set *ErrorSet) Error() string {
	errs := set.Flatten()

	buf := bytes.Buffer{}
	errs.Each(
		func(f *FieldError) {
			buf.WriteString(fmt.Sprintf("%s %s", f.Field, f.Error))
			buf.WriteRune('\n')
		},
	)

	return buf.String()
}

type FieldError struct {
	Field KeyPath
	Error error `json:"msg"`
}

type KeyPath []interface{}

func (keyPath KeyPath) String() string {
	buf := &bytes.Buffer{}
	for i := 0; i < len(keyPath); i++ {
		switch keyOrIndex := keyPath[i].(type) {
		case string:
			if buf.Len() > 0 {
				buf.WriteRune('.')
			}
			buf.WriteString(keyOrIndex)
		case int:
			buf.WriteString(fmt.Sprintf("[%d]", keyOrIndex))
		}
	}
	return buf.String()
}

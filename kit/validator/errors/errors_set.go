package errors

import (
	"bytes"
	"container/list"
	"fmt"
)

func NewErrorSet(root string) *ErrorSet {
	return &ErrorSet{
		root:   root,
		errors: list.New(),
	}
}

type ErrorSet struct {
	root   string
	errors *list.List
}

func (errs *ErrorSet) AddErr(err error, keyPathNodes ...interface{}) {
	if err == nil {
		return
	}
	errs.errors.PushBack(&FieldError{
		Field: keyPathNodes,
		Error: err,
	})
}

func (errs *ErrorSet) Each(cb func(fieldErr *FieldError)) {
	l := errs.errors
	for e := l.Front(); e != nil; e = e.Next() {
		if fieldErr, ok := e.Value.(*FieldError); ok {
			cb(fieldErr)
		}
	}
}

func (errs *ErrorSet) Flatten() *ErrorSet {
	set := NewErrorSet(errs.root)

	errs.Each(func(fieldErr *FieldError) {
		if subSet, ok := fieldErr.Error.(*ErrorSet); ok {
			subSet.Flatten().Each(func(subSetFieldErr *FieldError) {
				set.AddErr(subSetFieldErr.Error, append(fieldErr.Field, subSetFieldErr.Field...)...)
			})
		} else {
			set.AddErr(fieldErr.Error, fieldErr.Field...)
		}
	})

	return set
}

func (errs *ErrorSet) Len() int {
	return errs.Flatten().errors.Len()
}

func (errs *ErrorSet) Err() error {
	if errs.errors.Len() == 0 {
		return nil
	}
	return errs
}

func (errs *ErrorSet) Error() string {
	set := errs.Flatten()

	buf := bytes.Buffer{}
	set.Each(func(fieldErr *FieldError) {
		buf.WriteString(fmt.Sprintf("%s %s", fieldErr.Field, fieldErr.Error))
		buf.WriteRune('\n')
	})

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

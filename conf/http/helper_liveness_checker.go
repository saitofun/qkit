package http

import (
	"context"
	"reflect"

	"github.com/saitofun/qkit/kit/httptransport/httpx"
	"github.com/saitofun/qkit/kit/kit"
	"github.com/saitofun/qkit/x/reflectx"
)

type LivenessChecker interface {
	LivenessCheck() map[string]string
}

func RegisterCheckerBy(vs ...interface{}) {
	f :
	for _, v := range vs {
		rv := reflectx.Indirect(reflect.ValueOf(v))
		typ := rv.Type()

		if typ.Kind() != reflect.Struct {
			continue
		}

		if with, ok := rv.Interface().(LivenessChecker); ok {
			RegisterChecker()
		}

		for i := 0; i < rv.NumField(); i++ {
			value := rv.Field(i)
			name := typ.Field(i).Name

			if livenessChecker, ok := value.Interface().(LivenessChecker); ok {
				RegisterChecker(name, livenessChecker)
			}
		}
	}
}

func RegisterChecker(k string, checker LivenessChecker) { checkers[k] = checker }

type LivenessCheckers map[string]LivenessChecker

var checkers = LivenessCheckers{}

func (cs LivenessCheckers) Statuses() map[string]string {
	m := map[string]string{}

	for name := range cs {
		if cs[name] != nil {
			entry := cs[name].LivenessCheck()
			for key, v := range entry {
				if key != "" {
					m[name+"/"+key] = v
				} else {
					m[name] = v
				}
			}
		}
	}

	return m
}

var LivenessRouter = kit.NewRouter(&Liveness{})

type Liveness struct{ httpx.MethodGet }

func (Liveness) Path() string { return "/liveness" }

func (Liveness) Output(ctx context.Context) (interface{}, error) {
	return checkers.Statuses(), nil
}

// type LivenessStatuses map[string]string

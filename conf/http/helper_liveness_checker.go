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
	for _, v := range vs {
		rv := reflectx.Indirect(reflect.ValueOf(v))
		rt := rv.Type()

		if rt.Kind() != reflect.Struct {
			continue
		}

		if chk, ok := rv.Interface().(LivenessChecker); ok {
			RegisterChecker(rt.Name(), chk)
			continue
		}

		for i := 0; i < rv.NumField(); i++ {
			fv := rv.Field(i)
			name := rt.Field(i).Name

			if chk, ok := fv.Interface().(LivenessChecker); ok {
				RegisterChecker(name, chk)
			}
		}
	}
}

func RegisterChecker(k string, chk LivenessChecker) { checkers[k] = chk }

func ResetRegistered() { checkers = LivenessCheckers{} }

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

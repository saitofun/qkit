package builder

import (
	"context"
	"strconv"
)

type limit struct {
	count  int64
	offset int64
	AdditionType
}

func Limit(count int64) *limit { return &limit{AdditionType: AdditionLimit, count: count} }

func (l limit) Offset(offset int64) *limit { l.offset = offset; return &l }

func (l *limit) IsNil() bool { return l == nil || l.count <= 0 }

func (l *limit) Ex(ctx context.Context) *Ex {
	e := ExactlyExpr("LIMIT ")
	e.WriteQuery(strconv.FormatInt(l.count, 10))
	if l.offset > 0 {
		e.WriteQuery(" OFFSET ")
		e.WriteQuery(strconv.FormatInt(l.offset, 10))
	}
	return e.Ex(ctx)
}

type Pager struct {
	Size   int64 `name:"size,omitempty"   in:"query" default:"10" validate:"@int64[-1,]"`
	Offset int64 `name:"offset,omitempty" in:"query" default:"10" validate:"@int64[0,]"`
}

func (p *Pager) Addition() Addition {
	if p.Size != -1 {
		return Limit(p.Size).Offset(p.Offset)
	}
	return nil
}

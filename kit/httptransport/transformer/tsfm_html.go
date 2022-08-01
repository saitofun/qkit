package transformer

import (
	"context"
	"io"
	"net/textproto"
	"reflect"

	"github.com/saitofun/qkit/kit/httptransport/httpx"
	"github.com/saitofun/qkit/x/typesx"
	"github.com/saitofun/qlib/encoding/qtext"
)

const HTMLTsfName = "text/html"

func init() { DefaultFactory.Register(&HTMLText{}) }

type HTMLText struct{}

func (t *HTMLText) String() string { return HTMLTsfName }

func (HTMLText) Names() []string { return []string{HTMLTsfName, "html"} }

func (HTMLText) NamedByTag() string { return "" }

func (HTMLText) New(context.Context, typesx.Type) (Transformer, error) { return &HTMLText{}, nil }

func (t *HTMLText) EncodeTo(ctx context.Context, w io.Writer, v interface{}) error {
	rv, ok := v.(reflect.Value)
	if !ok {
		rv = reflect.ValueOf(v)
	}

	httpx.MaybeWriteHeader(ctx, w, t.String(), map[string]string{
		"charset": "utf-8",
	})

	data, err := qtext.MarshalText(rv, true)
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	return err
}

func (HTMLText) DecodeFrom(_ context.Context, r io.Reader, v interface{}, _ ...textproto.MIMEHeader) error {
	rv, ok := v.(reflect.Value)
	if !ok {
		rv = reflect.ValueOf(v)
	}
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	return qtext.UnmarshalText(rv, data, true)
}

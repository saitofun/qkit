package transformer_test

import (
	"context"
	"testing"
	_ "unsafe"

	"github.com/saitofun/qkit/kit/httptransport/transformer"
)

var (
	bgctx = context.Background()
)

func TestTransformer(t *testing.T) {
	tfs := transformer.Transformers()
	for _, name := range tfs {
		t.Log(name)
	}
}

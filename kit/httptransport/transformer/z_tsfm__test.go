package transformer_test

import (
	"context"
	"testing"
	_ "unsafe"
)

var (
	bgctx = context.Background()
)

//go:linkname transformers github.com/saitofun/qkit/kit/httptransport/transformer.transformers
func transformers() []string

func TestTransformer(t *testing.T) {
	tfs := transformers()
	for _, name := range tfs {
		t.Log(name)
	}
}

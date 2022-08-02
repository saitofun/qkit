package httptransport_test

import (
	"context"
	"os"
	"testing"

	. "github.com/onsi/gomega"

	. "github.com/saitofun/qkit/kit/httptransport"
)

func TestServiceMeta(t *testing.T) {
	meta := &ServiceMeta{}

	name, version := "srv-test", "1.1.1"

	os.Setenv(EnvProjectName, name)
	meta.SetDefault()
	NewWithT(t).Expect(meta.String()).To(Equal(name))

	os.Setenv(EnvProjectVersion, version)
	meta.SetDefault()
	NewWithT(t).Expect(meta.String()).To(Equal(name + "@" + version))
}

func TestServiceMetaWithContext(t *testing.T) {
	meta := ServiceMeta{Name: "test"}
	ctx := ContextWithServiceMeta(context.Background(), meta)
	got := ServiceMetaFromContext(ctx)
	NewWithT(t).Expect(got).To(Equal(meta))
}

var (
	bgctx = context.Background()
)

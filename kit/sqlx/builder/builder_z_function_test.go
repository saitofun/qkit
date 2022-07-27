package builder_test

import (
	"testing"

	"github.com/onsi/gomega"
	"github.com/saitofun/qkit/testutil/buildertestutil"

	"github.com/saitofun/qkit/kit/sqlx/builder"
)

func TestFunc(t *testing.T) {
	t.Run("Invalid", func(t *testing.T) {
		gomega.NewWithT(t).Expect(builder.Func("")).To(buildertestutil.BeExpr(""))
	})
	t.Run("Count", func(t *testing.T) {
		gomega.NewWithT(t).Expect(builder.Count()).To(buildertestutil.BeExpr("COUNT(1)"))
	})
	t.Run("Avg", func(t *testing.T) {
		gomega.NewWithT(t).Expect(builder.Avg()).To(buildertestutil.BeExpr("AVG(*)"))
	})
}

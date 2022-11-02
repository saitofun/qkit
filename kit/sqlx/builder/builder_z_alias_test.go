package builder_test

import (
	"testing"

	g "github.com/onsi/gomega"

	. "github.com/saitofun/qkit/kit/sqlx/builder"
	. "github.com/saitofun/qkit/testutil/buildertestutil"
)

func TestAlias(t *testing.T) {
	t.Run("Alias", func(t *testing.T) {
		g.NewWithT(t).Expect(Alias(Expr("f_id"), "id")).
			To(BeExpr("f_id AS id"))
	})
}

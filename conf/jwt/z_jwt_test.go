package jwt_test

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	. "github.com/onsi/gomega"

	"github.com/saitofun/qkit/base/types"
	. "github.com/saitofun/qkit/conf/jwt"
)

func TestJwt(t *testing.T) {
	c := &Jwt{
		Issuer: "jwt_test",
		ExpIn:  *types.AsDuration(types.Seconds(1).Duration()),
		// Method:  jwt.SIGNING_METHOD__ECDSA256,
		SignKey: "xxx",
	}

	t.Run("GenerateAndParse", func(t *testing.T) {
		token, err := c.GenerateTokenByPayload("100")
		NewWithT(t).Expect(err).To(BeNil())
		NewWithT(t).Expect(token).NotTo(BeEmpty())

		claim, err := c.ParseToken(token)
		NewWithT(t).Expect(err).To(BeNil())
		v, ok := claim.Payload.(string)
		NewWithT(t).Expect(ok).To(BeTrue())
		NewWithT(t).Expect(v).To(Equal("100"))
	})

	t.Run("TokenExpired", func(t *testing.T) {
		token, err := c.GenerateTokenByPayload("100")
		NewWithT(t).Expect(err).To(BeNil())
		NewWithT(t).Expect(token).NotTo(BeEmpty())

		time.Sleep(2 * time.Second)

		_, err = c.ParseToken(token)
		NewWithT(t).Expect(err).NotTo(BeNil())
		ve, ok := err.(*jwt.ValidationError)
		NewWithT(t).Expect(ok).To(BeTrue())
		NewWithT(t).Expect(ve.Errors | jwt.ValidationErrorExpired).To(Equal(jwt.ValidationErrorExpired))
	})
}

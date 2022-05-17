package reflectx_test

import (
	"reflect"
	"testing"

	. "github.com/onsi/gomega"
	. "github.com/sincospro/qkit/x/reflectx"
)

func TestNatureType(t *testing.T) {
	type Foo struct{}

	var v, vp = Foo{}, &Foo{}

	NewWithT(t).Expect(NatureType(v).String()).To(Equal("reflectx_test.Foo"))
	NewWithT(t).Expect(NatureType(vp).String()).To(Equal("reflectx_test.Foo"))
	NewWithT(t).Expect(NatureType(&vp).String()).To(Equal("reflectx_test.Foo"))
	NewWithT(t).Expect(NatureType(reflect.TypeOf(v)).String()).To(Equal("reflectx_test.Foo"))
	NewWithT(t).Expect(NatureType(reflect.TypeOf(vp)).String()).To(Equal("reflectx_test.Foo"))
	NewWithT(t).Expect(NatureType(reflect.TypeOf(&vp)).String()).To(Equal("reflectx_test.Foo"))
}

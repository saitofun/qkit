package dep

import "github.com/sincospro/qkit/base/types"

type Assert struct {
	Addr    types.Address
	Name    string
	Version string
	Type    AssertType
}

type AssertType uint8

const (
	AssertTypeUnknown = iota
	AssertTypeRPM
	AssertTypeDep
	AssertTypeDocker
)

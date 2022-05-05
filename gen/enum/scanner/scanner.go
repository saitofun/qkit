package scanner

import (
	"go/types"
	"sort"

	"golang.org/x/tools/go/packages"
)

type Scanner struct {
	pkg *packages.Package
	res map[*types.TypeName]Options
}

func NewScanner(pkg *packages.Package) *Scanner { return &Scanner{pkg: pkg} }

func (s *Scanner) Options(tn *types.TypeName) (Options, bool) {
	return nil, false
}

func (s *Scanner) Append(tn *types.TypeName, opt *Option) {
	if s.res == nil {
		s.res = make(map[*types.TypeName]Options)
	}
	s.res[tn] = append(s.res[tn], *opt)
	sort.Sort(s.res[tn])
}

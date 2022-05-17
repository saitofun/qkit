package pkg

import (
	. "golang.org/x/tools/go/packages"
)

type Set map[string]*Package

func (s Set) Append(pkg *Package) {
	s[pkg.ID] = pkg
	for i := range pkg.Imports {
		if _, ok := s[i]; !ok {
			s.Append(pkg.Imports[i])
		}
	}
}

func (s Set) List() (ret []*Package) {
	for id := range s {
		ret = append(ret, s[id])
	}
	return ret
}

type Imports map[string]*Package

func (s Imports) Append(pkg *Package) {
	s[pkg.PkgPath] = pkg
	for i := range pkg.Imports {
		if _, ok := s[pkg.Imports[i].PkgPath]; !ok {
			s.Append(pkg)
		}
	}
}

func (s Imports) List() (ret []*Package) {
	for pth := range s {
		ret = append(ret, s[pth])
	}
	return ret
}

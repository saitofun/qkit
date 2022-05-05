package pkg

import "golang.org/x/tools/go/packages"

type pkgs map[string]*packages.Package

func (s pkgs) add(pkg *packages.Package) {
	s[pkg.ID] = pkg
	for i := range pkg.Imports {
		if _, ok := s[i]; !ok {
			s.add(pkg.Imports[i])
		}
	}
}

func (s pkgs) packages() (ret []*packages.Package) {
	for id := range s {
		ret = append(ret, s[id])
	}
	return ret
}

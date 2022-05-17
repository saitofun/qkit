package pkg

import (
	"go/ast"
	"go/token"
	"go/types"

	. "golang.org/x/tools/go/packages"
)

type Pkg struct {
	*Package
	lst []*Package
	set *Set
}

type Pos interface{ Pos() token.Pos }

type End interface{ End() token.Pos }

func LoadFrom(pattern string) (*Pkg, error) {
	lst, err := Load(
		&Config{Mode: LoadMode(0b111111111111)},
		pattern,
	)
	if err != nil {
		return nil, err
	}
	return New(lst[0]), nil
}

func New(pkg *Package) *Pkg {
	set := &Set{}
	set.Append(pkg)
	return &Pkg{
		Package: pkg,
		lst:     set.List(),
		set:     set,
	}
}

func (p *Pkg) Imports() []*Package { return p.lst }

func (p *Pkg) Const(name string) *types.Const {
	for ident, def := range p.TypesInfo.Defs {
		if t, ok := def.(*types.Const); ok && ident.Name == name {
			return t
		}
	}
	return nil
}

func (p *Pkg) TypeName(name string) *types.TypeName {
	for ident, def := range p.TypesInfo.Defs {
		if t, ok := def.(*types.TypeName); ok && ident.Name == name {
			return t
		}
	}
	return nil
}

func (p *Pkg) Var(name string) *types.Var {
	for ident, def := range p.TypesInfo.Defs {
		if t, ok := def.(*types.Var); ok && ident.Name == name {
			return t
		}
	}
	return nil
}

func (p *Pkg) Func(name string) *types.Func {
	for ident, def := range p.TypesInfo.Defs {
		if t, ok := def.(*types.Func); ok && ident.Name == name {
			return t
		}
	}
	return nil
}

func (p *Pkg) Pkg(path string) *Package {
	for _, pkg := range p.lst {
		if path == pkg.PkgPath {
			return pkg
		}
	}
	return nil
}

func (p *Pkg) PkgOf(n Pos) *types.Package {
	for _, pkg := range p.lst {
		for _, file := range pkg.Syntax {
			if file.Pos() <= n.Pos() && file.End() > n.Pos() {
				return pkg.Types
			}
		}
	}
	return nil
}

func (p *Pkg) PkgInfoOf(n Pos) *types.Info {
	for _, pkg := range p.lst {
		for _, file := range pkg.Syntax {
			if file.Pos() <= n.Pos() && file.End() > n.Pos() {
				return pkg.TypesInfo
			}
		}
	}
	return nil
}

func (p *Pkg) FileOf(n Pos) *ast.File {
	for _, pkg := range p.lst {
		for _, file := range pkg.Syntax {
			if file.Pos() <= n.Pos() && file.End() > n.Pos() {
				return file
			}
		}
	}
	return nil
}

func (p *Pkg) IdentOf(obj types.Object) *ast.Ident {
	info := p.Pkg(obj.Pkg().Path())

	for ident, def := range info.TypesInfo.Defs {
		if def == obj {
			return ident
		}
	}
	return nil
}

func (p *Pkg) CommentsOf(n ast.Node) string {
	if f := p.FileOf(n); f == nil {
		return ""
	} else {
		return NewCommentScanner(p.Fset, f).CommentsOf(n)
	}
}

func (p *Pkg) Eval(expr ast.Expr) (types.TypeAndValue, error) {
	return types.Eval(
		p.Fset, p.PkgOf(expr), expr.Pos(), StringifyNode(p.Fset, expr),
	)
}

func (p *Pkg) FuncDeclOf(fn *types.Func) (decl *ast.FuncDecl) {
	ast.Inspect(p.FileOf(fn), func(node ast.Node) bool {
		fd, ok := node.(*ast.FuncDecl)
		if !ok {
			return true
		}
		if fd.Pos() <= fn.Pos() && fd.Body != nil && fn.Pos() < fd.Body.Pos() {
			decl = fd
			return false
		}
		return true
	})
	return
}

func (p *Pkg) ResultsOf(callee *ast.CallExpr) (Results, int) {
	typ := p.PkgInfoOf(callee).TypeOf(callee)
	res := Results{}

	switch t := typ.(type) {
	case *types.Tuple:
		for i := 0; i < t.Len(); i++ {
			p.AppendResult(res, i, TypeAndValueExpr{
				TypeAndValue: types.TypeAndValue{Type: t.At(i).Type()},
				Expr:         callee,
			})
		}
	default:
		p.AppendResult(res, 0, TypeAndValueExpr{
			TypeAndValue: types.TypeAndValue{Type: t},
			Expr:         callee,
		})
	}

	return res, len(res)
}

func (p *Pkg) AssignedValueOf(ident *ast.Ident, pos token.Pos) []TypeAndValueExpr {
	var (
		ass  *ast.AssignStmt
		blk  *ast.BlockStmt
		idx  = 0
		file = p.FileOf(ident)
	)

	ast.Inspect(file, func(node ast.Node) bool {
		switch fn := node.(type) {
		case *ast.FuncLit:
			if fn.Pos() <= ident.Pos() && ident.Pos() <= fn.End() {
				blk = fn.Body
			}
			return false
		case *ast.FuncDecl:
			if fn.Pos() <= ident.Pos() && ident.Pos() <= fn.End() {
				blk = fn.Body
			}
			return false
		}
		return true
	})

	if blk == nil {
		return nil
	}

	func(n ast.Node) {
		nodes := []ast.Node{n}
		for len(nodes) > 0 {
			n, nodes = nodes[0], nodes[1:]
			ast.Inspect(n, func(node ast.Node) bool {
				if node == nil || node.Pos() > pos {
					return false
				}
				switch stmt := node.(type) {
				case *ast.CaseClause:
					return IsContainsReturn(stmt) ||
						stmt.Pos() <= pos && pos < stmt.End()
				case *ast.IfStmt:
					if stmt.Else != nil {
						nodes = append(nodes, stmt.Else)
					}
					return IsContainsReturn(stmt) ||
						stmt.Body.Pos() <= pos && pos < stmt.Body.End()
				case *ast.AssignStmt:
					for i := range stmt.Lhs {
						id, ok := stmt.Lhs[i].(*ast.Ident)
						if ok && ident.Obj == id.Obj {
							ass, idx = stmt, i
						}
					}
				}
				return true
			})
		}
	}(blk)

	if ass == nil {
		return nil
	}
	res := Results{}
	p.SetResultsBy(res, ass.Rhs...)
	return res[idx]
}

func (p *Pkg) AppendResult(res Results, i int, tve TypeAndValueExpr) {
	if _, ok := tve.Type.(*types.Interface); !ok {
		res[i] = append(res[i], tve)
		return
	}
	switch expr := tve.Expr.(type) {
	case *ast.Ident:
		res[i] = append(res[i], p.AssignedValueOf(expr, expr.Pos())...)
	case *ast.SelectorExpr:
		res[i] = append(res[i], p.AssignedValueOf(expr.Sel, expr.Sel.Pos())...)
	default:
		res[i] = append(res[i], tve)
	}
}

func (p *Pkg) SetResultsBy(res Results, exprs ...ast.Expr) {
	for i := range exprs {
		switch e := exprs[i].(type) {
		case *ast.CallExpr:
			_res, _len := p.ResultsOf(e)
			for j := 0; j < _len; j++ {
				if j > 0 {
					i++
				}
				for _, tve := range _res[j] {
					res[i] = append(res[i], TypeAndValueExpr{
						TypeAndValue: tve.TypeAndValue,
						Expr:         tve.Expr,
					})
				}
			}
		default:
			tv, _ := p.Eval(e)
			p.AppendResult(res, i, TypeAndValueExpr{TypeAndValue: tv, Expr: e})
		}
	}
}

func (p *Pkg) FuncResultsOf(fn *types.Func) (Results, int) {
	if fn == nil {
		return nil, 0
	}
	decl := p.FuncDeclOf(fn)
	if decl == nil {
		return nil, 0
	}
	return p.FuncResultsOfSignature(
		fn.Type().(*types.Signature),
		decl.Body,
		decl.Type,
	)
}

func (p *Pkg) FuncResultsOfSignature(sig *types.Signature, body *ast.BlockStmt, typ *ast.FuncType) (Results, int) {
	results := sig.Results()
	if results.Len() == 0 {
		return nil, 0
	}

	named := make([]*ast.Ident, 0)

	for _, field := range typ.Results.List {
		named = append(named, field.Names...)
	}

	returns := func() []*ast.ReturnStmt {
		lst := make([]*ast.ReturnStmt, 0)
		ast.Inspect(body, func(node ast.Node) bool {
			switch node := node.(type) {
			case *ast.FuncLit:
				return false // inline declaration
			case *ast.ReturnStmt:
				lst = append(lst, node)
			}
			return true
		})
		return lst
	}()

	finals := Results{}

	for _, stmt := range returns {
		if stmt.Results != nil {
			p.SetResultsBy(finals, stmt.Results...)
			continue
		}
		for i := 0; i < results.Len(); i++ {
			// named return
			for _, tve := range p.AssignedValueOf(named[i], stmt.Pos()) {
				p.AppendResult(finals, i, tve)
			}
		}
	}

	for i := range finals {
		for j := range finals[i] {
			tve := finals[i][j]
			switch t := results.At(i).Type().(type) {
			case *types.Interface:
				// nothing
			case *types.Named:
				if t.String() != "error" {
					tve.Type = t
				}
			default:
				tve.Type = t
			}
			finals[i][j] = tve
		}
	}

	return finals, results.Len()
}

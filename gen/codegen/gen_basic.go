package codegen

import (
	"bytes"
	"fmt"
	"go/token"
	"sort"
	"strconv"
	"strings"

	"github.com/saitofun/qlib/util/qnaming"
)

type (
	Snippet interface{ Bytes() []byte }

	SnippetSpec interface {
		Snippet
		IfSpec
	}

	SnippetType interface {
		Snippet
		IfType
	}

	SnippetCanAddr interface {
		Snippet
		IfCanAddr
	}

	SnippetKVExpr struct {
		Snippet
		K, V Snippet
	}

	SnippetAddrExpr struct {
		V SnippetCanAddr
	}
)

type (
	IfSpec          interface{ _spec() }
	IfCanAddr       interface{ _canAddr() }
	IfCanBeIfMethod interface{ _canBeInterfaceMethod() }
	IfType          interface{ _type() }

	FnAlaise func(string) string
)

type SnippetLiteral string

var _ Snippet = SnippetLiteral("")

func (s SnippetLiteral) Bytes() []byte { return []byte(s) }

type SnippetLiteralCompose struct {
	Type     SnippetType
	Elements []Snippet
}

var _ Snippet = (*SnippetLiteralCompose)(nil)

func (s *SnippetLiteralCompose) Bytes() []byte {
	buf := bytes.NewBuffer(nil)

	if s.Type != nil {
		buf.Write(s.Type.Bytes())
	}

	buf.WriteRune('{')

	for _, e := range s.Elements {
		buf.WriteRune('\n')
		buf.Write(e.Bytes())
		buf.WriteRune(',')
	}

	buf.WriteRune('\n')
	buf.WriteRune('}')
	return buf.Bytes()
}

/*
	SnippetBlock code block, like
	```go
		var a = "Hello CodeGen"
		fmt.Print(a)
	```
*/
type SnippetBlock []Snippet

var _ Snippet = SnippetBlock(nil)

func (s SnippetBlock) Bytes() []byte {
	buf := bytes.NewBuffer(nil)

	for _, sn := range s {
		if sn == nil {
			continue
		}
		buf.Write(sn.Bytes())
		buf.WriteRune('\n')
	}

	return buf.Bytes()
}

/*
	SnippetBlockWithBrace code block quote with '{' and '}', like
	```go
	{
		var a = "Hello CodeGen"
		fmt.Print(a)
	}
	```
*/
type SnippetBlockWithBrace []Snippet

var _ Snippet = SnippetBlockWithBrace(nil)

func (s SnippetBlockWithBrace) Bytes() []byte {
	buf := bytes.NewBuffer(nil)

	buf.WriteRune('{')
	buf.WriteRune('\n')
	for _, sn := range s {
		if sn == nil {
			continue
		}
		buf.Write(sn.Bytes())
		buf.WriteRune('\n')
	}
	buf.WriteRune('}')

	return buf.Bytes()
}

// SnippetBuiltIn built-in symbols, like `int`, `println`
type SnippetBuiltIn string

var _ Snippet = SnippetBuiltIn("")

func (s SnippetBuiltIn) Bytes() []byte { return []byte(s) }

// SnippetIdent idenifier `a := 0 // a is a identifier`
type SnippetIdent string

var _ SnippetCanAddr = SnippetIdent("")

func (s SnippetIdent) Bytes() []byte { return []byte(s) }

func (s SnippetIdent) _canAddr() {}

func (s SnippetIdent) UpperCamelCase() *SnippetIdent {
	return Ident(qnaming.UpperCamelCase(string(s)))
}

func (s SnippetIdent) LowerCamelCase() *SnippetIdent {
	return Ident(qnaming.LowerCamelCase(string(s)))
}

func (s SnippetIdent) UpperSnakeCase() *SnippetIdent {
	return Ident(qnaming.UpperSnakeCase(string(s)))
}

func (s SnippetIdent) LowerSnakeCase() *SnippetIdent {
	return Ident(qnaming.LowerSnakeCase(string(s)))
}

// SnippetComments comment code
type SnippetComments []string

var _ Snippet = SnippetComments([]string{})

func (s SnippetComments) Bytes() []byte {
	buf := bytes.NewBuffer(nil)

	for _, c := range s {
		buf.WriteString("// ")
		buf.WriteString(c)
		buf.WriteRune('\n')
	}
	return buf.Bytes()
}

// SnippetExpr expression `a == 0 // a == 0 is a expression`
type SnippetExpr string

var _ Snippet = SnippetExpr("")

func (s SnippetExpr) Bytes() []byte { return []byte(s) }

type SnippetTypeDecl struct {
	Token token.Token
	Specs []SnippetSpec
}

var _ Snippet = (*SnippetTypeDecl)(nil)

func (s *SnippetTypeDecl) Bytes() []byte {
	buf := bytes.NewBuffer(nil)
	mul := len(s.Specs) > 1

	buf.WriteString(s.Token.String())
	buf.WriteRune(' ')

	if mul {
		buf.WriteRune('(')
		buf.WriteRune('\n')
	}

	for i, spec := range s.Specs {
		if i > 0 {
			buf.WriteRune('\n')
		}
		buf.Write(spec.Bytes())
	}

	if mul {
		buf.WriteRune('\n')
		buf.WriteRune(')')
	}

	return buf.Bytes()
}

type SnippetField struct {
	SnippetSpec
	SnippetCanAddr

	Type  SnippetType
	Names []*SnippetIdent
	Tag   string
	Alias bool

	SnippetComments
}

var _ Snippet = (*SnippetField)(nil)

func (s *SnippetField) Bytes() []byte {
	buf := bytes.NewBuffer(nil)

	if s.SnippetComments != nil {
		buf.Write(s.SnippetComments.Bytes())
	}

	for i := range s.Names {
		if i > 0 {
			buf.WriteRune(',')
			buf.WriteRune(' ')
		}
		buf.Write(s.Names[i].Bytes())
	}

	if len(s.Names) > 0 {
		if s.Alias {
			buf.WriteRune(' ')
			buf.WriteRune('=')
			buf.WriteRune(' ')
		} else {
			buf.WriteRune(' ')
		}
	}

	buf.Write(s.Type.Bytes())

	if s.Tag != "" {
		buf.WriteRune(' ')
		buf.WriteRune('`')
		buf.WriteString(s.Tag)
		buf.WriteRune('`')
	}

	return buf.Bytes()
}

func (s SnippetField) WithTag(tag string) *SnippetField {
	s.Tag = tag
	return &s
}

func (s SnippetField) WithTags(tags map[string][]string) *SnippetField {
	buf := bytes.NewBuffer(nil)

	names := make([]string, 0)
	for tag := range tags {
		names = append(names, tag)
	}
	sort.Strings(names)

	for i, tag := range names {
		if i > 0 {
			buf.WriteRune(' ')
		}
		values := make([]string, 0)
		for j := range tags[tag] {
			v := tags[tag][j]
			if v != "" {
				values = append(values, v)
			}
		}
		buf.WriteString(tag)
		buf.WriteRune(':')
		buf.WriteString(strconv.Quote(strings.Join(values, ",")))
	}

	s.Tag = buf.String()
	return &s
}

func (s SnippetField) WithoutTag() *SnippetField {
	s.Tag = ""
	return &s
}

func (s SnippetField) WithComments(cmt ...string) *SnippetField {
	s.SnippetComments = Comments(cmt...)
	return &s
}

func (s SnippetField) AsAlias() *SnippetField {
	s.Alias = true
	return &s
}

type SnippetCaseClause struct {
	Case []Snippet
	Blk  SnippetBlock
}

var _ Snippet = (*SnippetCaseClause)(nil)

func (s *SnippetCaseClause) Bytes() []byte {
	buf := bytes.NewBuffer(nil)

	if len(s.Case) == 0 {
		buf.WriteString(token.DEFAULT.String())
	} else {
		buf.WriteString(token.CASE.String() + " ")
		for i, c := range s.Case {
			if i > 0 {
				buf.WriteString(", ")
			}
			buf.Write(c.Bytes())
		}
	}
	buf.WriteRune(':')
	for _, b := range s.Blk {
		buf.WriteRune('\n')
		buf.Write(b.Bytes())
	}

	buf.WriteRune('\n')

	return buf.Bytes()
}

func (s SnippetCaseClause) Do(bodies ...Snippet) *SnippetCaseClause {
	s.Blk = bodies
	return &s
}

type SnippetTypeAssertExpr struct {
	Asserter Snippet
	Type     SnippetType
}

var _ Snippet = (*SnippetTypeAssertExpr)(nil)

func (s *SnippetTypeAssertExpr) Bytes() []byte {
	buf := bytes.NewBuffer(nil)

	buf.Write(s.Asserter.Bytes())

	buf.WriteRune('.')
	buf.WriteRune('(')
	buf.Write(s.Type.Bytes())
	buf.WriteRune(')')

	return buf.Bytes()
}

type SnippetAssignStmt struct {
	IfSpec
	Token token.Token
	Ls    []SnippetCanAddr
	Rs    []Snippet
}

var _ Snippet = (*SnippetAssignStmt)(nil)

func (s *SnippetAssignStmt) Bytes() []byte {
	buf := bytes.NewBuffer(nil)

	for i, l := range s.Ls {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.Write(l.Bytes())
	}

	if len(s.Rs) > 0 {
		buf.WriteRune(' ')
		buf.WriteString(s.Token.String())
		buf.WriteRune(' ')

		for i, r := range s.Rs {
			if i > 0 {
				buf.WriteString(", ")
			}
			buf.Write(r.Bytes())
		}
	}

	return buf.Bytes()
}

func (s SnippetAssignStmt) By(rs ...Snippet) *SnippetAssignStmt {
	s.Rs = rs
	return &s
}

type SnippetReturnStmt struct {
	Results []Snippet
}

var _ Snippet = (*SnippetReturnStmt)(nil)

func (s *SnippetReturnStmt) Bytes() []byte {
	buf := bytes.NewBuffer(nil)

	buf.WriteString("return")

	for i, r := range s.Results {
		if i > 0 {
			buf.WriteRune(',')
		}
		buf.WriteRune(' ')
		buf.Write(r.Bytes())
	}

	return buf.Bytes()
}

type SnippetSelectStmt struct {
	Clauses []*SnippetCaseClause
}

var _ Snippet = (*SnippetSelectStmt)(nil)

func (s *SnippetSelectStmt) Bytes() []byte {
	buf := bytes.NewBuffer(nil)

	buf.WriteString(token.SELECT.String())
	buf.WriteString(" {\n")

	for _, c := range s.Clauses {
		buf.Write(c.Bytes())
	}

	buf.WriteRune('}')

	return buf.Bytes()
}

type SnippetSwitchStmt struct {
	Init, Cond Snippet
	Clauses    []*SnippetCaseClause
}

var _ Snippet = (*SnippetSwitchStmt)(nil)

func (s *SnippetSwitchStmt) Bytes() []byte {
	buf := bytes.NewBuffer(nil)

	buf.WriteString(token.SWITCH.String())
	if s.Cond != nil {
		if s.Init != nil {
			buf.WriteRune(' ')
			buf.Write(s.Init.Bytes())
			buf.WriteString(";")
		}
		buf.WriteRune(' ')
		buf.Write(s.Cond.Bytes())
	}

	buf.WriteString(" {\n")

	for _, c := range s.Clauses {
		buf.Write(c.Bytes())
	}

	buf.WriteRune('}')

	return buf.Bytes()
}

func (s SnippetSwitchStmt) InitWith(init Snippet) *SnippetSwitchStmt {
	s.Init = init
	return &s
}

func (s SnippetSwitchStmt) When(clauses ...*SnippetCaseClause) *SnippetSwitchStmt {
	s.Clauses = append(s.Clauses, clauses...)
	return &s
}

type SnippetForRangeStmt struct {
	K, V   SnippetIdent
	Ranger Snippet
	Blk    SnippetBlockWithBrace
}

var _ Snippet = (*SnippetForRangeStmt)(nil)

func (s *SnippetForRangeStmt) Bytes() []byte {
	buf := bytes.NewBuffer(nil)

	buf.WriteString(token.FOR.String())
	buf.WriteRune(' ')

	if s.K != Anonymous || s.V != Anonymous {
		buf.Write(s.K.Bytes())
		buf.WriteRune(',')
		buf.WriteRune(' ')
		buf.Write(s.V.Bytes())
		buf.WriteRune(' ')
		buf.WriteString(token.DEFINE.String())
		buf.WriteRune(' ')
	}

	buf.WriteString(token.RANGE.String())
	buf.WriteRune(' ')
	buf.Write(s.Ranger.Bytes())
	buf.WriteRune(' ')
	buf.Write(s.Blk.Bytes())

	return buf.Bytes()
}

func (s SnippetForRangeStmt) Do(blk ...Snippet) *SnippetForRangeStmt {
	s.Blk = blk
	return &s
}

type SnippetForStmt struct {
	Init, Cond, Post Snippet
	Blk              SnippetBlockWithBrace
}

var _ Snippet = (*SnippetForStmt)(nil)

func (s *SnippetForStmt) Bytes() []byte {
	buf := bytes.NewBuffer(nil)

	buf.WriteString(token.FOR.String())
	if s.Init != nil {
		buf.WriteRune(' ')
		buf.Write(s.Init.Bytes())
		buf.WriteRune(';')
	}
	if s.Cond != nil {
		buf.WriteRune(' ')
		buf.Write(s.Cond.Bytes())
	}
	if s.Post != nil {
		buf.WriteRune(';')
		buf.WriteRune(' ')
		buf.Write(s.Post.Bytes())
	}
	buf.WriteRune(' ')
	buf.Write(s.Blk.Bytes())

	return buf.Bytes()
}

func (s SnippetForStmt) Do(blk ...Snippet) *SnippetForStmt {
	s.Blk = blk
	return &s
}

type SnippetIfStmt struct {
	Init, Cond Snippet
	Blk        SnippetBlockWithBrace
	Else       []*SnippetIfStmt
}

var _ Snippet = (*SnippetIfStmt)(nil)

func (s *SnippetIfStmt) Bytes() []byte {
	buf := bytes.NewBuffer(nil)

	if s.Cond != nil {
		buf.WriteString(token.IF.String())
	}

	if s.Init != nil {
		buf.WriteRune(' ')
		buf.Write(s.Init.Bytes())
		buf.WriteRune(';')
	}

	if s.Cond != nil {
		buf.WriteRune(' ')
		buf.Write(s.Cond.Bytes())
	}

	buf.WriteRune(' ')
	buf.Write(s.Blk.Bytes())

	for _, then := range s.Else {
		buf.WriteRune(' ')
		buf.WriteString(token.ELSE.String())
		if then.Cond != nil {
			buf.WriteRune(' ')
		}
		buf.Write(then.Bytes())
	}

	return buf.Bytes()
}

type SnippetRefExpr struct {
	Lead Snippet
	Refs []Snippet
}

var _ Snippet = (*SnippetRefExpr)(nil)

func (s *SnippetRefExpr) Bytes() []byte {
	buf := bytes.NewBuffer(nil)

	buf.Write(s.Lead.Bytes())
	for _, ref := range s.Refs {
		buf.WriteRune('.')
		buf.Write(ref.Bytes())
	}

	return buf.Bytes()
}

type SnippetStarExpr struct {
	SnippetCanAddr
	SnippetType
	T SnippetType
}

var _ Snippet = (*SnippetStarExpr)(nil)

func (s *SnippetStarExpr) Bytes() []byte {
	buf := bytes.NewBuffer(nil)

	buf.WriteString(token.MUL.String())
	buf.Write(s.T.Bytes())

	return buf.Bytes()
}

var _ Snippet = (*SnippetAddrExpr)(nil)

func (s *SnippetAddrExpr) Bytes() []byte {
	buf := bytes.NewBuffer(nil)

	buf.WriteString(token.AND.String())
	buf.Write(s.V.Bytes())

	return buf.Bytes()
}

type SnippetParenExpr struct {
	SnippetCanAddr
	V Snippet
}

var _ Snippet = (*SnippetParenExpr)(nil)

func (s *SnippetParenExpr) Bytes() []byte {
	buf := bytes.NewBuffer(nil)

	buf.WriteRune('(')
	buf.Write(s.V.Bytes())
	buf.WriteRune(')')

	return buf.Bytes()
}

type SnippetIncExpr struct {
	Value SnippetCanAddr
}

var _ Snippet = (*SnippetIncExpr)(nil)

func (s *SnippetIncExpr) Bytes() []byte {
	buf := bytes.NewBuffer(nil)

	buf.Write(s.Value.Bytes())
	buf.WriteString(token.INC.String())

	return buf.Bytes()
}

type SnippetDecExpr struct {
	Value SnippetCanAddr
}

var _ Snippet = (*SnippetDecExpr)(nil)

func (s *SnippetDecExpr) Bytes() []byte {
	buf := bytes.NewBuffer(nil)

	buf.Write(s.Value.Bytes())
	buf.WriteString(token.INC.String())

	return buf.Bytes()
}

type SnippetCallExpr struct {
	Callee   Snippet
	Args     []Snippet
	Ellipsis bool
	Modifier token.Token
}

var _ Snippet = (*SnippetCallExpr)(nil)

func (s *SnippetCallExpr) Bytes() []byte {
	buf := bytes.NewBuffer(nil)

	if s.Modifier > token.ILLEGAL {
		buf.WriteString(s.Modifier.String())
		buf.WriteRune(' ')
	}

	buf.Write(s.Callee.Bytes())

	buf.WriteRune('(')
	for i, a := range s.Args {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.Write(a.Bytes())
	}

	if s.Ellipsis {
		buf.WriteString(token.ELLIPSIS.String())
	}
	buf.WriteRune(')')

	return buf.Bytes()
}

func (s SnippetCallExpr) AsDefer() *SnippetCallExpr {
	s.Modifier = token.DEFER
	return &s
}

func (s SnippetCallExpr) AsRotine() *SnippetCallExpr {
	s.Modifier = token.GO
	return &s
}

func (s SnippetCallExpr) WithEllipsis() *SnippetCallExpr {
	s.Ellipsis = true
	return &s
}

type BuiltInType string

var _ SnippetType = BuiltInType("")

func (t BuiltInType) Bytes() []byte { return []byte(t) }

func (t BuiltInType) _type() {}

type MapType struct {
	SnippetType
	Tk, Tv SnippetType
}

var _ SnippetType = (*MapType)(nil)

func (m *MapType) Bytes() []byte {
	buf := bytes.NewBuffer(nil)

	buf.WriteString(token.MAP.String())
	buf.WriteRune('[')
	buf.Write(m.Tk.Bytes())
	buf.WriteRune(']')
	buf.Write(m.Tv.Bytes())

	return buf.Bytes()
}

type ArrayType struct {
	SnippetType
	T   SnippetType
	Len int
}

var _ SnippetType = (*ArrayType)(nil)

func (s *ArrayType) Bytes() []byte {
	buf := bytes.NewBuffer(nil)

	buf.WriteString(fmt.Sprintf("[%d]", s.Len))
	buf.Write(s.T.Bytes())

	return buf.Bytes()
}

type SliceType struct {
	SnippetType
	T SnippetType
}

var _ SnippetType = (*SliceType)(nil)

func (s *SliceType) Bytes() []byte {
	buf := bytes.NewBuffer(nil)

	buf.WriteRune('[')
	buf.WriteRune(']')
	buf.Write(s.T.Bytes())

	return buf.Bytes()
}

type ChanType struct {
	SnippetType
	T SnippetType
}

var _ SnippetType = (*ChanType)(nil)

func (s *ChanType) Bytes() []byte {
	buf := bytes.NewBuffer(nil)

	buf.WriteString(token.CHAN.String())
	buf.WriteRune(' ')
	buf.Write(s.T.Bytes())

	return buf.Bytes()
}

type EllipsisType struct {
	SnippetType
	T SnippetType
}

var _ SnippetType = (*EllipsisType)(nil)

func (s *EllipsisType) Bytes() []byte {
	buf := bytes.NewBuffer(nil)

	buf.WriteString(token.ELLIPSIS.String())
	buf.Write(s.T.Bytes())

	return buf.Bytes()
}

type FuncType struct {
	SnippetType
	IfCanBeIfMethod
	Name    *SnippetIdent
	Recv    *SnippetField
	Args    []*SnippetField
	Rets    []*SnippetField
	Blk     SnippetBlockWithBrace
	noToken bool
}

var _ SnippetType = (*FuncType)(nil)

func (f *FuncType) Bytes() []byte {
	buf := bytes.NewBuffer(nil)

	if !f.noToken {
		buf.WriteString(token.FUNC.String())
		buf.WriteRune(' ')
	}

	if f.Recv != nil {
		buf.WriteRune('(')
		buf.Write(f.Recv.Bytes())
		buf.WriteRune(')')
		buf.WriteRune(' ')
	}

	if f.Name != nil {
		buf.Write(f.Name.Bytes())
	}

	buf.WriteByte('(')
	for i := range f.Args {
		if i > 0 {
			buf.WriteRune(',')
			buf.WriteRune(' ')
		}
		buf.Write(f.Args[i].WithoutTag().Bytes())
	}
	buf.WriteByte(')')

	if len(f.Rets) > 1 {
		buf.WriteRune(' ')
		buf.WriteRune('(')
	}

	for i := range f.Rets {
		if i > 0 {
			buf.WriteRune(',')
			buf.WriteRune(' ')
		}
		buf.Write(f.Rets[i].WithoutTag().Bytes())
	}

	if len(f.Rets) > 1 {
		buf.WriteRune(')')
	}

	if f.Blk != nil {
		buf.WriteRune(' ')
		buf.Write(f.Blk.Bytes())
	}
	return buf.Bytes()
}

func (f FuncType) WithoutToken() *FuncType {
	f.noToken = true
	return &f
}

type StructType struct {
	SnippetType
	Fields []*SnippetField
}

var _ SnippetType = (*StructType)(nil)

func (s *StructType) Bytes() []byte {
	buf := bytes.NewBuffer(nil)

	buf.WriteString(token.STRUCT.String())

	buf.WriteRune(' ')
	buf.WriteRune('{')

	for i := range s.Fields {
		buf.WriteRune('\n')
		buf.Write(s.Fields[i].Bytes())
	}

	buf.WriteRune('\n')
	buf.WriteRune('}')

	return buf.Bytes()
}

type InterfaceType struct {
	SnippetType
	Methods []IfCanBeIfMethod
}

var _ SnippetType = (*InterfaceType)(nil)

func (s *InterfaceType) Bytes() []byte {
	buf := bytes.NewBuffer(nil)

	buf.WriteString(token.INTERFACE.String())
	buf.WriteRune(' ')
	buf.WriteRune('{')
	buf.WriteRune('\n')
	for i := range s.Methods {
		switch method := s.Methods[i].(type) {
		case *FuncType:
			buf.Write(method.WithoutToken().Bytes())
		case *NamedType:
			buf.Write(method.Bytes())
		}
		buf.WriteRune('\n')
	}
	buf.WriteRune('}')

	return buf.Bytes()
}

type NamedType struct {
	SnippetType
	IfCanBeIfMethod
	IfCanAddr
	Name *SnippetIdent
}

var _ SnippetType = (*NamedType)(nil)

func (t *NamedType) Bytes() []byte { return t.Name.Bytes() }

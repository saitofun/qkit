package codegen_test

import (
	"fmt"
	"go/token"

	. "github.com/saitofun/qkit/gen/codegen"
)

func ExampleVar() {
	var v = Var(String, "a")

	fmt.Println(string(v.Bytes()))
	fmt.Println(string(v.WithComments("A is a string").Bytes()))

	v = v.WithTag(`json:"a"`)
	fmt.Println(string(v.Bytes()))

	v = v.WithoutTag()
	fmt.Println(string(v.Bytes()))

	v = Var(String, "AliasStringType").AsAlias().
		WithComments("AliasStringType is a Alias of string")
	fmt.Println(string(v.Bytes()))

	// Output:
	// a string
	// // A is a string
	// a string
	// a string `json:"a"`
	// a string
	// // AliasStringType is a Alias of string
	// AliasStringType = string
}

func ExampleType() {

	var t = Type("Name")

	fmt.Println(string(t.Bytes()))
	fmt.Println(string(Chan(t).Bytes()))
	fmt.Println(string(Chan(Star(t)).Bytes()))
	fmt.Println(string(Array(t, 10).Bytes()))
	fmt.Println(string(Array(Star(t), 10).Bytes()))
	fmt.Println(string(Slice(t).Bytes()))
	fmt.Println(string(Map(Bool, Star(t)).Bytes()))

	// Output:
	// Name
	// chan Name
	// chan *Name
	// [10]Name
	// [10]*Name
	// []Name
	// map[bool]*Name
}

func ExampleTypeAssert() {

	ta := TypeAssert(String, Ident("value"))
	fmt.Println(string(ta.Bytes()))

	ta = TypeAssert(Type("Name"), Ident("value"))
	fmt.Println(string(ta.Bytes()))

	ta = TypeAssert(Star(Type("Name")), Ident("value"))
	fmt.Println(string(ta.Bytes()))

	// Output:
	// value.(string)
	// value.(Name)
	// value.(*Name)
}

// func Test_KeyValue(t *testing.T) {
// 	t.Log(string(KeyValue(Ident("k"), Ident("v")).Bytes()))
// }

func ExampleIdent() {
	ids := Idents("a", "b", "C", "Type.Field")
	for _, id := range ids {
		fmt.Println(string(id.Bytes()))
	}

	// Output:
	// a
	// b
	// C
	// Type.Field
}

func ExampleDefine() {
	fmt.Println(Stringify(
		Define(
			Ident("a"),
			Ident("b"),
		).By(
			Ident("c"),
			Call("Fn", Value(1)),
		)),
	)

	fmt.Println(Stringify(
		Define(
			Ident("a"),
			Ident("b"),
		).By(
			Ident("c"),
			Call("Fn", Value("abc")),
		)),
	)

	fmt.Println(Stringify(
		DeclVar(
			Var(Int, "a"),
			Assign(Var(String, "a", "b")).By(Call("Fn", Value(1))),
		),
	))

	fmt.Println(Stringify(
		DeclConst(
			Assign(Var(Int, "a")).By(Iota),
			Assign(Ident("b")),
			Assign(Ident("c")),
		),
	))

	fmt.Println(Stringify(
		DeclType(
			Var(Type("time.Time"), "aliasTime").AsAlias(),
		),
	))

	// Output:
	// a, b := c, Fn(1)
	// a, b := c, Fn("abc")
	// var (
	// a int
	// a, b string = Fn(1)
	// )
	// const (
	// a int = iota
	// b
	// c
	// )
	// type aliasTime = time.Time
}

func ExampleFor() {
	i := Ident("i")
	fmt.Println(
		Stringify(
			For(
				Define(i).By(Value(0)),
				AssignWith(token.LSS, i).By(Value(10)),
				Inc(i),
			).Do(),
		),
	)

	// Output:
	// for i := 0; i < 10; i++ {
	// }
}

func ExampleForRange() {
	k, v, ranger := Ident("k"), Ident("v"), Ident("ranger")
	fmt.Println(
		Stringify(
			ForRange(ranger, *k, *v).Do(),
		),
	)

	// Output:
	// for k, v := range ranger {
	// }
}

func ExampleSnippetBlock() {
	blk := SnippetBlock{
		SnippetExpr("a := 1"),
		SnippetExpr("b := 2"),
		SnippetExpr("a, b = b, a"),
	}
	fmt.Print(string(blk.Bytes()))
	fmt.Print(string((SnippetBlockWithBrace)(blk).Bytes()))

	// Output:
	// a := 1
	// b := 2
	// a, b = b, a
	// {
	// a := 1
	// b := 2
	// a, b = b, a
	// }
}

package pkg_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
	"testing"

	. "github.com/sincospro/qkit/x/pkg"
)

func TestCommentScanner(t *testing.T) {
	var (
		fs     = token.NewFileSet()
		src, _ = os.ReadFile("__tests__/demo.go")
		f, _   = parser.ParseFile(fs, "__tests__/demo.go", src, parser.ParseComments)
		cs     = NewCommentScanner(fs, f)
		// tt   = require.New(t)
	)

	ast.Inspect(f, func(node ast.Node) bool {
		if node == nil {
			return true
		}
		defer func() {
			recover()
		}()
		comments := cs.CommentsOf(node)
		_len := len(strings.Split(cs.CommentsOf(node), "\n"))
		t.Log(_len)
		// tt.GreaterOrEqual(3, _len)
		if _len != 3 {
			t.Log(comments)
			t.Log(_len)
			for _, c := range node.(*ast.CommentGroup).List {
				t.Log(c.Text)
			}
		}
		return true
	})
}

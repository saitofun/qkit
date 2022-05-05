package pkg

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/token"
	"strings"
)

func StringifyNode(fs *token.FileSet, n ast.Node) string {
	buf := bytes.NewBuffer(nil)
	if err := format.Node(buf, fs, n); err != nil {
		panic(err)
	}
	return buf.String()
}

func StringifyCommentGroup(lst ...*ast.CommentGroup) (ret string) {
	if len(lst) == 0 {
		return ""
	}
	for _, cg := range lst {
		for _, line := range strings.Split(cg.Text(), "\n") {
			if strings.HasPrefix(line, "go:") {
				continue
			}
			ret = ret + "\n" + line
		}
	}
	return
}

func Import(path string) string {
	parts := strings.Split(path, "/vendor/")
	return parts[len(parts)-1]
}

func ImportPathAndExpose(s string) (string, string) {
	args := strings.Split(s, ".")
	if _len := len(args); _len > 1 {
		return Import(strings.Join(args[0:_len-1], ",")), args[_len-1]
	}
	return "", s
}

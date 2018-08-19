package evaluator

import (
	"github.com/wasanx25/gopter/ast"
	"github.com/wasanx25/gopter/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}

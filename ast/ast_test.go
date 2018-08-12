package ast_test

import (
	"testing"

	"github.com/wasanx25/gopter/ast"
	"github.com/wasanx25/gopter/token"
)

func TestString(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myBar",
				},
				Value: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "let myBar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}

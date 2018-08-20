package ast_test

import (
	"reflect"
	"testing"

	"github.com/wasanx25/gopter/ast"
)

func TestModify(t *testing.T) {
	one := func() ast.Expression { return &ast.IntegerLiteral{Value: 1} }
	two := func() ast.Expression { return &ast.IntegerLiteral{Value: 2} }

	turnOneIntoTwo := func(node ast.Node) ast.Node {
		integer, ok := node.(*ast.IntegerLiteral)
		if !ok {
			return node
		}

		if integer.Value != 1 {
			return node
		}

		integer.Value = 2
		return integer
	}

	tests := []struct {
		input    ast.Node
		expected ast.Node
	}{
		{one(), two()},
		{
			&ast.Program{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{Expression: one()},
				},
			},
			&ast.Program{
				Statements: []ast.Statement{
					&ast.ExpressionStatement{Expression: two()},
				},
			},
		},
		{
			&ast.InfixExpression{Left: one(), Operator: "+", Right: two()},
			&ast.InfixExpression{Left: two(), Operator: "+", Right: two()},
		},
		{
			&ast.InfixExpression{Left: two(), Operator: "+", Right: one()},
			&ast.InfixExpression{Left: two(), Operator: "+", Right: two()},
		},
		{
			&ast.PrefixExpression{Operator: "-", Right: one()},
			&ast.PrefixExpression{Operator: "-", Right: two()},
		},
		{
			&ast.IndexExpression{Left: one(), Index: one()},
			&ast.IndexExpression{Left: two(), Index: two()},
		},
		{
			&ast.IfExpression{
				Condition: one(),
				Consequence: &ast.BlockStatement{
					Statements: []ast.Statement{
						&ast.ExpressionStatement{Expression: one()},
					},
				},
				Alternative: &ast.BlockStatement{
					Statements: []ast.Statement{
						&ast.ExpressionStatement{Expression: one()},
					},
				},
			},
			&ast.IfExpression{
				Condition: two(),
				Consequence: &ast.BlockStatement{
					Statements: []ast.Statement{
						&ast.ExpressionStatement{Expression: two()},
					},
				},
				Alternative: &ast.BlockStatement{
					Statements: []ast.Statement{
						&ast.ExpressionStatement{Expression: two()},
					},
				},
			},
		},
		{
			&ast.ReturnStatement{ReturnValue: one()},
			&ast.ReturnStatement{ReturnValue: two()},
		},
		{
			&ast.LetStatement{Value: one()},
			&ast.LetStatement{Value: two()},
		},
		{
			&ast.FunctionLiteral{
				Parameters: []*ast.Identifier{},
				Body: &ast.BlockStatement{
					Statements: []ast.Statement{
						&ast.ExpressionStatement{Expression: one()},
					},
				},
			},
			&ast.FunctionLiteral{
				Parameters: []*ast.Identifier{},
				Body: &ast.BlockStatement{
					Statements: []ast.Statement{
						&ast.ExpressionStatement{Expression: two()},
					},
				},
			},
		},
		{
			&ast.ArrayLiteral{Elements: []ast.Expression{one(), one()}},
			&ast.ArrayLiteral{Elements: []ast.Expression{two(), two()}},
		},
	}

	for _, tt := range tests {
		modified := ast.Modify(tt.input, turnOneIntoTwo)

		equal := reflect.DeepEqual(modified, tt.expected)
		if !equal {
			t.Errorf("not equal. got=%#v, want=%#v", modified, tt.expected)
		}
	}
}

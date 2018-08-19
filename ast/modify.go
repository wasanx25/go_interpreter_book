package ast

type ModifierFunc func(Node) Node

func Modify(node Node, modifier ModifierFunc) Node {
	switch modifyNode := node.(type) {
	case *Program:
		for i, statement := range modifyNode.Statements {
			modifyNode.Statements[i], _ = Modify(statement, modifier).(Statement)
		}
	case *ExpressionStatement:
		modifyNode.Expression, _ = Modify(modifyNode.Expression, modifier).(Expression)
	}

	return modifier(node)
}

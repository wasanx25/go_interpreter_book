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
	case *InfixExpression:
		modifyNode.Left, _ = Modify(modifyNode.Left, modifier).(Expression)
		modifyNode.Right, _ = Modify(modifyNode.Right, modifier).(Expression)
	case *PrefixExpression:
		modifyNode.Right, _ = Modify(modifyNode.Right, modifier).(Expression)
	case *IndexExpression:
		modifyNode.Left, _ = Modify(modifyNode.Left, modifier).(Expression)
		modifyNode.Index, _ = Modify(modifyNode.Index, modifier).(Expression)
	}

	return modifier(node)
}

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
	case *IfExpression:
		modifyNode.Condition, _ = Modify(modifyNode.Condition, modifier).(Expression)
		modifyNode.Consequence, _ = Modify(modifyNode.Consequence, modifier).(*BlockStatement)
		if modifyNode.Alternative != nil {
			modifyNode.Alternative, _ = Modify(modifyNode.Alternative, modifier).(*BlockStatement)
		}
	case *BlockStatement:
		for i := range modifyNode.Statements {
			modifyNode.Statements[i], _ = Modify(modifyNode.Statements[i], modifier).(Statement)
		}
	case *ReturnStatement:
		modifyNode.ReturnValue, _ = Modify(modifyNode.ReturnValue, modifier).(Expression)
	case *LetStatement:
		modifyNode.Value, _ = Modify(modifyNode.Value, modifier).(Expression)
	case *FunctionLiteral:
		for i := range modifyNode.Parameters {
			modifyNode.Parameters[i], _ = Modify(modifyNode.Parameters[i], modifier).(*Identifier)
		}
		modifyNode.Body, _ = Modify(modifyNode.Body, modifier).(*BlockStatement)
	}

	return modifier(node)
}

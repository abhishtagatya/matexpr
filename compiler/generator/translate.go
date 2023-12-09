package generator

import (
	"fmt"
	"matexpr/compiler/common"
)

// buildStackPostfix Creates Stacks from Abstract Syntax Tree through Post Order Traversal
// Result is a Stack in the order of Postfix Notation of our target code
func buildStackPostfix(root *common.SyntaxTreeNode) *common.StackItem {

	var opStack *common.StackItem

	stack := []*common.SyntaxTreeNode{root}

	for len(stack) > 0 {
		pq := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if pq.Left != nil {
			stack = append(stack, pq.Left)
		}

		if pq.Right != nil {
			stack = append(stack, pq.Right)
		}

		newStack := &common.StackItem{
			Object: pq,
			Prev:   opStack,
		}
		opStack = newStack
	}

	return opStack
}

// PostfixTranslate Translates the Abstract Syntax Tree through Stack
// of Postfix to Infix Notation
func PostfixTranslate(root *common.SyntaxTreeNode) string {
	var opStack = buildStackPostfix(root)
	var infixStack *common.StackItem

	var prevPrecedence = 4

	for opStack != nil {
		currStack := opStack.Object.(*common.SyntaxTreeNode)
		if currStack.Type == common.TFunction {
			var result string
			yItem := infixStack.Object.(*common.SyntaxTreeNode)
			infixStack = infixStack.Prev

			xItem := infixStack.Object.(*common.SyntaxTreeNode)
			infixStack = infixStack.Prev

			if common.PrecedenceMap[currStack.Attribute] >= prevPrecedence {
				if xItem.Type == common.PFormedExpr {
					xItem.Attribute = fmt.Sprintf("(%s)", xItem.Attribute)
				}

				if yItem.Type == common.PFormedExpr {
					yItem.Attribute = fmt.Sprintf("(%s)", yItem.Attribute)
				}

				result = fmt.Sprintf("%s %s %s",
					xItem.Attribute, common.FuncMap[currStack.Attribute], yItem.Attribute,
				)
			} else {
				result = fmt.Sprintf("%s %s %s",
					xItem.Attribute, common.FuncMap[currStack.Attribute], yItem.Attribute,
				)
			}
			prevPrecedence = common.PrecedenceMap[currStack.Attribute]

			newInfix := &common.StackItem{
				Object: &common.SyntaxTreeNode{
					Type:      common.PFormedExpr,
					Attribute: result,
				},
				Prev: infixStack,
			}

			infixStack = newInfix
		} else {
			newInfix := &common.StackItem{
				Object: currStack,
				Prev:   infixStack,
			}

			infixStack = newInfix
		}

		opStack = opStack.Prev
	}

	if infixStack != nil {
		return infixStack.Object.(*common.SyntaxTreeNode).Attribute
	}

	return "TranslationError: InfixStack is Empty"
}

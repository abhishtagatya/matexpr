package generator

import (
	"fmt"
	"matexpr/compiler/common"
)

// buildStackPostfix Creates Stacks from Abstract Syntax Tree through DFS Visits
// Result is a Stack in the order of Postfix Notation of our target code
func buildStackPostfix(root *common.ParseTreeNode) *common.StackItem {

	var opStack *common.StackItem

	stack := []*common.ParseTreeNode{root}

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
func PostfixTranslate(root *common.ParseTreeNode) string {
	var opStack = buildStackPostfix(root)
	var infixStack *common.StackItem

	var prevPrecedence = 4

	for opStack != nil {
		currStack := opStack.Object.(*common.ParseTreeNode)
		if currStack.Type == common.TFunction {
			var result string
			yItem := infixStack.Object.(*common.ParseTreeNode)
			infixStack = infixStack.Prev

			xItem := infixStack.Object.(*common.ParseTreeNode)
			infixStack = infixStack.Prev

			if common.PrecedenceMap[currStack.Attribute] > prevPrecedence {
				result = fmt.Sprintf("%s %s (%s)",
					xItem.Attribute, common.FuncMap[currStack.Attribute], yItem.Attribute,
				)
				fmt.Println(common.PrecedenceMap[currStack.Attribute], prevPrecedence, common.FuncMap[currStack.Attribute])

			} else {
				result = fmt.Sprintf("%s %s %s",
					xItem.Attribute, common.FuncMap[currStack.Attribute], yItem.Attribute,
				)
			}
			prevPrecedence = common.PrecedenceMap[currStack.Attribute]

			newInfix := &common.StackItem{
				Object: &common.ParseTreeNode{
					Type:      common.TNumeric,
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

	finalResult := infixStack.Object.(*common.ParseTreeNode).Attribute

	return finalResult
}

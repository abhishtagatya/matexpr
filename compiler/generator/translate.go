package generator

import (
	"fmt"
	"matexpr/compiler/common"
	"strings"
)

func updateString(translation *strings.Builder, x, y, operator string) {
	translation.WriteString(
		fmt.Sprintf(
			"%s %s %s",
			x, operator, y,
		),
	)
}

func buildStack(root *common.ParseTreeNode) *common.StackItem {
	var opStack *common.StackItem

	queue := []*common.ParseTreeNode{root}
	for len(queue) > 0 {
		pq := queue[0]
		if pq.Left != nil {
			queue = append(queue, pq.Left)
		}

		if pq.Right != nil {
			queue = append(queue, pq.Right)
		}

		newStack := &common.StackItem{
			Object: pq,
			Prev:   opStack,
		}
		opStack = newStack

		queue = queue[1:]
	}

	return opStack
}

func StackTranslate(root *common.ParseTreeNode) (string, error) {

	var translation strings.Builder

	var opStack = buildStack(root)
	var tempStack *common.StackItem

	var prevPrecedence = 4 // Set to Highest Precedence

	for opStack != nil {
		currStack := opStack.Object.(*common.ParseTreeNode)
		if currStack.Type == common.TFunction {
			xTemp := tempStack.Object.(*common.ParseTreeNode)
			tempStack = tempStack.Prev

			if translation.Len() > 0 {
				yTemp := translation.String()
				translation.Reset()

				if common.PrecedenceMap[currStack.Attribute] > prevPrecedence {
					yTemp = fmt.Sprintf("(%s)", yTemp)
				}
				prevPrecedence = common.PrecedenceMap[currStack.Attribute]

				updateString(&translation, xTemp.Attribute, yTemp, common.FuncMap[currStack.Attribute])

			} else {
				yTemp := tempStack.Object.(*common.ParseTreeNode).Attribute
				tempStack = tempStack.Prev

				if common.PrecedenceMap[currStack.Attribute] > prevPrecedence {
					yTemp = fmt.Sprintf("(%s)", yTemp)
				}
				prevPrecedence = common.PrecedenceMap[currStack.Attribute]

				updateString(&translation, xTemp.Attribute, yTemp, common.FuncMap[currStack.Attribute])
			}
		}

		if currStack.Type == common.TNumeric {
			newTemp := &common.StackItem{
				Object: currStack,
				Prev:   tempStack,
			}
			tempStack = newTemp
		}

		opStack = opStack.Prev
	}

	return translation.String(), nil
}

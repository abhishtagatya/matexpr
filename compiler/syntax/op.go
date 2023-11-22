package syntax

import (
	"fmt"
	"matexpr/compiler/common"
)

// EXPR -> FUNC ( EXPR , EXPR )
// EXPR -> NUMBER | EXPR

/* Deprecated */

func OpParse(tokens []common.Token) ([]string, error) {
	tokens = append(tokens, common.Token{
		Type:      common.TEndMark,
		Attribute: "$",
	})

	tailStack := &common.StackItem{
		Object: common.TEndMark,
		Prev:   nil,
	}

	topStackObj := tailStack.Object.(string)
	topTokenObj := tokens[0].Type

	iterToken := 0
	for topTokenObj != common.TEndMark {

		lookup := common.OpTable[topStackObj][topTokenObj]
		fmt.Println(topStackObj, topTokenObj, lookup)

		if lookup == "" {
			return nil, nil
		}

		if lookup == common.LLE {
			tailStack = tailStack.Prev
			topStackObj = tailStack.Object.(string)
		}

		if lookup == common.RLE {
			newStack := common.StackItem{
				Object: topTokenObj,
				Prev:   tailStack,
			}
			tailStack = &newStack
			topStackObj = tailStack.Object.(string)

			iterToken += 1
			topTokenObj = tokens[iterToken].Type
		}
	}

	for tailStack != nil {
		tailStack = tailStack.Prev
	}

	fmt.Println()

	return nil, nil
}

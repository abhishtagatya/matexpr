package syntax

import (
	"errors"
	"fmt"
	"matexpr/compiler/common"
)

/*
Rule Explanation :
* Uppercase are Non-terminals, Lowercase are Terminals

1. ROOT -> FUNC
2. FUNC -> sym_opp EXPR comma EXPR sym_clp
3. EXPR -> FUNC
4. EXPR -> number
*/

// funcExprParse Parse Function Components : FUNC -> sym_opp EXPR comma EXPR sym_clp
func funcExprParse(head *common.ParseTreeNode, tokens []common.Token) ([]common.Token, error) {

	var err error
	var currToken common.Token

	// Check for SYM_OPP (Open Parenthesis)
	currToken = tokens[0]
	if currToken.Type != common.TSymOpp {
		return tokens, errors.New(fmt.Sprintf("SyntaxError: Expected SYM_OPP, Got %s", currToken.Type))
	}

	// Check for Left Expression : Fill in head.Left (*ParseTreeNode)
	tokens, err = leftExprParse(head, tokens[1:])
	if err != nil {
		return tokens, err
	}

	// Check for SYM_COMMA (Comma)
	currToken = tokens[0]
	if currToken.Type != common.TSymComma {
		return tokens, errors.New(fmt.Sprintf("SyntaxError: Expected SYM_COMMA, Got %s", currToken.Type))
	}

	// Check for Right Expression : Fill in head.Right (*ParseTreeNode)
	tokens, err = rightExprParse(head, tokens[1:])
	if err != nil {
		return tokens, err
	}

	// Check for SYM_CLP (Closed Parenthesis)
	currToken = tokens[0]
	if currToken.Type != common.TSymClp {
		return tokens, errors.New(fmt.Sprintf("SyntaxError: Expected SYM_CLP, Got %s", currToken.Type))
	}

	return tokens[1:], nil
}

// leftExprParse Parse Left Side of Expression to Fill in head.Left
func leftExprParse(head *common.ParseTreeNode, tokens []common.Token) ([]common.Token, error) {

	var currToken = tokens[0]
	if currToken.Type == common.TNumeric {
		head.Left = &common.ParseTreeNode{
			Type:      currToken.Type,
			Attribute: currToken.Attribute,
			Left:      nil,
			Right:     nil,
		}

		return tokens[1:], nil
	}

	if currToken.Type == common.TFunction {
		var nextNode = &common.ParseTreeNode{
			Type:      currToken.Type,
			Attribute: currToken.Attribute,
		}
		head.Left = nextNode

		return funcExprParse(nextNode, tokens[1:])
	}

	return tokens, errors.New(fmt.Sprintf("SyntaxError: Expected NUMERIC or FUNCTION, Got %s", currToken.Type))
}

// rightExprParse Parse Right Side of Expression to Fill in head.Right
func rightExprParse(head *common.ParseTreeNode, tokens []common.Token) ([]common.Token, error) {

	var currToken = tokens[0]
	if currToken.Type == common.TNumeric {
		head.Right = &common.ParseTreeNode{
			Type:      currToken.Type,
			Attribute: currToken.Attribute,
		}
		return tokens[1:], nil
	}

	if currToken.Type == common.TFunction {
		var nextNode = &common.ParseTreeNode{
			Type:      currToken.Type,
			Attribute: currToken.Attribute,
		}
		head.Right = nextNode

		return funcExprParse(nextNode, tokens[1:])
	}

	return tokens, errors.New(fmt.Sprintf("SyntaxError: Expected NUMERIC or FUNCTION, Got %s", currToken.Type))
}

// rootExprParse Parse Root Expression (Look for Function Only) and Create head
func rootExprParse(head *common.ParseTreeNode, tokens []common.Token) (*common.ParseTreeNode, []common.Token, error) {

	var err error
	var currToken = tokens[0]

	if currToken.Type == common.TFunction {
		node := &common.ParseTreeNode{
			Type:      currToken.Type,
			Attribute: currToken.Attribute,
			Left:      nil,
			Right:     nil,
		}
		head = node

		tokens, err = funcExprParse(node, tokens[1:])
		return head, tokens, err
	}

	return nil, tokens, errors.New(fmt.Sprintf("SyntaxError: Expected FUNCTION, Got %s", currToken.Type))
}

// RecursiveDescent Parses Top-Down through recursion of Grammatical Rules (Rules Identified Above)
// Creates a Abstract Syntax Tree for Further Processing
func RecursiveDescent(tokens []common.Token) (*common.ParseTreeNode, error) {
	var err error
	var head *common.ParseTreeNode

	// Set End Mark and End of Line
	tokens = append(tokens, common.Token{
		Type:      common.TEndMark,
		Attribute: "$",
	})

	if head, tokens, err = rootExprParse(head, tokens); err != nil {
		return nil, err
	}

	if tokens[0].Type != common.TEndMark {
		return nil, errors.New(fmt.Sprintf("SyntaxError: Expected ENM, Got %s", tokens[0].Type))
	}

	return head, nil
}

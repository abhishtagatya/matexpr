package syntax

import (
	"errors"
	"fmt"
	"matexpr/compiler/common"
)

func funcExprParse(head *common.ParseTreeNode, tokens []common.Token) ([]common.Token, error) {

	var err error
	var currToken common.Token

	currToken = tokens[0]
	if currToken.Type != common.TSymOpp {
		return tokens, errors.New(fmt.Sprintf("SyntaxError: Expected SYM_OPP, Got %s", currToken.Type))
	}

	tokens, err = leftExprParse(head, tokens[1:])
	if err != nil {
		return tokens, err
	}

	currToken = tokens[0]
	if currToken.Type != common.TSymComma {
		return tokens, errors.New(fmt.Sprintf("SyntaxError: Expected SYM_COMMA, Got %s", currToken.Type))
	}

	tokens, err = rightExprParse(head, tokens[1:])
	if err != nil {
		return tokens, err
	}

	currToken = tokens[0]
	if currToken.Type != common.TSymClp {
		return tokens, errors.New(fmt.Sprintf("SyntaxError: Expected SYM_CLP, Got %s", currToken.Type))
	}

	return tokens[1:], nil
}

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

func RecursiveDescent(tokens []common.Token) (*common.ParseTreeNode, error) {
	var err error
	var head *common.ParseTreeNode

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

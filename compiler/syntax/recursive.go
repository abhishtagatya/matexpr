package syntax

import (
	"errors"
	"fmt"
	"matexpr/compiler/common"
)

func commaParse(tokens []common.Token, position int) (int, error) {

	if tokens[position].Type != common.TSymComma {
		return position, errors.New(fmt.Sprintf("SyntaxError: Expected SYM_COMMA, Got %s", tokens[position].Type))
	}

	return position, nil
}

func clpParse(tokens []common.Token, position int) (int, error) {

	if tokens[position].Type == common.TSymClp {
		return position, nil
	}

	return position, errors.New(fmt.Sprintf("SyntaxError: Expected SYM_CLP, Got %s", tokens[position].Type))
}

func oppParse(tokens []common.Token, position int) (int, error) {

	if tokens[position].Type != common.TSymOpp {
		return position, errors.New(fmt.Sprintf("SyntaxError: Expected SYM_OPP, Got %s", tokens[position].Type))
	}

	return position, nil
}

func numParse(tokens []common.Token, position int) (int, error) {

	if tokens[position].Type == common.TNumeric {
		return position, nil
	}

	return position, errors.New(fmt.Sprintf("SyntaxError: Expected NUMERIC, Got %s", tokens[position].Type))
}

func expParse(tokens []common.Token, position int) (int, error) {

	if tokens[position].Type == common.TNumeric {
		return numParse(tokens, position)
	}

	if tokens[position].Type == common.TFunction {
		return funcParse(tokens, position)
	}

	return position, errors.New(fmt.Sprintf("SyntaxError: Expected NUMERIC or FUNCTION, Got %s", tokens[position].Type))
}

func funcParse(tokens []common.Token, position int) (int, error) {

	var err error

	if tokens[position].Type != common.TFunction {
		return position, errors.New(fmt.Sprintf("SyntaxError: Expected FUNCTION, Got %s", tokens[position].Type))
	}

	position += 1
	if position, err = oppParse(tokens, position); err != nil {
		return position, err
	}

	position += 1
	if position, err = expParse(tokens, position); err != nil {
		return position, err
	}

	position += 1
	if position, err = commaParse(tokens, position); err != nil {
		return position, err
	}

	position += 1
	if position, err = expParse(tokens, position); err != nil {
		return position, err
	}

	position += 1
	if position, err = clpParse(tokens, position); err != nil {
		return position, err
	}

	return position, nil
}

func RecursiveParse(tokens []common.Token) error {
	var err error
	var position = 0

	tokens = append(tokens, common.Token{
		Type:      common.TEndMark,
		Attribute: "$",
	})

	if tokens[position].Type == common.TFunction {
		if position, err = funcParse(tokens, position); err != nil {
			return err
		}
	}

	position += 1
	if tokens[position].Type != common.TEndMark {
		return errors.New(fmt.Sprintf("SyntaxError: Expected END_MARK, Got %s", tokens[position].Type))
	}

	if position != len(tokens)-1 {
		return errors.New("SyntaxError: Invalid Syntax")
	}

	return nil
}

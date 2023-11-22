package lexical

import (
	"errors"
	"fmt"
	"matexpr/compiler/common"
)

// Tokenize Source Input to List of Tokens (Type, Attribute) Tuple -> ([]Token, error)
func Tokenize(source string, optimize bool) ([]common.Token, error) {

	var tokens []common.Token

	// Identify and Store Tokens as List
	var lexPos = 0
	var cSource = source
	for len(cSource) > 0 {

		// Identify Numeric Literals
		if t := common.NumRE.FindStringIndex(cSource); len(t) > 1 && t[0] == 0 {
			lexPos = t[1]
			tokens = append(tokens, common.Token{Type: common.TNumeric, Attribute: cSource[:lexPos]})
			cSource = cSource[lexPos:]
			continue
		}

		// Identify Functions
		if t := common.FuncRE.FindStringIndex(cSource); len(t) > 1 && t[0] == 0 {
			lexPos = t[1]
			tokens = append(tokens, common.Token{Type: common.TFunction, Attribute: cSource[:lexPos]})
			cSource = cSource[lexPos:]
			continue
		}

		// Identify Opening Parenthesis (Symbol)
		if t := common.OpParRE.FindStringIndex(cSource); len(t) > 1 && t[0] == 0 {
			lexPos = t[1]
			tokens = append(tokens, common.Token{Type: common.TSymOpp, Attribute: cSource[:lexPos]})
			cSource = cSource[lexPos:]
			continue
		}

		// Identify Closing Parenthesis (Symbol)
		if t := common.ClParRE.FindStringIndex(cSource); len(t) > 1 && t[0] == 0 {
			lexPos = t[1]
			tokens = append(tokens, common.Token{Type: common.TSymClp, Attribute: cSource[:lexPos]})
			cSource = cSource[lexPos:]
			continue
		}

		// Identify Comma (Symbol)
		if t := common.CommaRE.FindStringIndex(cSource); len(t) > 1 && t[0] == 0 {
			lexPos = t[1]
			tokens = append(tokens, common.Token{Type: common.TSymComma, Attribute: cSource[:lexPos]})
			cSource = cSource[lexPos:]
			continue
		}

		// Identify Comment Statement
		if t := common.CommentRE.FindStringIndex(cSource); len(t) > 1 && t[0] == 0 {
			lexPos = t[1]
			tokens = append(tokens, common.Token{Type: common.TComment, Attribute: cSource[:lexPos]})
			cSource = cSource[lexPos:]
			continue
		}

		// Identify Whitespaces
		if t := common.SpaceRE.FindStringIndex(cSource); len(t) > 1 && t[0] == 0 {
			lexPos = t[1]
			tokens = append(tokens, common.Token{Type: "WHITESPACE", Attribute: cSource[:lexPos]})
			cSource = cSource[lexPos:]
			continue
		}

		// If outside the FA or Language
		return nil, errors.New(fmt.Sprintf("LexicalError: Unidentified token type for `%s`", cSource))
	}

	// Optimize To Remove Unessential
	if optimize {
		var optTokens []common.Token
		for _, token := range tokens {
			if token.Type == common.TWhitespace {
				continue
			}

			if token.Type == common.TComment {
				continue
			}

			optTokens = append(optTokens, token)
		}

		return optTokens, nil
	}

	return tokens, nil
}

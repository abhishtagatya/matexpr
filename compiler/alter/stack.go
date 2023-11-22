package alter

// Alternative Implementation using Explicitly Stack

import (
	"fmt"
	"matexpr/compiler/common"
)

type OpStack struct {
	Operator string
	X        string
	Y        string
	Result   string
	Prev     *OpStack
}

func StackImplementation(tokens []common.Token) string {
	funcMap := map[string]string{
		"add": "+",
		"sub": "-",
		"mul": "*",
		"div": "/",
		"mod": "%",
		"pow": "^",
	}

	opPredMap := map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
		"%": 2,
		"^": 3,
	}

	var tailOp *OpStack = nil

	// Create Stack
	for _, token := range tokens {
		if token.Type == common.TFunction {
			newOp := OpStack{Operator: funcMap[token.Attribute], Prev: tailOp}
			tailOp = &newOp
		}

		if token.Type == common.TNumeric {
			if tailOp.X == "" {
				tailOp.X = token.Attribute
			} else {
				tailOp.Y = token.Attribute
			}

			if tailOp.X != "" && tailOp.Y != "" {
				tailOp.Result = fmt.Sprintf("%s %s %s", tailOp.X, tailOp.Operator, tailOp.Y)
			}
		}
	}

	// Resolve Stack
	var prevOp = tailOp
	var resolvedStack *OpStack = nil

	for tailOp != nil {

		if tailOp.X == "" {
			tailOp.X = resolvedStack.Result
			resolvedStack = resolvedStack.Prev
		}

		if tailOp.Y == "" {
			tailOp.Y = resolvedStack.Result
			resolvedStack = resolvedStack.Prev
		}

		if tailOp.X != "" && tailOp.Y != "" {

			// Handle Parenthesis
			if opPredMap[tailOp.Operator] > opPredMap[prevOp.Operator] {
				tailOp.Result = fmt.Sprintf("%s %s (%s)", tailOp.X, tailOp.Operator, tailOp.Y)
			} else {
				tailOp.Result = fmt.Sprintf("%s %s %s", tailOp.X, tailOp.Operator, tailOp.Y)
			}

			newResolved := OpStack{
				Operator: tailOp.Operator,
				X:        tailOp.X,
				Y:        tailOp.Y,
				Result:   tailOp.Result,
				Prev:     resolvedStack,
			}
			resolvedStack = &newResolved
		}

		prevOp = tailOp
		tailOp = tailOp.Prev
	}

	return prevOp.Result
}

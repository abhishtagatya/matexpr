package main

import (
	"fmt"
	"matexpr/compiler/alter"
	"matexpr/compiler/lexical"
	"matexpr/compiler/syntax"
)

func main() {
	tokens, err := lexical.Tokenize("add(55, add(55, 55))", true)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if err = syntax.RecursiveParse(tokens); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(alter.StackImplementation(tokens))
}

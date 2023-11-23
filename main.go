package main

import (
	"fmt"
	"matexpr/compiler/generator"
	"matexpr/compiler/lexical"
	"matexpr/compiler/syntax"
	"os"
)

func main() {

	if len(os.Args) <= 1 {
		return
	}

	arg := os.Args[1]
	if arg == "" {
		return
	}

	tokens, err := lexical.Tokenize(arg, true)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	pst, err := syntax.RecursiveDescent(tokens)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	translation, err := generator.StackTranslate(pst)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(translation)
}

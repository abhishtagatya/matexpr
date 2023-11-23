package main

import (
	"fmt"
	"matexpr/compiler/common"
	"matexpr/compiler/lexical"
	"matexpr/compiler/syntax"
)

func main() {
	tokens, err := lexical.Tokenize("add(5, mul(3, sub(10, pow(6, 4))))", true)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	pst, err := syntax.RecursiveDescent(tokens)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//fmt.Println(alter.StackImplementation(tokens))

	fmt.Println(pst)
	fmt.Println(pst.Left)
	fmt.Println(pst.Right)

	q := []*common.ParseTreeNode{pst}
	for len(q) > 0 {
		pq := q[0]
		if pq.Left != nil {
			//fmt.Println("Left")
			q = append(q, pq.Left)
		}

		if pq.Right != nil {
			//fmt.Println("Right")
			q = append(q, pq.Right)
		}

		fmt.Println(pq.Attribute)
		q = q[1:]
	}

}

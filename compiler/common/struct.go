package common

import "fmt"

// Token Object (Type, Attribute)
type Token struct {
	Type      string
	Attribute string
}

// StackItem Object (Object interface{}, Prev)
type StackItem struct {
	Object interface{}
	Prev   *StackItem
}

// SyntaxTreeNode Object (Type, Attribute, Left, Right)
// Used to Create Abstract Syntax Tree
type SyntaxTreeNode struct {
	Type      string
	Attribute string
	Left      *SyntaxTreeNode
	Right     *SyntaxTreeNode
}

// PrintSyntaxTree Auxiliary function to print syntax tree (BFS)
func PrintSyntaxTree(root *SyntaxTreeNode) {

	queue := []*SyntaxTreeNode{root}

	for len(queue) > 0 {
		pq := queue[0]

		if pq.Left != nil {
			fmt.Println(fmt.Sprintf("%s.Left -> %s", pq.Attribute, pq.Left.Attribute))
			queue = append(queue, pq.Left)
		}

		if pq.Right != nil {
			fmt.Println(fmt.Sprintf("%s.Right -> %s", pq.Attribute, pq.Right.Attribute))
			queue = append(queue, pq.Right)
		}

		queue = queue[1:]
	}
}

package common

type Token struct {
	Type      string
	Attribute string
}

type StackItem struct {
	Object interface{}
	Prev   *StackItem
}

type ParseTreeNode struct {
	Type      string
	Attribute string
	Left      *ParseTreeNode
	Right     *ParseTreeNode
}

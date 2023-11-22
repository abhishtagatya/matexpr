package common

type Token struct {
	Type      string
	Attribute string
}

type StackItem struct {
	Object interface{}
	Prev   *StackItem
}

type ParseNode struct {
	Type      string
	Attribute string
	Left      *ParseNode
	Right     *ParseNode
}

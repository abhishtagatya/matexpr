package common

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

// ParseTreeNode Object (Type, Attribute, Left, Right)
// Used for Parse Simulated Tree in Syntax Analyzer
// and Conversion for Direct Translation
type ParseTreeNode struct {
	Type      string
	Attribute string
	Left      *ParseTreeNode
	Right     *ParseTreeNode
}

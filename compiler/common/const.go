package common

import "regexp"

const TNumeric = "NUMERIC"       // Type for Numeric Token
const TFunction = "FUNCTION"     // Type for Function Token
const TSymOpp = "SYM_OPP"        // Type for Open Parenthesis (Symbol) Token
const TSymClp = "SYM_CLP"        // Type for Close Parenthesis (Symbol) Token
const TSymComma = "SYM_COMMA"    // Type for Comma (Symbol) Token
const TComment = "COMMENT"       // Type for Comment Token
const TWhitespace = "WHITESPACE" // Type for Whitespace Token
const TEndMark = "ENM"           // Type for End Marker ($) Token

const PFormedExpr = "PFE" // Parsed Form Expression

var FuncRE = regexp.MustCompile(`\b(add|sub|mul|div|mod|pow)\b`)                     // Regular Expression for Function
var NumRE = regexp.MustCompile(`[+-]?(\d+([.]\d*)?(e[+-]?\d+)?|[.]\d+(e[+-]?\d+)?)`) // Regular Expression for Numeric (with Scientific Notation Enabled)
var OpParRE = regexp.MustCompile(`\(`)                                               // Regular Expression for Open Parenthesis
var ClParRE = regexp.MustCompile(`\)`)                                               // Regular Expression for Closed Parenthesis
var CommaRE = regexp.MustCompile(`,`)                                                // Regular Expression for Comma
var SpaceRE = regexp.MustCompile(`\s+`)                                              // Regular Expression for Whitespace
var CommentRE = regexp.MustCompile(`/\*.*\*/`)                                       // Regular Expression for Comment

// FAdd FSub FMul FDiv FMod FPow Function Name Constants
const FAdd = "add"
const FSub = "sub"
const FMul = "mul"
const FDiv = "div"
const FMod = "mod"
const FPow = "pow"

// FuncMap Conversion of function names to mathematical symbols
var FuncMap = map[string]string{
	FAdd: "+",
	FSub: "-",
	FMul: "*",
	FDiv: "/",
	FMod: "%",
	FPow: "^",
}

// PrecedenceMap Lookup for function precedence
var PrecedenceMap = map[string]int{
	FAdd: 1,
	FSub: 1,
	FMul: 2,
	FDiv: 2,
	FMod: 2,
	FPow: 3,
}

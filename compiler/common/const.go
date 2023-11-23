package common

import "regexp"

const TNumeric = "NUMERIC"
const TFunction = "FUNCTION"
const TSymOpp = "SYM_OPP"
const TSymClp = "SYM_CLP"
const TSymComma = "SYM_COMMA"
const TComment = "COMMENT"
const TWhitespace = "WHITESPACE"
const TEndMark = "ENM"

const RLE = "<"
const LLE = ">"
const LEQ = "="

var FuncRE = regexp.MustCompile(`\b(add|sub|mul|div|mod|pow)\b`)
var NumRE = regexp.MustCompile(`[+-]?(\d+([.]\d*)?(e[+-]?\d+)?|[.]\d+(e[+-]?\d+)?)`)
var OpParRE = regexp.MustCompile(`\(`)
var ClParRE = regexp.MustCompile(`\)`)
var CommaRE = regexp.MustCompile(`,`)
var SpaceRE = regexp.MustCompile(`\s+`)
var CommentRE = regexp.MustCompile(`/\*.*\*/`)

var FuncMap = map[string]string{
	"add": "+",
	"sub": "-",
	"mul": "*",
	"div": "/",
	"mod": "%",
	"pow": "^",
}

var PrecedenceMap = map[string]int{
	"add": 1,
	"sub": 1,
	"mul": 2,
	"div": 2,
	"mod": 2,
	"pow": 3,
}

// EXPR -> FUNC( EXPR , EXPR )
// EXPR -> NUMBER | EXPR

// F -> ID T
// T -> < ( P ) >
// P -> < E , E >
// E -> F | NUM | E

var OpTable = map[string]map[string]string{
	TEndMark: {
		TFunction: RLE,
		TSymClp:   LLE,
	},
	TNumeric: {
		TSymClp:   LLE,
		TSymComma: LLE,
	},
	TFunction: {
		TSymOpp:  RLE,
		TSymClp:  LLE,
		TEndMark: LLE,
	},
	TSymOpp: {
		TNumeric:  RLE,
		TFunction: RLE,
		TSymComma: RLE, // Flat
		TSymClp:   LLE,
	},
	TSymClp: {
		TSymComma: LLE,
	},
	TSymComma: {
		TNumeric:  RLE,
		TFunction: RLE,
		TSymClp:   LLE,
	},
}

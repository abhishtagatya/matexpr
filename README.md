# MatExpr : A Formal Language Project
> Transforms Function Composition into Mathematical Expression

### Quick Run

This project is running on Golang, make sure to have installed version (~1.21).
Use the command below to run this project.

```shell
make brun EXPR="add(5, mul(3, sub(10, pow(6, 4))))"
```

The Merlin Server is running a version of Golang that supports the verification of this
project.

### Implementation

In this section, the implementation details are covered for this project. Discussing the components
of the compilers and how this program implements those components to obtain the objective.

#### Lexical Analyzer (Scanner)

The lexical analyzer is responsible for tokenizing the input string. The process of tokenization can be further
found in [tokenize.go](compiler%2Flexical%2Ftokenize.go), where it utilizes [Regular Expression](https://pkg.go.dev/regexp) to identify lexemes
of the input string. The regular expressions can be found in [const.go](compiler%2Fcommon%2Fconst.go).

```plain/txt
Input  : "add(5, mul(10,32.3e+10))"
Output : [{FUNCTION add} {SYM_OPP (} {NUMERIC 5} {SYM_COMMA ,} {FUNCTION add} {SYM_OPP (} {NUMERIC 10} {SYM_COMMA ,} {NUMERIC 32.3e+10} {SYM_CLP )} {SYM_CLP )}]
```

During incorrect scanning, the program will raise a `LexicalError` signifying that there are unknown lexemes / character
that doesnt fit any of the regular expressions.

```plain/txt
Input  : "add(5,gcd(10, 5))"
Output : "LexicalError: Unidentified token type for `gcd(10, 5))`"
```

#### Syntax Analyzer (Parser)

The parser utilizes Top-Down Parsing of Recursive Descent. This process is tailored towards the requirements grammatical rules
and is executed if tokenization is successful. The implementation can be found in [recursive.go](compiler%2Fsyntax%2Frecursive.go).

```plain/text
Rule Explanation :
* Uppercase are Non-terminals, Lowercase are Terminals

1. ROOT -> FUNC
2. FUNC -> sym_opp EXPR comma EXPR sym_clp
3. EXPR -> FUNC
4. EXPR -> number
```

The recursive descent implementation also creates an Abstract Syntax Tree
upon which requires to operate Syntax Direct Translation. Here is Tree visualized with [PrintSyntaxTree](compiler%2Fcommon%2Fstruct.go) function.

```plain\txt
Input  : "add(5,sub(10, 5))"
Output : add.Left -> 5
         add.Right -> sub
         sub.Left -> 10
         sub.Right -> 5
```

However, if the given expression does not conform to the grammatical rules, then the program will
raise a `SyntaxError` as follows.

```plain\txt
Input  : "add(5, mod(10))"
Output : "SyntaxError: Expected SYM_COMMA, Got SYM_CLP"
```

#### Syntax-Directed Translation (Postfix to Infix)

The implementation towards generating the final output of mathematical expression involves the process of traversing
the outputted Abstract Syntax Tree from the Parsing process and converting a Postfix Notation into an Infix Notation.

In [translate.go](compiler%2Fgenerator%2Ftranslate.go), the program utilizes an auxiliary function to traverse the Syntax
Tree using PostOrder Traversal and form a Postfix Notation Stack. This we then operate to generate our Infix Notation.

```plain/txt
Input          : "mul(5, pow(10, add(10, 10)))"
Postfix        : 5 10 10 10 + ^ *
Infix (Target) : 5 * 10 ^ (10 + 10)
```

### Author
- Abhishta Gatya Adyatma [(xadyata00)](mailto:xadyata00@stud.fit.vutbr.cz) - 255965
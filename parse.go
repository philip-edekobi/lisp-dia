package main

/*
type ast struct{
	car *ast // head
	cdr *ast // tail / rest
}

/* parse(["(", "+", "13", "(", "-", "12", "1", ")", ")"]
	should produce => ast {
		car: "+",
func parse(Tokens []Token) ast {}
*/

// ValueKind is the type of the value of the Token
// in the ast
type ValueKind uint

const (
	literalValue ValueKind = iota
	listValue
)

// Value represents a value of a node in the ast
type Value struct {
	Kind    ValueKind
	Literal *Token
	List    *AST
}

// AST is a simple list of values of Tokens organized in a
// tree-structured format
type AST []Value

/*
Parse transform the list of Tokens from the lexer into an ast
Parse (["(", "+", "13", "(", "-", "12", "1", ")", ")"]

	should produce => ast = [
		Value{
			Kind: literalValue,
			literal: "+",
		},
		Value: {
			Kind: literalValue,
			literal: "13"
		},
		Value: {
			Kind: listValue,
			list: [
				Value{
					Kind: literalValue,
					literal: "-"
				},
				Value{
					Kind: literalValue,
					literal: "12"
				}
				Value{
					Kind: literalValue",
					literal: "1"
				}
			]
	]
*/
func Parse(Tokens []Token, index int) (AST, int) {
	var tree AST

	for index < len(Tokens) {
		Token := Tokens[index]
		if Token.Kind == syntaxToken && Token.Value == "(" {
			child, nextIndex := Parse(Tokens, index+1)
			tree = append(tree, Value{
				Kind: listValue,
				List: &child,
			})
			index = nextIndex
			continue
		}

		if Token.Kind == syntaxToken && Token.Value == ")" {
			return tree, index + 1
		}
		tree = append(tree, Value{
			Kind:    literalValue,
			Literal: &Token,
		})
		index++
	}
	return tree, index
}

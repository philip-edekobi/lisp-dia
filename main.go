package main

import (
	"fmt"
	"os"
)

func main() {
	program, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	Tokens := Lex([]rune(string(program)))

	ast, _ := Parse(Tokens, 0)

	initializeBuiltins()
	ctx := map[string]any{}

	for _, stmt := range ast {
		fmt.Println(interpret(stmt, ctx))
	}

	// TODO: compile the AST to js? go?
	// compile(ast)
}

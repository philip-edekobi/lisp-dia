package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type parseTestCase struct {
	tokens      []Token
	expectedAST *AST
}

func TestParse(t *testing.T) {
	source1 := []rune("( + 1 2)")
	source2 := []rune("(- 12 (+ 3 4))")
	source3 := []rune("(+ 1 2)\n(+ 2 3)")
	source4 := []rune("(+ (- 1 0) 1)")
	source5 := []rune("((a 1 2) 1)")

	parseTestCases := []parseTestCase{
		{
			tokens: Lex(source5),
			expectedAST: &AST{
				{
					Kind: listValue,
					List: &AST{
						{
							Kind: listValue,
							List: &AST{
								{
									Kind: literalValue,
									Literal: &Token{
										Kind:     identifierToken,
										Value:    "a",
										Location: 2,
									},
								},
								{
									Kind: literalValue,
									Literal: &Token{
										Kind:     integerToken,
										Value:    "1",
										Location: 4,
									},
								},
								{
									Kind: literalValue,
									Literal: &Token{
										Kind:     integerToken,
										Value:    "2",
										Location: 6,
									},
								},
							},
						},
						{
							Kind: literalValue,
							Literal: &Token{
								Kind:     integerToken,
								Value:    "1",
								Location: 9,
							},
						},
					},
				},
			},
		},
		{
			tokens: Lex(source4),
			expectedAST: &AST{
				{
					Kind: listValue,
					List: &AST{
						{
							Kind: literalValue,
							Literal: &Token{
								Kind:     identifierToken,
								Value:    "+",
								Location: 1,
							},
						},
						{
							Kind: listValue,
							List: &AST{
								{
									Kind: literalValue,
									Literal: &Token{
										Kind:     identifierToken,
										Value:    "-",
										Location: 4,
									},
								},
								{
									Kind: literalValue,
									Literal: &Token{
										Kind:     integerToken,
										Value:    "1",
										Location: 6,
									},
								},
								{
									Kind: literalValue,
									Literal: &Token{
										Kind:     integerToken,
										Value:    "0",
										Location: 8,
									},
								},
							},
						},
						{
							Kind: literalValue,
							Literal: &Token{
								Kind:     integerToken,
								Value:    "1",
								Location: 11,
							},
						},
					},
				}},
		},
		{
			tokens: Lex(source3),
			expectedAST: &AST{
				{
					Kind: listValue,
					List: &AST{
						{
							Kind: literalValue,
							Literal: &Token{
								Kind:     identifierToken,
								Value:    "+",
								Location: 1,
							},
						},
						{
							Kind: literalValue,
							Literal: &Token{
								Kind:     integerToken,
								Value:    "1",
								Location: 3,
							},
						},
						{
							Kind: literalValue,
							Literal: &Token{
								Kind:     integerToken,
								Value:    "2",
								Location: 5,
							},
						},
					},
				},
				{
					Kind: listValue,
					List: &AST{
						{
							Kind: literalValue,
							Literal: &Token{
								Kind:     identifierToken,
								Value:    "+",
								Location: 9,
							},
						},
						{
							Kind: literalValue,
							Literal: &Token{
								Kind:     integerToken,
								Value:    "2",
								Location: 11,
							},
						},
						{
							Kind: literalValue,
							Literal: &Token{
								Kind:     integerToken,
								Value:    "3",
								Location: 13,
							},
						},
					},
				},
			},
		},
		{
			tokens: Lex(source2),
			expectedAST: &AST{
				{
					Kind: listValue,
					List: &AST{
						{
							Kind: literalValue,
							Literal: &Token{
								Kind:     identifierToken,
								Value:    "-",
								Location: 1,
							},
						},
						{
							Kind: literalValue,
							Literal: &Token{
								Kind:     integerToken,
								Value:    "12",
								Location: 3,
							},
						},
						{
							Kind: listValue,
							List: &AST{
								{
									Kind: literalValue,
									Literal: &Token{
										Kind:     identifierToken,
										Value:    "+",
										Location: 7,
									},
								},
								{
									Kind: literalValue,
									Literal: &Token{
										Kind:     integerToken,
										Value:    "3",
										Location: 9,
									},
								},
								{
									Kind: literalValue,
									Literal: &Token{
										Kind:     integerToken,
										Value:    "4",
										Location: 11,
									},
								},
							},
						},
					},
				},
			},
		},
		{
			tokens: Lex(source1),
			expectedAST: &AST{
				{
					Kind: listValue,
					List: &AST{
						{
							Kind: literalValue,
							Literal: &Token{
								Kind:     identifierToken,
								Value:    "+",
								Location: 2,
							},
						},
						{
							Kind: literalValue,
							Literal: &Token{
								Kind:     integerToken,
								Value:    "1",
								Location: 4,
							},
						},
						{
							Kind: literalValue,
							Literal: &Token{
								Kind:     integerToken,
								Value:    "2",
								Location: 6,
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range parseTestCases {
		ast, _ := Parse(tc.tokens, 0)
		require.Equal(t, *tc.expectedAST, ast)
	}
}

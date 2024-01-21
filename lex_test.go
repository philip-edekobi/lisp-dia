package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type LexTestCase struct {
	source         []rune
	cursor         int
	expectedCursor int
	expectedToken  Token
}

func TestEatWhitespace(t *testing.T) {
	str := []rune("  norma lly ")
	x := eatWhitespace(str, 0)
	require.Equal(t, x, 2)
}

func TestLexIntegerToken(t *testing.T) {
	sampleSrcs := []LexTestCase{
		{
			source:         []rune("foo a 123"),
			cursor:         6,
			expectedCursor: 9,
			expectedToken: Token{
				Value:    "123",
				Kind:     integerToken,
				Location: 6,
			},
		},
		{
			source:         []rune("foo 12 4"),
			cursor:         4,
			expectedCursor: 6,
			expectedToken: Token{
				Value:    "12",
				Kind:     integerToken,
				Location: 4,
			},
		},
		{
			source:         []rune("foo 12a 3"),
			cursor:         4,
			expectedCursor: 6,
			expectedToken: Token{
				Value:    "12",
				Kind:     integerToken,
				Location: 4,
			},
		},
		{
			source:         []rune("fooooo"),
			cursor:         0,
			expectedCursor: 0,
		},
		{
			source:         []rune("foo "),
			cursor:         4,
			expectedCursor: 4,
		},
	}

	for i, src := range sampleSrcs {
		cursor, toke := lexIntegerToken(src.source, src.cursor)

		switch i {
		case 3:
			require.Equal(t, cursor, src.expectedCursor)
			require.Nil(t, toke)

		case 4:
			require.Equal(t, cursor, src.expectedCursor)
			require.Nil(t, toke)

		default:
			require.Equal(t, cursor, src.expectedCursor)
			require.Equal(t, *toke, src.expectedToken)
		}
	}
}

func TestLexIndentifierToken(t *testing.T) {
	sampleSrcs := []LexTestCase{
		{
			source:         []rune("foo 123"),
			cursor:         0,
			expectedCursor: 3,
			expectedToken: Token{
				Value:    "foo",
				Kind:     identifierToken,
				Location: 0,
			},
		},
		{
			source:         []rune("+ 12 4"),
			cursor:         0,
			expectedCursor: 1,
			expectedToken: Token{
				Value:    "+",
				Kind:     identifierToken,
				Location: 0,
			},
		},
		{
			source:         []rune("foo aba 3"),
			cursor:         4,
			expectedCursor: 7,
			expectedToken: Token{
				Value:    "aba",
				Kind:     identifierToken,
				Location: 4,
			},
		},
		{
			source:         []rune("123"),
			cursor:         0,
			expectedCursor: 0,
		},
		{
			source:         []rune("foo "),
			cursor:         4,
			expectedCursor: 4,
		},
	}

	for i, src := range sampleSrcs {
		cursor, toke := lexIdentifierToken(src.source, src.cursor)

		switch i {
		case 3:
			require.Equal(t, cursor, src.expectedCursor)
			require.Nil(t, toke)

		case 4:
			require.Equal(t, cursor, src.expectedCursor)
			require.Nil(t, toke)

		default:
			require.Equal(t, cursor, src.expectedCursor)
			require.Equal(t, *toke, src.expectedToken)
		}
	}

}

func TestLex(t *testing.T) {
	sampleSource := []rune("(+ 12 31)")
	expectedOut := []string{"(", "+", "12", "31", ")"}
	var output []string

	for _, Token := range Lex(sampleSource) {
		output = append(output, Token.Value)
	}

	require.Equal(t, output, expectedOut)
}

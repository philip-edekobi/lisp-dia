package main

import (
	"unicode"
)

type tokenKind uint

const (
	syntaxToken     tokenKind = iota // eg "(", ")"
	identifierToken                  // eg "+", "add"
	integerToken
)

// Token represents a value in the program source
type Token struct {
	Value    string
	Kind     tokenKind
	Location int
}

// TODO: implement functionality to step throught the source and show
// where the Token is
// func (t Token) debug(source []rune) {}

// eatWhitespace skips the next whitespace characters in a character stream
func eatWhitespace(source []rune, cursor int) int {
	for cursor < len(source) {
		if unicode.IsSpace(source[cursor]) {
			cursor++
			continue
		}

		break
	}

	return cursor
}

func lexSyntaxToken(source []rune, cursor int) (int, *Token) {
	for cursor < len(source) {
		if source[cursor] == '(' || source[cursor] == ')' {
			return cursor + 1, &Token{
				Value:    string([]rune{source[cursor]}),
				Kind:     syntaxToken,
				Location: cursor,
			}
		}

		break
	}

	return cursor, nil
}

func lexIntegerToken(source []rune, cursor int) (int, *Token) {
	var Value []rune
	originalPosition := cursor

	for cursor < len(source) {
		char := source[cursor]
		if char >= '0' && char <= '9' {
			Value = append(Value, char)
			cursor++
			continue
		}

		break
	}

	if len(Value) == 0 {
		return cursor, nil
	}

	return cursor, &Token{
		Value:    string(Value),
		Kind:     integerToken,
		Location: originalPosition,
	}
}

func lexIdentifierToken(source []rune, cursor int) (int, *Token) {
	var Value []rune
	var leadingChar rune
	originalPosition := cursor

	for cursor < len(source) {
		char := source[cursor]
		if !unicode.IsSpace(char) && char != '(' && char != ')' {
			if !unicode.IsPrint(leadingChar) {
				if unicode.IsNumber(char) {
					break
				}
				leadingChar = char
			}

			Value = append(Value, char)
			cursor++
			continue
		} else {
			break
		}
	}

	if len(Value) == 0 {
		return cursor, nil
	}

	return cursor, &Token{
		Value:    string(Value),
		Kind:     identifierToken,
		Location: originalPosition,
	}
}

// Lex takes an input stream and returns the Tokens in that stream
// for example: "(+ 1 11)"
// lex("(+ 1 11)") should produce: ["(", "+", "1", "11", ")"]
func Lex(source []rune) []Token {
	var Tokens []Token
	var t *Token

	cursor := 0

	for cursor < len(source) {
		cursor = eatWhitespace(source, cursor)

		cursor, t = lexSyntaxToken(source, cursor)
		if t != nil {
			Tokens = append(Tokens, *t)
			continue
		}

		cursor, t = lexIntegerToken(source, cursor)
		if t != nil {
			Tokens = append(Tokens, *t)
			continue
		}

		cursor, t = lexIdentifierToken(source, cursor)
		if t != nil {
			Tokens = append(Tokens, *t)
			continue
		}

		// Nothing lexed
		if cursor < len(source) {
			panic("lexing failed")
		}
	}

	return Tokens
}

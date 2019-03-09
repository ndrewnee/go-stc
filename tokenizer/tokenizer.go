package tokenizer

import (
	"fmt"
	"unicode"
)

const (
	TypeParen  Type = "paren"
	TypeNumber Type = "number"
	TypeString Type = "string"
	TypeName   Type = "name"
)

type Type string

type Token struct {
	Type  Type
	Value string
}

func Tokenize(input string) ([]Token, error) {
	inputRunes := []rune(input)

	var tokens []Token
	current := 0

	for current < len(inputRunes) {
		char := inputRunes[current]
		switch {
		case char == '(':
			tokens = append(tokens, Token{TypeParen, "("})
			current++
		case char == ')':
			tokens = append(tokens, Token{TypeParen, ")"})
			current++
		case unicode.IsSpace(char):
			current++
		case unicode.IsNumber(char):
			value := ""
			for unicode.IsNumber(char) {
				value += string(char)
				current++
				char = inputRunes[current]
			}

			tokens = append(tokens, Token{TypeNumber, value})
		case char == '"':
			value := ""
			current++
			char = inputRunes[current]

			for char != '"' {
				value += string(char)
				current++
				char = inputRunes[current]
			}

			current++
			tokens = append(tokens, Token{TypeString, value})
		case unicode.IsLetter(char):
			value := ""
			for unicode.IsLetter(char) {
				value += string(char)
				current++
				char = inputRunes[current]
			}

			tokens = append(tokens, Token{TypeName, value})
		default:
			return nil, fmt.Errorf("tokenize failed. unknown character: %v, current: %v", string(char), current)
		}
	}

	return tokens, nil
}

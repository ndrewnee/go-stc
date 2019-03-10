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
	var tokens []Token

	inputRunes := []rune(input)
	length := len(inputRunes)

	for current := 0; current < length; {
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
				if current >= length {
					return tokens, nil
				}

				char = inputRunes[current]
			}

			tokens = append(tokens, Token{TypeNumber, value})
		case char == '"':
			value := ""

			current++
			if current >= length {
				return tokens, nil
			}

			char = inputRunes[current]
			for char != '"' {
				value += string(char)

				current++
				if current >= length {
					return tokens, nil
				}

				char = inputRunes[current]
			}

			tokens = append(tokens, Token{TypeString, value})
			current++
		case unicode.IsLetter(char):
			value := ""
			for unicode.IsLetter(char) {
				value += string(char)

				current++
				if current >= length {
					return tokens, nil
				}

				char = inputRunes[current]
			}

			tokens = append(tokens, Token{TypeName, value})
		default:
			return nil, fmt.Errorf("tokenize failed. unknown character: %v, current: %v", string(char), current)
		}
	}

	return tokens, nil
}

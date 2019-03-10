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

	for index := 0; index < length; {
		char := inputRunes[index]
		switch {
		case char == '(':
			tokens = append(tokens, Token{TypeParen, "("})
			index++
		case char == ')':
			tokens = append(tokens, Token{TypeParen, ")"})
			index++
		case unicode.IsSpace(char):
			index++
		case unicode.IsNumber(char):
			value := ""
			for unicode.IsNumber(char) {
				value += string(char)

				index++
				if index >= length {
					return tokens, nil
				}

				char = inputRunes[index]
			}

			tokens = append(tokens, Token{TypeNumber, value})
		case char == '"':
			value := ""

			index++
			if index >= length {
				return tokens, nil
			}

			char = inputRunes[index]
			for char != '"' {
				value += string(char)

				index++
				if index >= length {
					return tokens, nil
				}

				char = inputRunes[index]
			}

			tokens = append(tokens, Token{TypeString, value})
			index++
		case unicode.IsLetter(char):
			value := ""
			for unicode.IsLetter(char) {
				value += string(char)

				index++
				if index >= length {
					return tokens, nil
				}

				char = inputRunes[index]
			}

			tokens = append(tokens, Token{TypeName, value})
		default:
			return nil, fmt.Errorf("tokenize failed. unknown character: %v, index: %v", string(char), index)
		}
	}

	return tokens, nil
}

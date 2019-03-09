package compiler

import (
	"github.com/ndrewnee/go-stc/generator"
	"github.com/ndrewnee/go-stc/parser"
	"github.com/ndrewnee/go-stc/tokenizer"
	"github.com/ndrewnee/go-stc/transformer"
)

func Compile(input string) (string, error) {
	tokens, err := tokenizer.Tokenize(input)
	if err != nil {
		return "", err
	}

	ast, err := parser.Parse(tokens)
	if err != nil {
		return "", err
	}

	newAST, err := transformer.Transform(ast)
	if err != nil {
		return "", err
	}

	output, err := generator.GenerateCode(newAST)
	if err != nil {
		return "", err
	}

	return output, nil
}

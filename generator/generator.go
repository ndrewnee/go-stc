package generator

import (
	"fmt"
	"strings"

	"github.com/ndrewnee/go-stc/parser"
)

func GenerateCode(ast parser.Node) (string, error) {
	switch ast.Type {
	case parser.NodeTypeProgram:
		result := make([]string, 0, len(ast.Body))
		for _, child := range ast.Body {
			code, err := GenerateCode(child)
			if err != nil {
				return "", err
			}

			result = append(result, code)
		}

		return strings.Join(result, "\n"), nil
	case parser.NodeTypeExpressionStatement:
		code, err := GenerateCode(*ast.Expression)
		if err != nil {
			return "", err
		}

		return code + ";", nil
	case parser.NodeTypeCallExpression:
		callee, err := GenerateCode(*ast.Callee)
		if err != nil {
			return "", err
		}

		args := make([]string, 0, len(*ast.Arguments))
		for _, arg := range *ast.Arguments {
			argCode, err := GenerateCode(arg)
			if err != nil {
				return "", err
			}

			args = append(args, argCode)
		}

		return callee + "(" + strings.Join(args, ", ") + ")", nil
	case parser.NodeTypeIdentifier:
		return ast.Name, nil
	case parser.NodeTypeNumberLiteral:
		return ast.Value, nil
	case parser.NodeTypeStringLiteral:
		return `"` + ast.Value + `"`, nil
	default:
		return "", fmt.Errorf("generate code failed. unknown node type: %+v", ast)
	}
}

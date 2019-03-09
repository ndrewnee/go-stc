package generator

import (
	"fmt"
	"strings"

	"github.com/ndrewnee/go-stc/parser"
)

func GenerateCode(node parser.Node) (string, error) {
	switch node.Type {
	case parser.NodeTypeProgram:
		result := make([]string, 0, len(node.Body))
		for _, child := range node.Body {
			code, err := GenerateCode(child)
			if err != nil {
				return "", err
			}

			result = append(result, code)
		}

		return strings.Join(result, "\n"), nil
	case parser.NodeTypeExpressionStatement:
		code, err := GenerateCode(node)
		if err != nil {
			return "", err
		}

		return code + ";", nil
	case parser.NodeTypeCallExpression:
		callee, err := GenerateCode(*node.Callee)
		if err != nil {
			return "", err
		}

		args := make([]string, 0, len(*node.Arguments))
		for _, arg := range *node.Arguments {
			argCode, err := GenerateCode(arg)
			if err != nil {
				return "", err
			}

			args = append(args, argCode)
		}

		return callee + "(" + strings.Join(args, ", ") + ")", nil
	case parser.NodeTypeIdentifier:
		return node.Name, nil
	case parser.NodeTypeNumberLiteral:
		return node.Value, nil
	case parser.NodeTypeStringLiteral:
		return `"` + node.Value + `"`, nil
	default:
		return "", fmt.Errorf("generate code failed. unknown node type: %+v", node)
	}
}

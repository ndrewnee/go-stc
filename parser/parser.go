package parser

import (
	"fmt"

	"github.com/ndrewnee/go-stc/tokenizer"
)

const (
	NodeTypeProgram             NodeType = "Program"
	NodeTypeNumberLiteral       NodeType = "NumberLiteral"
	NodeTypeStringLiteral       NodeType = "StringLiteral"
	NodeTypeCallExpression      NodeType = "CallExpression"
	NodeTypeIdentifier          NodeType = "Identifier"
	NodeTypeExpressionStatement NodeType = "ExpressionStatement"
)

type NodeType string

type Node struct {
	Type       NodeType
	Value      string
	Name       string
	Body       []Node
	Params     []Node
	Context    *[]Node
	Callee     *Node
	Arguments  *[]Node
	Expression *Node
}

func Parse(tokens []tokenizer.Token) (Node, error) {
	current := 0

	ast := Node{Type: NodeTypeProgram}

	for current < len(tokens) {
		node, innerCurrent, err := walk(tokens, current)
		if err != nil {
			return Node{}, err
		}

		current = innerCurrent
		ast.Body = append(ast.Body, node)
	}

	return ast, nil
}

func walk(tokens []tokenizer.Token, current int) (Node, int, error) {
	token := tokens[current]
	switch {
	case token.Type == tokenizer.TypeNumber:
		current++
		return Node{Type: NodeTypeNumberLiteral, Value: token.Value}, current, nil
	case token.Type == tokenizer.TypeString:
		current++
		return Node{Type: NodeTypeStringLiteral, Value: token.Value}, current, nil
	case token.Type == tokenizer.TypeParen && token.Value == "(":
		current++
		token = tokens[current]

		node := Node{Type: NodeTypeCallExpression, Name: token.Value}

		current++
		token = tokens[current]

		for (token.Type != tokenizer.TypeParen) ||
			(token.Type == tokenizer.TypeParen && token.Value != ")") {

			param, innerCurrent, err := walk(tokens, current)
			if err != nil {
				return Node{}, 0, err
			}

			node.Params = append(node.Params, param)
			current = innerCurrent
			token = tokens[current]
		}

		current++
		return node, current, nil
	default:
		return Node{}, 0, fmt.Errorf("parse failed. unknown token type: %+v, current: %v", token, current)
	}
}

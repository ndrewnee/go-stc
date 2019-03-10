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
	index := 0

	ast := Node{Type: NodeTypeProgram}

	for index < len(tokens) {
		node, innerIndex, err := walk(tokens, index)
		if err != nil {
			return Node{}, err
		}

		index = innerIndex
		ast.Body = append(ast.Body, node)
	}

	return ast, nil
}

func walk(tokens []tokenizer.Token, index int) (Node, int, error) {
	token := tokens[index]
	switch {
	case token.Type == tokenizer.TypeNumber:
		index++
		return Node{Type: NodeTypeNumberLiteral, Value: token.Value}, index, nil
	case token.Type == tokenizer.TypeString:
		index++
		return Node{Type: NodeTypeStringLiteral, Value: token.Value}, index, nil
	case token.Type == tokenizer.TypeParen && token.Value == "(":
		index++
		token = tokens[index]

		node := Node{Type: NodeTypeCallExpression, Name: token.Value}

		index++
		token = tokens[index]

		for (token.Type != tokenizer.TypeParen) ||
			(token.Type == tokenizer.TypeParen && token.Value != ")") {

			param, innerIndex, err := walk(tokens, index)
			if err != nil {
				return Node{}, 0, err
			}

			node.Params = append(node.Params, param)
			index = innerIndex
			token = tokens[index]
		}

		index++
		return node, index, nil
	default:
		return Node{}, 0, fmt.Errorf("parse failed. unknown token type: %+v, index: %v", token, index)
	}
}

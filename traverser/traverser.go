package traverser

import (
	"fmt"

	"github.com/ndrewnee/go-stc/parser"
)

type VisitorMethods struct {
	Enter func(node *parser.Node, parent parser.Node)
	Exit  func(node *parser.Node, parent parser.Node)
}

type Visitor map[parser.NodeType]VisitorMethods

func Traverse(ast parser.Node, visitor Visitor) error {
	if visitor == nil {
		visitor = make(Visitor)
	}

	return traverseNode(ast, parser.Node{}, visitor)
}

func traverseArray(array []parser.Node, parent parser.Node, visitor Visitor) error {
	for _, child := range array {
		if err := traverseNode(child, parent, visitor); err != nil {
			return err
		}
	}

	return nil
}

func traverseNode(node parser.Node, parent parser.Node, visitor Visitor) error {
	methods, ok := visitor[node.Type]
	if ok && methods.Enter != nil {
		methods.Enter(&node, parent)
	}

	switch node.Type {
	case parser.NodeTypeProgram:
		if err := traverseArray(node.Body, node, visitor); err != nil {
			return err
		}
	case parser.NodeTypeCallExpression:
		if err := traverseArray(node.Params, node, visitor); err != nil {
			return err
		}
	case parser.NodeTypeNumberLiteral, parser.NodeTypeStringLiteral:
	default:
		return fmt.Errorf("traverse failed. unknown node type: %+v", node)
	}

	if ok && methods.Exit != nil {
		methods.Exit(&node, parent)
	}

	return nil
}

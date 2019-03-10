package transformer

import (
	"github.com/ndrewnee/go-stc/parser"
	"github.com/ndrewnee/go-stc/traverser"
)

func Transform(ast parser.Node) (parser.Node, error) {
	newAST := parser.Node{Type: parser.NodeTypeProgram, Body: []parser.Node{}}

	ast.Context = &newAST.Body

	visitor := traverser.Visitor{
		parser.NodeTypeNumberLiteral: traverser.VisitorMethods{
			Enter: func(node *parser.Node, parent parser.Node) {
				*parent.Context = append(*parent.Context, parser.Node{
					Type:  parser.NodeTypeNumberLiteral,
					Value: node.Value,
				})
			},
		},
		parser.NodeTypeStringLiteral: traverser.VisitorMethods{
			Enter: func(node *parser.Node, parent parser.Node) {
				*parent.Context = append(*parent.Context, parser.Node{
					Type:  parser.NodeTypeStringLiteral,
					Value: node.Value,
				})
			},
		},
		parser.NodeTypeCallExpression: traverser.VisitorMethods{
			Enter: func(node *parser.Node, parent parser.Node) {
				expression := parser.Node{
					Type:      parser.NodeTypeCallExpression,
					Arguments: new([]parser.Node),
					Callee: &parser.Node{
						Type: parser.NodeTypeIdentifier,
						Name: node.Name,
					},
				}

				node.Context = expression.Arguments

				if parent.Type != parser.NodeTypeCallExpression {
					newExpression := parser.Node{
						Type:       parser.NodeTypeExpressionStatement,
						Expression: &expression,
					}

					*parent.Context = append(*parent.Context, newExpression)
					return
				}

				*parent.Context = append(*parent.Context, expression)
			},
		},
	}

	if err := traverser.Traverse(ast, visitor); err != nil {
		return parser.Node{}, err
	}

	return newAST, nil
}

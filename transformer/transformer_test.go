package transformer

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ndrewnee/go-stc/parser"
)

func TestTransform(t *testing.T) {
	type args struct {
		ast parser.Node
	}
	tests := []struct {
		name    string
		skip    bool
		args    args
		wantErr assert.ErrorAssertionFunc
		want    parser.Node
	}{
		{
			name:    "should fail because node type is unknown",
			wantErr: assert.Error,
		},
		{
			name: "should traverse through ast",
			args: args{
				ast: parser.Node{
					Type:    parser.NodeTypeProgram,
					Context: new([]parser.Node),
					Body: []parser.Node{
						{
							Type: parser.NodeTypeCallExpression,
							Name: "add",
							Params: []parser.Node{
								{
									Type:  parser.NodeTypeNumberLiteral,
									Value: "2",
								},
								{
									Type: parser.NodeTypeCallExpression,
									Name: "subtract",
									Params: []parser.Node{
										{
											Type:  parser.NodeTypeNumberLiteral,
											Value: "4",
										},
										{
											Type:  parser.NodeTypeNumberLiteral,
											Value: "2",
										},
									},
								},
							},
						},
					},
				},
			},
			wantErr: assert.NoError,
			want: parser.Node{
				Type: parser.NodeTypeProgram,
				Body: []parser.Node{
					{
						Type: parser.NodeTypeExpressionStatement,
						Expression: &parser.Node{
							Type: parser.NodeTypeCallExpression,
							Callee: &parser.Node{
								Type: parser.NodeTypeIdentifier,
								Name: "add",
							},
							Arguments: &[]parser.Node{
								{
									Type:  parser.NodeTypeNumberLiteral,
									Value: "2",
								},
								{
									Type: parser.NodeTypeCallExpression,
									Callee: &parser.Node{
										Type: parser.NodeTypeIdentifier,
										Name: "subtract",
									},
									Arguments: &[]parser.Node{
										{
											Type:  parser.NodeTypeNumberLiteral,
											Value: "4",
										},
										{
											Type:  parser.NodeTypeNumberLiteral,
											Value: "2",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.skip {
				t.Skip()
			}

			got, err := Transform(tt.args.ast)
			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

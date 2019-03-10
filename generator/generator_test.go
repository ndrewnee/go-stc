package generator

import (
	"testing"

	"github.com/ndrewnee/go-stc/parser"
	"github.com/stretchr/testify/assert"
)

func TestGenerateCode(t *testing.T) {
	type args struct {
		ast parser.Node
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "should fail because node type is invalid",
			wantErr: assert.Error,
		},
		{
			name: "should generate correct code",
			args: args{
				ast: parser.Node{
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
			wantErr: assert.NoError,
			want:    "add(2, subtract(4, 2));",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateCode(tt.args.ast)
			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

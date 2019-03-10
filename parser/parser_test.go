package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ndrewnee/go-stc/tokenizer"
)

func TestParse(t *testing.T) {
	type args struct {
		tokens []tokenizer.Token
	}
	tests := []struct {
		name    string
		skip    bool
		args    args
		want    Node
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "should fail because token type is invalid",
			args: args{
				tokens: []tokenizer.Token{
					{
						Type: tokenizer.Type("invalid"),
					},
				},
			},
			wantErr: assert.Error,
		},
		{
			name: "should parse tokens",
			args: args{
				tokens: []tokenizer.Token{
					{
						Type:  tokenizer.TypeParen,
						Value: "(",
					},
					{
						Type:  tokenizer.TypeName,
						Value: "add",
					},
					{
						Type:  tokenizer.TypeNumber,
						Value: "2",
					},
					{
						Type:  tokenizer.TypeParen,
						Value: "(",
					},
					{
						Type:  tokenizer.TypeName,
						Value: "subtract",
					},
					{
						Type:  tokenizer.TypeNumber,
						Value: "4",
					},
					{
						Type:  tokenizer.TypeNumber,
						Value: "2",
					},
					{
						Type:  tokenizer.TypeParen,
						Value: ")",
					},
					{
						Type:  tokenizer.TypeParen,
						Value: ")",
					},
					{
						Type:  tokenizer.TypeParen,
						Value: "(",
					},
					{
						Type:  tokenizer.TypeName,
						Value: "concat",
					},
					{
						Type:  tokenizer.TypeString,
						Value: "foo",
					},
					{
						Type:  tokenizer.TypeString,
						Value: "bar",
					},
					{
						Type:  tokenizer.TypeParen,
						Value: ")",
					},
				},
			},
			wantErr: assert.NoError,
			want: Node{
				Type: NodeTypeProgram,
				Body: []Node{
					{
						Type: NodeTypeCallExpression,
						Name: "add",
						Params: []Node{
							{
								Type:  NodeTypeNumberLiteral,
								Value: "2",
							},
							{
								Type: NodeTypeCallExpression,
								Name: "subtract",
								Params: []Node{
									{
										Type:  NodeTypeNumberLiteral,
										Value: "4",
									},
									{
										Type:  NodeTypeNumberLiteral,
										Value: "2",
									},
								},
							},
						},
					},
					{
						Type: NodeTypeCallExpression,
						Name: "concat",
						Params: []Node{
							{
								Type:  NodeTypeStringLiteral,
								Value: "foo",
							},
							{
								Type:  NodeTypeStringLiteral,
								Value: "bar",
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

			got, err := Parse(tt.args.tokens)
			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

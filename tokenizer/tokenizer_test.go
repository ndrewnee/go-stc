package tokenizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenize(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		skip    bool
		args    args
		want    []Token
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "should fail because invalid token found in input",
			args: args{
				input: "input with invalid token: $",
			},
			wantErr: assert.Error,
		},
		{
			name: "should correctly tokenize input",
			args: args{
				input: `(add 2 (subtract 4 2)) (concat "foo" "bar")`,
			},
			wantErr: assert.NoError,
			want: []Token{
				{
					Type:  TypeParen,
					Value: "(",
				},
				{
					Type:  TypeName,
					Value: "add",
				},
				{
					Type:  TypeNumber,
					Value: "2",
				},
				{
					Type:  TypeParen,
					Value: "(",
				},
				{
					Type:  TypeName,
					Value: "subtract",
				},
				{
					Type:  TypeNumber,
					Value: "4",
				},
				{
					Type:  TypeNumber,
					Value: "2",
				},
				{
					Type:  TypeParen,
					Value: ")",
				},
				{
					Type:  TypeParen,
					Value: ")",
				},
				{
					Type:  TypeParen,
					Value: "(",
				},
				{
					Type:  TypeName,
					Value: "concat",
				},
				{
					Type:  TypeString,
					Value: "foo",
				},
				{
					Type:  TypeString,
					Value: "bar",
				},
				{
					Type:  TypeParen,
					Value: ")",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.skip {
				t.Skip()
			}

			got, err := Tokenize(tt.args.input)
			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

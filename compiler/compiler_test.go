package compiler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateCode(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "should fail because input is invalid",
			args: args{
				input: "$",
			},
			wantErr: assert.Error,
		},
		{
			name: "should generate correct code",
			args: args{
				input: `(add 2 (subtract 4 2)) (concat "foo" "bar")`,
			},
			wantErr: assert.NoError,
			want: `add(2, subtract(4, 2));
concat("foo", "bar");`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Compile(tt.args.input)
			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

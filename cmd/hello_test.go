/*
Copyright © 2022 y16ra

*/
package cmd

import (
	"testing"
)

func Test_reverseString(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "nomal",
			args: args{
				input: "Hello, World!",
			},
			want: "!dlroW ,olleH",
		},
		{
			name: "args",
			args: args{
				input: "Taro Hanaco",
			},
			want: "ocanaH oraT",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseString(tt.args.input); got != tt.want {
				t.Errorf("reverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}

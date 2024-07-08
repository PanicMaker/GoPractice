package stack

import (
	"testing"
)

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{s: "()"}, true},
		{"2", args{s: "()[]{}"}, true},
		{"3", args{s: "(]"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLC71(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{s: "/home/"}, "/home"},
		{"2", args{s: "/../"}, "/"},
		{"3", args{s: "/home//foo/"}, "/home/foo"},
		{"4", args{s: "/a/./b/../../c/"}, "/c"},
		{"5", args{s: "/home/user/Documents/../Pictures"}, "/home/user/Pictures"},
		{"6", args{s: "/.../a/../b/c/../d/./"}, "/.../b/d"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.args.s); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLC150(t *testing.T) {
	type args struct {
		tokens []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{tokens: []string{"2", "1", "+", "3", "*"}}, 9},
		{"2", args{tokens: []string{"4", "13", "5", "/", "+"}}, 6},
		{"3", args{tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}}, 22},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evalRPN(tt.args.tokens); got != tt.want {
				t.Errorf("evalRPN() = %v, want %v", got, tt.want)
			}
		})
	}
}

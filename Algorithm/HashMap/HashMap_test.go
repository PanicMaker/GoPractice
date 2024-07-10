package HashMap

import "testing"

func TestLC205(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{s: "egg", t: "add"}, true},
		{"2", args{s: "foo", t: "bar"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIsomorphicII(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isIsomorphic() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestLC290(t *testing.T) {
	type args struct {
		pattern string
		t       string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{pattern: "abba", t: "dog cat cat dog"}, true},
		{"2", args{pattern: "abba", t: "dog cat cat fish"}, false},
		{"3", args{pattern: "aaaa", t: "dog cat cat dog"}, false},
		{"4", args{pattern: "aaa", t: "aa aa aa aa"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordPattern(tt.args.pattern, tt.args.t); got != tt.want {
				t.Errorf("wordPattern() = %v, want %v", got, tt.want)
			}
		})
	}

}

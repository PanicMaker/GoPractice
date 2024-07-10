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

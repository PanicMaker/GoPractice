package String

import (
	"testing"
)

func TestLC13(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{s: "III"}, 3},
		{"2", args{s: "IV"}, 4},
		{"3", args{s: "IX"}, 9},
		{"4", args{s: "LVIII"}, 58},
		{"5", args{s: "MCMXCIV"}, 1994},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToIntII(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
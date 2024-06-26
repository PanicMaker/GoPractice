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
			if got := romanToIntI(tt.args.s); got != tt.want {
				t.Errorf("romanToIntI() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToIntII(tt.args.s); got != tt.want {
				t.Errorf("romanToIntII() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLC14(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{strs: []string{"flower", "flow", "flight"}}, "fl"},
		{"2", args{strs: []string{"dog", "racecar", "car"}}, ""},
		{"3", args{strs: []string{"ab", "a"}}, "a"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefixII(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefixII() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestLC58(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{s: "Hello World"}, 5},
		{"2", args{s: "   fly me   to   the moon  "}, 4},
		{"3", args{s: "luffy is still joyboy"}, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWordIII(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestLC125(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{s: "A man, a plan, a canal: Panama"}, true},
		{"2", args{s: "race a car"}, false},
		{"3", args{s: " "}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}

}

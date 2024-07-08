package String

import (
	"reflect"
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
			if got := isPalindromeII(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestLC392(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{s: "abc", t: "ahbgdc"}, true},
		{"2", args{s: "axc", t: "ahbgdc"}, false},
		{"3", args{s: "acb", t: "ahbgdc"}, false},
		{"4", args{s: "bb", t: "ahbgdc"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequenceI(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestLC76(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{s: "ADOBECODEBANC", t: "ABC"}, "BANC"},
		{"2", args{s: "a", t: "a"}, "a"},
		{"3", args{s: "a", t: "aa"}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindowII(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestLC567(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{s1: "ab", s2: "eidbaooo"}, true},
		{"2", args{s1: "ab", s2: "eidboaoo"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkInclusion(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("checkInclusion() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestLC30(t *testing.T) {
	type args struct {
		s     string
		words []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{s: "barfoothefoobarman", words: []string{"foo", "bar"}}, []int{0, 9}},
		{"2", args{s: "wordgoodgoodgoodbestword", words: []string{"word", "good", "best", "word"}}, []int{}},
		{"3", args{s: "barfoofoobarthefoobarman", words: []string{"bar", "foo", "the"}}, []int{6, 9, 12}},
		{"4", args{s: "wordgoodgoodgoodbestword", words: []string{"word", "good", "best", "good"}}, []int{8}},
		{"5", args{s: "lingmindraboofooowingdingbarrwingmonkeypoundcake", words: []string{"fooo", "barr", "wing", "ding", "wing"}}, []int{13}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSubstring(tt.args.s, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSubstring() = %v, want %v", got, tt.want)
			}
		})
	}

}

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

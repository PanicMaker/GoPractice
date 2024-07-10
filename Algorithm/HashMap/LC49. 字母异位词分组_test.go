package HashMap

import (
	"reflect"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{"1", args{strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"}}, [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
		{"2", args{strs: []string{""}}, [][]string{{""}}},
		{"3", args{strs: []string{"a"}}, [][]string{{"a"}}},
		{"4", args{strs: []string{"", ""}}, [][]string{{"", ""}}},
		{"5", args{strs: []string{"", "b", ""}}, [][]string{{"", ""}, {"b"}}},
		{"6", args{strs: []string{"ac", "c"}}, [][]string{{"ac"}, {"c"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagramsII(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}

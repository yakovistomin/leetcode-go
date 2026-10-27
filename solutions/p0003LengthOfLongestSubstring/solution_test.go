package p0003LengthOfLongestSubstring

import (
	"testing"

	"pgregory.net/rapid"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  int
	}{
		{
			name:  "first example",
			value: "abcabcbb",
			want:  3,
		},
		{
			name:  "second example",
			value: "pwwkew",
			want:  3,
		},
		{
			name:  "equal symbols",
			value: "bbbbb",
			want:  1,
		},
		{
			name:  "empty string",
			value: "",
			want:  0,
		},
		{
			name:  "one synbol",
			value: "a",
			want:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLongestSubstring(tt.value)

			if got != tt.want {
				t.Errorf(
					"lengthOfLongestSubstring(%s) = %d, want %d",
					tt.value,
					got,
					tt.want,
				)
			}
		})
	}
}

func bruteForce(s string) int {
	best := 0
	for i := range len(s) {
		var seen [128]bool
		for j := i; j < len(s); j++ {
			if seen[s[j]] {
				break
			}
			seen[s[j]] = true
			best = max(best, j-i+1)
		}
	}
	return best
}

func TestLengthOfLongestSubstringProperty(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		s := rapid.StringOfN(rapid.RuneFrom([]rune("abc")), 0, 30, -1).Draw(t, "s")

		if got, want := lengthOfLongestSubstring(s), bruteForce(s); got != want {
			t.Fatalf("got %d, want %d", got, want)
		}
	})
}

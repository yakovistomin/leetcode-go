package p0003LengthOfLongestSubstring

import "testing"

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

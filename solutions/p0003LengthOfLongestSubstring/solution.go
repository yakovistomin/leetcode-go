package p0003LengthOfLongestSubstring

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	lastIndex := make([]int, 128)
	result := 0

	idx1 := 0
	for idx2 := 0; idx2 < n; idx2++ {
		curSymbol := s[idx2]
		if lastIndex[curSymbol] > idx1 {
			idx1 = lastIndex[curSymbol]
		}
		if idx2-idx1+1 > result {
			result = idx2 - idx1 + 1
		}
		lastIndex[curSymbol] = idx2 + 1
	}
	return result
}

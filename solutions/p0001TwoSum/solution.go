package p0001TwoSum

func twoSum(nums []int, target int) []int {
	index := make(map[int]int)

	for idx, value := range nums {
		firstValue := target - value
		secondIdx, ok := index[firstValue]
		if ok {
			return []int{secondIdx, idx}
		}
		index[value] = idx
	}
	return nil
}

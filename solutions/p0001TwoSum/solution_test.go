package p0001TwoSum

import (
	"slices"
	"testing"

	"pgregory.net/rapid"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "first example",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "second example",
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			name:   "duplicate numbers",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)

			if !slices.Equal(got, tt.want) {
				t.Errorf(
					"twoSum(%v, %d) = %v, want %v",
					tt.nums,
					tt.target,
					got,
					tt.want,
				)
			}
		})
	}
}

type twoSumCase struct {
	nums   []int
	target int
}

func twoSumCases() *rapid.Generator[twoSumCase] {
	return rapid.Custom(func(t *rapid.T) twoSumCase {
		nums := rapid.SliceOfN(rapid.IntRange(-10_000, 10_000), 2, 100).Draw(t, "nums")

		i := rapid.IntRange(0, len(nums)-1).Draw(t, "i")
		j := rapid.IntRange(0, len(nums)-2).Draw(t, "j")
		if j >= i {
			j++
		}

		return twoSumCase{nums: nums, target: nums[i] + nums[j]}
	})
}

func TestTwoSumProperties(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		c := twoSumCases().Draw(t, "case")

		result := twoSum(slices.Clone(c.nums), c.target)

		if len(result) != 2 {
			t.Fatalf("expected 2 indices, got %v", result)
		}
		i, j := result[0], result[1]

		if i == j {
			t.Fatalf("indices must differ: %d, %d", i, j)
		}
		if i < 0 || i >= len(c.nums) || j < 0 || j >= len(c.nums) {
			t.Fatalf("index out of range: %d, %d (len %d)", i, j, len(c.nums))
		}
		if c.nums[i]+c.nums[j] != c.target {
			t.Fatalf("nums[%d] + nums[%d] = %d, want %d", i, j, c.nums[i]+c.nums[j], c.target)
		}
	})
}

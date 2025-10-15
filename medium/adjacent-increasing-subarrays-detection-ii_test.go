package medium

import "testing"

func TestMaxIncreasingSubarrays(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 3}, // k can be 3: [1,2,3] and [4,5,6]
		{[]int{1, 2, 1, 2, 1, 2}, 2}, // only k=2 possible
		{[]int{5, 4, 3, 2, 1}, 1},    // k=1 valid (single-element subarrays)
		{[]int{1, 2, 3, 1, 2, 3, 1, 2, 3}, 3},
		{[]int{1}, 0},
	}

	for _, tc := range tests {
		if got := maxIncreasingSubarrays(tc.nums); got != tc.want {
			t.Fatalf("maxIncreasingSubarrays(%v) = %d, want %d", tc.nums, got, tc.want)
		}
	}
}

package easy

import "testing"

func TestHasIncreasingSubarrays(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want bool
	}{
		{
			name: "basic true",
			nums: []int{1, 2, 3, 4},
			k:    2,
			want: true,
		},
		{
			name: "basic false",
			nums: []int{4, 3, 2, 1},
			k:    2,
			want: false,
		},
		{
			name: "two length-3 increasing subarrays",
			nums: []int{1, 3, 5, 4, 6, 8},
			k:    3,
			want: true,
		},
		{
			name: "one increasing one not",
			nums: []int{1, 2, 3, 2, 1, 0},
			k:    3,
			want: false,
		},
		{
			name: "equal elements not increasing",
			nums: []int{1, 1, 1, 1},
			k:    2,
			want: false,
		},
		{
			name: "k equals 1 always true when n>=2",
			nums: []int{5, 5},
			k:    1,
			want: true,
		},
		{
			name: "insufficient length",
			nums: []int{1, 2, 3},
			k:    2,
			want: false,
		},
		{
			name: "long contiguous increasing halves",
			nums: []int{1, 2, 3, 4, 5, 6},
			k:    3,
			want: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := hasIncreasingSubarrays(tc.nums, tc.k)
			if got != tc.want {
				t.Fatalf("hasIncreasingSubarrays(%v, %d) = %v, want %v", tc.nums, tc.k, got, tc.want)
			}
		})
	}
}

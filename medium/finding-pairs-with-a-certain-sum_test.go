package medium

import "testing"

func TestFindSumPairs_Basic(t *testing.T) {
	nums1 := []int{1, 1, 2}
	nums2 := []int{1, 2, 3}
	fp := Constructor(append([]int(nil), nums1...), append([]int(nil), nums2...))

	// initial: pairs that sum to 4: (1,3) appears 2*1 + (2,2) 1*1 = 3
	if got := fp.Count(4); got != 3 {
		t.Fatalf("initial Count(4) = %d, want 3", got)
	}

	// add 1 to nums2[0], nums2 becomes [2,2,3]
	fp.Add(0, 1)
	// now pairs sum to 4: (1,3) 2*1 + (2,2) 1*2 = 4
	if got := fp.Count(4); got != 4 {
		t.Fatalf("after Add Count(4) = %d, want 4", got)
	}
}

func TestFindSumPairs_EdgeAndUpdates(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{3, 3}
	fp := Constructor(append([]int(nil), nums1...), append([]int(nil), nums2...))

	// Count(4): only 1 pairs with both 3s (1*2) => 2
	if got := fp.Count(4); got != 2 {
		t.Fatalf("Count(4) = %d, want 2", got)
	}

	// decrement nums2[1] by 1 -> becomes 2, nums2 = [3,2]
	fp.Add(1, -1)

	// Count(3): 1 pairs with 2 -> 1*1 =1; 2 pairs with 1 -> 0
	if got := fp.Count(3); got != 1 {
		t.Fatalf("after Add Count(3) = %d, want 1", got)
	}
}

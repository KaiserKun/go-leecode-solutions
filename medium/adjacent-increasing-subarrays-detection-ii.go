package medium

// link: https://leetcode.cn/problems/adjacent-increasing-subarrays-detection-ii/
// thought: 动态规划，记录以i处结尾时，严格递增的数量
func maxIncreasingSubarrays(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	// inc[i]记录以i结尾的严格递增子数组的长度
	inc := make([]int, n)
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			inc[i] = inc[i-1] + 1
		} else {
			inc[i] = 0
		}
	}

	// 闭包，检查是否存在长度为k的子数组
	check := func(k int) bool {
		if 2*k > n {
			return false
		}
		for i := 0; i <= n-2*k; i++ {
			if inc[i+k-1] >= k-1 && inc[i+2*k-1] >= k-1 {
				return true
			}
		}
		return false
	}

	// 二分查找最大的k
	lo, hi := 1, n/2
	ans := 0
	for lo <= hi {
		mid := (lo + hi) / 2
		if check(mid) {
			ans = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return ans
}

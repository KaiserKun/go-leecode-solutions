package medium

// link: https://leetcode.cn/problems/product-of-array-except-self/
// thought: 从头和从尾各开始一次
func productExceptSelf(nums []int) []int {
	n := len(nums)

	// preStart中，第i个位置是前i个元素的乘积，除自身以外
	preStart := make([]int, n+1)
	preStart[0] = 1
	for i := 0; i < n; i++ {
		preStart[i+1] = preStart[i] * nums[i]
	}

	// preEnd中，第i+1个位置是i元素之后的元素乘积
	preEnd := make([]int, n+1)
	preEnd[n] = 1
	for i := n; i > 0; i-- {
		preEnd[i-1] = preEnd[i] * nums[i-1]
	}

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = preStart[i] * preEnd[i+1]
	}

	return ans
}

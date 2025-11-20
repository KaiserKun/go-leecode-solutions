package medium

// link: https://leetcode.cn/problems/contiguous-array/
// thought: 将0视为-1，这样当前缀和为0时，代表前面i个元素具有相同数量的0和1
func findMaxLength(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	// prefix[i] 是前缀和为i时的首次出现位置
	prefix := make(map[int]int)
	prefix[0] = -1

	// pre表示前n个元素的和
	pre := 0
	ans := 0
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			pre -= 1
		} else {
			pre += 1
		}

		if start, ok := prefix[pre]; ok {
			ans = max(ans, i-start)
		} else {
			prefix[pre] = i
		}
	}

	return ans
}

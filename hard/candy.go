package hard

// link: https://leetcode.cn/problems/candy/
// thought: 两次遍历，分别处理递增和递减的情况
func candy(ratings []int) int {
	n := len(ratings)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		cnt[i] = 1
	}
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			cnt[i] = cnt[i-1] + 1
		}
	}
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			// 取最大值，同时满足左右两个条件
			cnt[i] = max(cnt[i], cnt[i+1]+1)
		}
	}
	ans := 0
	for _, v := range cnt {
		ans += v
	}
	return ans
}

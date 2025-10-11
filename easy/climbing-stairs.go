package easy

// link: https://leetcode.cn/problems/climbing-stairs/
// thought:到达当前楼梯的方法数目是到达前两层的方法数目之和
func climbStairs(n int) int {
	// 初始值第一个楼梯有1种方法，此时前面两个楼梯无法到达，所以都为0
	p, q, r := 0, 0, 1
	for i := 0; i < n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}

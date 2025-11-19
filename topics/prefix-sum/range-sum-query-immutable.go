package ps

// link: https://leetcode.cn/problems/range-sum-query-immutable/
// thought: 利用数组prefix[i]记录前i个元素的和

type NumArray struct {
	prefix []int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	prefix := make([]int, n+1)
	prefix[0] = 0
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}

	return NumArray{prefix}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.prefix[right+1] - this.prefix[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */

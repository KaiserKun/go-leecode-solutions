package medium

// link: https://leetcode.cn/problems/maximum-number-of-events-that-can-be-attended/
// thought: 定义一个结构体，包含两个数组和一个哈希表，哈希表用于存储nums2中每个元素的出现次数
type FindSumPairs struct {
	nums1 []int
	nums2 []int
	cnt   map[int]int
}

// thought: 使用哈希表存储nums2中每个元素的出现次数，这样在count方法中可以快速查找满足条件的元素个数
func Constructor(nums1 []int, nums2 []int) FindSumPairs {
	mp := make(map[int]int)

	for _, num := range nums2 {
		mp[num]++
	}
	return FindSumPairs{nums1, nums2, mp}
}

// thought: 在Add方法中，先减少旧值的计数，再增加新值的计数
func (this *FindSumPairs) Add(index int, val int) {
	t := this.nums2[index]
	this.cnt[t]--
	this.nums2[index] += val
	this.cnt[this.nums2[index]]++
}

// thought: 在count方法中，遍历nums1，对于每个元素num，计算rest=tot-num，然后在哈希表中查找rest的出现次数并累加
func (this *FindSumPairs) Count(tot int) int {
	ans := 0
	for _, num := range this.nums1 {
		rest := tot - num
		ans += this.cnt[rest]
	}
	return ans
}

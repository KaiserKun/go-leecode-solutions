package easy

// Link: https://leetcode.cn/problems/two-sum/
// thought: 用哈希表存储已经遍历过的数字及其索引，遍历数组时检查目标值减去当前数字是否在哈希表中
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if p, ok := m[target-v]; ok {
			return []int{p, i}
		}
		m[v] = i
	}
	return nil
}

package medium

// link: https://leetcode.com/problems/search-in-rotated-sorted-array/
// thought: 从左上角或者右下角开始，这样一次可以排除一行或者一列
func findTargetIn2DPlants(plants [][]int, target int) bool {
	if len(plants) == 0 {
		return false
	}
	m, n := len(plants), len(plants[0])
	if n == 0 {
		return false
	}
	for i, j := 0, n-1; i < m && j >= 0; {
		if plants[i][j] == target {
			return true
		}
		if plants[i][j] < target {
			i++
		} else {
			j--
		}
	}
	return false
}

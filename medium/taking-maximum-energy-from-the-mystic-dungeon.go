package medium

// link: https://leetcode.cn/problems/taking-maximum-energy-from-the-mystic-dungeon/
// thought: 从终点往前走，每次遍历n/k个，遍历k次，也就是O(n),每次保留最大值和现在值；

import "math"

func maximumEnergy(energy []int, k int) int {
	if len(energy) == 0 {
		return energy[0]
	}

	maxGlobal := math.MinInt64
	for i := 0; i < k; i++ {
		cur := 0
		maxCur := math.MinInt64
		for j := len(energy) - i; j > 0; j -= k {
			cur += energy[j-1]
			maxCur = max(cur, maxCur)
		}
		maxGlobal = max(maxGlobal, maxCur)
	}

	return maxGlobal
}

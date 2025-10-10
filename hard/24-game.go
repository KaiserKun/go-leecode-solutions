package hard

import "math"

// link: https://leetcode.cn/problems/24-game/
// thought: 回溯+枚举所有可能的数字和运算符组合
func judgePoint24(cards []int) bool {
	floatcards := make([]float64, 4)
	for i, v := range cards {
		floatcards[i] = float64(v)
	}
	return judgePoint24ByFloat(floatcards)
}

func judgePoint24ByFloat(cards []float64) bool {
	// 退出递归的情况
	if len(cards) == 1 {
		return math.Abs(cards[0]-24) < 1e-6
	}
	for i := 0; i < len(cards); i++ {
		for j := i + 1; j < len(cards); j++ {
			x, y := cards[i], cards[j]
			var next []float64
			for k := 0; k < len(cards); k++ {
				if k != i && k != j {
					next = append(next, cards[k])
				}
			}
			// 生成所有可能的数字和运算符组合
			for _, v := range []float64{x + y, x - y, y - x, x * y} {
				if judgePoint24ByFloat(append(next, v)) {
					return true
				}
			}
			if y != 0 {
				if judgePoint24ByFloat(append(next, x/y)) {
					return true
				}
			}
			if x != 0 {
				if judgePoint24ByFloat(append(next, y/x)) {
					return true
				}
			}
		}
	}
	return false
}

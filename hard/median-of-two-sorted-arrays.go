package hard

// link: https://leetcode.cn/problems/median-of-two-sorted-arrays/
// thought: 二分查找，确保在较短的数组上进行二分查找，计算出两个数组的分割点，然后根据分割点计算中位数
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 确保 nums1 是较短的数组
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m, n := len(nums1), len(nums2)
	total := m + n
	half := (total + 1) / 2 // 左半部分元素个数

	i := m / 2
	for i >= 0 && i <= m {
		// 在 nums2 中的对应位置
		j := half - i

		// nums1 分割点左侧是否比 nums2 右侧大
		if i > 0 && j < n && nums1[i-1] > nums2[j] {
			// 分割点需要左移
			i--
		} else if j > 0 && i < m && nums2[j-1] > nums1[i] {
			// 检查 nums2 分割点左侧是否比 nums1 右侧大
			// 分割点需要右移
			i++
		} else {
			// 找到正确的分割点
			var maxLeft, minRight float64

			// 找出左侧最大值
			if i == 0 {
				maxLeft = float64(nums2[j-1])
			} else if j == 0 {
				maxLeft = float64(nums1[i-1])
			} else {
				maxLeft = float64(max(nums1[i-1], nums2[j-1]))
			}

			// 如果总元素数是奇数，中位数就是左侧最大值
			if total%2 == 1 {
				return maxLeft
			}

			// 如果总元素数是偶数，需要找出右侧最小值
			if i == m {
				minRight = float64(nums2[j])
			} else if j == n {
				minRight = float64(nums1[i])
			} else {
				minRight = float64(min(nums1[i], nums2[j]))
			}

			// 返回两个中间值的平均数
			return (maxLeft + minRight) / 2
		}
	}

	// 永远不会执行到这里，但需要返回语句
	return 0
}

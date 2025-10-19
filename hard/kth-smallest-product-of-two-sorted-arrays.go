package hard

import "sort"

// link: https://leetcode.cn/problems/kth-smallest-product-of-two-sorted-arrays/
// thought: 二分查找乘积的值x，使得乘积小于等于x的对数不少于k
// countLessOrEqual 计算对于给定的x1，有多少个nums2中的元素与x1的乘积小于等于v
func countLessOrEqual(nums2 []int, x1 int64, v int64) int {
	n2 := len(nums2)
	if x1 >= 0 {
		// 返回乘积小于等于v的对数
		return sort.Search(n2, func(i int) bool { return int64(nums2[i])*x1 > v })
	}
	left, right := 0, n2-1
	for left <= right {
		mid := (left + right) / 2
		prod := int64(nums2[mid]) * x1
		if prod > v {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return n2 - left
}

func kthSmallestProduct(nums1 []int, nums2 []int, k int64) int64 {
	n1 := len(nums1)
	n2 := len(nums2)
	min1 := int64(nums1[0])
	max1 := int64(nums1[n1-1])
	min2 := int64(nums2[0])
	max2 := int64(nums2[n2-1])
	candidates := []int64{min1 * min2, min1 * max2, max1 * min2, max1 * max2}
	minProd, maxProd := candidates[0], candidates[0]
	for _, c := range candidates[1:] {
		if c < minProd {
			minProd = c
		}
		if c > maxProd {
			maxProd = c
		}
	}
	left, right := minProd, maxProd
	for left <= right {
		mid := (left + right) / 2
		count := int64(0)
		for i := 0; i < n1; i++ {
			count += int64(countLessOrEqual(nums2, int64(nums1[i]), mid))
		}
		if count < k {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

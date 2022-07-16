/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	intersection := make([]int, 0)
	set1 := make(map[int]bool)
	for _, v := range nums1 {
		set1[v] = true
	}
	set2 := make(map[int]bool)
	for _, v := range nums2 {
		set2[v] = true
	}
	if len(set1) > len(set2) {
		set1, set2 = set2, set1
	}
	for v := range set1 {
		if _, has := set2[v]; has {
			intersection = append(intersection, v)
		}
	}
	return intersection
}

// @lc code=end


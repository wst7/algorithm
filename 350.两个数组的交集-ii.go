/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	hash := map[int]int{}
	result := []int{}

	for _, v := range nums1 {
		if hash[v] > 0 {
			hash[v]++
		} else {
			hash[v] = 1
		}
	}
	for _, v := range nums2 {
		if hash[v] > 0 {
			result = append(result, v)
			hash[v]--
		}
	}

	return result

}

// @lc code=end


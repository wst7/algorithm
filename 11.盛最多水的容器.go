/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start

func maxArea(height []int) int {
	left, right := 0, len(height) - 1
	var res int = 0
	for left < right {
		var area int
		if height[left] < height[right] {
			area = height[left] * (right - left)
			left ++
		} else {
			area = height[right] * (right - left)
			right --
		}
		if area > res {
			res = area
		}
	}
	return res
	
}

// @lc code=end


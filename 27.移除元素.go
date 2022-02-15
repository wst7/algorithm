/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	var n int = 0 // 指针移动
	for _, v := range nums {
		if v != val {
			nums[n] = v
			n++
		}
	}
	nums = nums[:n]
	return n
}

// @lc code=end


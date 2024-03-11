/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 *
 * https://leetcode.cn/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (42.00%)
 * Likes:    7031
 * Dislikes: 0
 * Total Accepted:    1.1M
 * Total Submissions: 2.6M
 * Testcase Example:  '[1,3]\n[2]'
 *
 * 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
 *
 * 算法的时间复杂度应该为 O(log (m+n)) 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums1 = [1,3], nums2 = [2]
 * 输出：2.00000
 * 解释：合并数组 = [1,2,3] ，中位数 2
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums1 = [1,2], nums2 = [3,4]
 * 输出：2.50000
 * 解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
 *
 *
 *
 *
 *
 *
 * 提示：
 *
 *
 * nums1.length == m
 * nums2.length == n
 * 0 <= m <= 1000
 * 0 <= n <= 1000
 * 1 <= m + n <= 2000
 * -10^6 <= nums1[i], nums2[i] <= 10^6
 *
 *
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	m, n := len(nums1), len(nums2)
	length := m + n
	var nums = make([]int, length)
	p1, p2, k := 0, 0, 0
	for {
		if !(p1 < m || p2 < n) {
			break
		}
		if p1 == m {
			nums[k] = nums2[p2]
			p2++
		} else if p2 == n {
			nums[k] = nums1[p1]
			p1++
		} else if nums1[p1] < nums2[p2] {
			nums[k] = nums1[p1]
			p1++
		} else {
			nums[k] = nums2[p2]
			p2++
		}
		k++
	}

	var median float64
	if length%2 == 1 {
		median = float64(nums[(length-1)/2])
	} else {
		median = float64(nums[length/2-1]+nums[length/2]) / 2
	}
	return median

}

// @lc code=end



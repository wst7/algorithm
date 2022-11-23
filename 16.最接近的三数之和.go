/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
// 确定第一个数字，寻找剩余两数之和最接近target - nums[i]

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	numsSlice := nums[:]
	sort.Ints(numsSlice)

	const INT_MAX = int(^uint(0) >> 1)
	delta := INT_MAX
	var i int
	for i = 0; i < len(numsSlice)-2; i++ {
		sum := numsSlice[i] + twoSumClosest(numsSlice, i+1, target-numsSlice[i])
		if math.Abs(float64(delta)) > math.Abs(float64(target-sum)) {
			delta = target - sum
		}
	}
	return target - delta
}

func twoSumClosest(nums []int, start int, target int) int {
	lo, hi := start, len(nums)-1
	const INT_MAX = int(^uint(0) >> 1)
	delta := INT_MAX
	for {
		if lo >= hi {
			break
		}
		sum := nums[lo] + nums[hi]
		if math.Abs(float64(delta)) > math.Abs(float64(target-sum)) {
			delta = target - sum
		}
		if sum < target {
			lo++
		} else {
			hi--
		}
	}
	return target - delta

}

// @lc code=end


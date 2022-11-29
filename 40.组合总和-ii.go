/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int
	var track []int

	var trackback func(candidates []int, target int, sum int, startIndex int)
	trackback = func(candidates []int, target int, sum int, startIndex int) {

		if sum == target {
			var temp []int
			temp = append(temp, track...)
			res = append(res, temp)
			return
		}
		if sum > target {
			return
		}

		for i := startIndex; i < len(candidates); i++ {
			if i > startIndex && candidates[i] == candidates[i-1] {
				continue
			}
			sum += candidates[i]
			track = append(track, candidates[i])
			trackback(candidates, target, sum, i+1)
			sum -= candidates[i]
			track = track[:len(track)-1]
		}
	}
	trackback(candidates, target, 0, 0)
	return res

}

// @lc code=end


/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {

	// 最后的结果
	var res [][]int
	// 记录回溯的路径
	var track []int
	backtrack(candidates, 0, target, 0, track, &res)
	return res
}

func backtrack(candidates []int, start int, target int, sum int, track []int, res *[][]int) {
	if target == sum {
		tmp := make([]int, len(track))
		copy(tmp, track) //拷贝
		*res = append(*res, tmp)
		return
	}
	if sum > target {
		return
	}
	for i := start; i < len(candidates); i++ {
		// 选择 candidates[i]
		track = append(track, candidates[i])
		sum += candidates[i]
		backtrack(candidates, i, target, sum, track, res)
		// 撤销选择 candidates[i]
		track = track[:len(track)-1]
		sum -= candidates[i]
	}
}

// @lc code=end


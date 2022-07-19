/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {
	if n <= 0 {
		return []string{""}
	}
	res := []string{}
	var dfs func(paths string, left int, right int)
	dfs = func(paths string, left int, right int) {
		if left > n || right > left {
			return
		}
		if len(paths) == n*2 {
			res = append(res, paths)
			return
		}
		dfs(paths+"(", left+1, right)
		dfs(paths+")", left, right+1)

	}

	dfs("", 0, 0)
	return res
}

// @lc code=end


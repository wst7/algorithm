/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 *
 * https://leetcode.cn/problems/add-binary/description/
 *
 * algorithms
 * Easy (53.01%)
 * Likes:    1169
 * Dislikes: 0
 * Total Accepted:    378.3K
 * Total Submissions: 713.8K
 * Testcase Example:  '"11"\n"1"'
 *
 * 给你两个二进制字符串 a 和 b ，以二进制字符串的形式返回它们的和。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入:a = "11", b = "1"
 * 输出："100"
 *
 * 示例 2：
 *
 *
 * 输入：a = "1010", b = "1011"
 * 输出："10101"
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= a.length, b.length <= 10^4
 * a 和 b 仅由字符 '0' 或 '1' 组成
 * 字符串如果不是 "0" ，就不含前导零
 *
 *
 */

// @lc code=start
func addBinary(a string, b string) string {
	var m = len(a) - 1
	var n = len(b) - 1
	var carry int = 0
	var str string = ""
	for m >= 0 || n >= 0 {
		if m >= 0 && a[m] == '1' {
			carry++
		}
		if n >= 0 && b[n] == '1' {
			carry++
		}
		str = strconv.Itoa(carry%2) + str
		carry = carry / 2
		m--
		n--
	}
	if carry != 0 {
		str = strconv.Itoa(carry) + str
	}
	return str
}

// @lc code=end


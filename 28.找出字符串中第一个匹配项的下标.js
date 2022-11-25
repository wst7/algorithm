/*
 * @lc app=leetcode.cn id=28 lang=javascript
 *
 * [28] 找出字符串中第一个匹配项的下标
 */
// abcde
// abc
// @lc code=start
var strStr = function(haystack, needle) {
	let len = needle.length
	let i = 0
	while(i <= haystack.length - len) {
		if (haystack.slice(i, i + len) === needle) {
			return i
		}
		i++
	}
	return -1
}

// @lc code=end


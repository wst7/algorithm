/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	var lastWord string = ""
	var p string = ""
	for _, v := range s {
		if v != ' ' {
			p += string(v)
			lastWord = p
		} else {
			p = ""
		}
	}
	return len(lastWord)
}

// @lc code=end


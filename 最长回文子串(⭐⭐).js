// https://leetcode-cn.com/problems/longest-palindromic-substring/


/**
 * 思路：从中间开始向两边扩散来判断回文串
 * @param {string} s
 * @return {string}
 */
 var longestPalindrome = function(s) {
  // 寻找最长回文串
  const palindrome = (s, l, r) => {
      let flag = false
      while(l >= 0 && r < s.length && s[l] == s[r]) {
          flag = true
          l -= 1
          r += 1
      }

      if (flag) {
          l = l + 1
      }

      return s.substring(l, r)
  }
  let res = ''
  for (let i = 0; i < s.length; i ++) {

      // 寻找长度为奇数的回文字串
      const s1 = palindrome(s, i, i)

      // 寻找长度为偶数的回文字串
      const s2 = palindrome(s, i, i + 1)

      res = res.length > s1.length ? res : s1
      res = res.length > s2.length ? res : s2
  }
  return res
};
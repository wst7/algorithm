// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/


/**
 * @param {string} s
 * @return {number}
 */
 var lengthOfLongestSubstring = function(s) {

  const length = s.length
  const occ = new Set()
  
  // 右指针
  let rk = -1
  let ans = 0

  for (let i = 0; i < length; i++) {

      // delete
      if(i !== 0) {
          occ.delete(s.charAt(i - 1))
      }

      while(rk + 1 < length && !occ.has(s.charAt(rk + 1))) {
          // add
          occ.add(s.charAt(rk + 1))
          ++rk
      }

      ans = Math.max(ans, rk - i + 1)
  }

  return ans

};
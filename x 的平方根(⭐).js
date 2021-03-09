// https://leetcode-cn.com/problems/sqrtx/

/**
 * @param {number} x
 * @return {number}
 */
 var mySqrt = function(x) {
  let l = 0
  let r = x
  let ans = 1
  while(l <= r) {
      let mid = parseInt(l + (r - l) / 2)
      if (mid * mid <= x) {
          ans = mid
          l = mid + 1
      } else {
          r = mid - 1
      }
  }
  return ans
};
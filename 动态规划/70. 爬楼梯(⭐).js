// https://leetcode-cn.com/problems/climbing-stairs/

/**
 * 动态规划：
 * f(x) = f(x - 1) + f(x - 2)
 * f(0) = 1
 * f(1) = 1
 * @param {number} n
 * @return {number}
 */
 var climbStairs = function(n) {
    
  if (n === 0 || n === 1) {
      return 1
  }

  let dpx
  let dpx_1 = 1
  let dpx_2 = 1
  for (let i = 2; i <= n; i++) {
      dpx = dpx_1 + dpx_2
      dpx_2 = dpx_1
      dpx_1 = dpx
      
  }

  return dpx
};
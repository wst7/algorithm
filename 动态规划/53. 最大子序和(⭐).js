// https://leetcode-cn.com/problems/maximum-subarray/

/**
 * @param {number[]} nums
 * @return {number}
 */
 var maxSubArray = function(nums) {

  // dp[i]表示以第i位元素结尾的数组的最大子序和
  const dp = []
  dp[0] = nums[0]
  for (let i = 1; i < nums.length; i++) {
     dp[i] = dp[i - 1] <= 0 ? nums[i] : dp[i - 1] + nums[i]
  }

  return Math.max(...dp)
};
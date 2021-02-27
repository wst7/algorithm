/**
 * 剑指 Offer 42. 连续子数组的最大和
 * @param {number[]} nums
 * @return {number}
 */
var maxSubArray = function(nums) {
  // dp[i] 代表以元素 nums[i] 为结尾的连续子数组最大和
  // dp[0] = nums[0]
  const dp = [...nums]
  dp[0] = nums[0]
  for (let i = 1; i < nums.length; i++) {
      // 当 dp[i - 1] > 0 时：执行 dp[i] = dp[i-1] + nums[i] ；
      // 当 dp[i - 1] ≤0 时：执行 dp[i] = nums[i] ；

      dp[i] = nums[i] + Math.max(dp[i - 1], 0)
  }
  return Math.max(...dp)
};
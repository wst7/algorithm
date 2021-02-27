/**
 * @param {number[]} nums
 * @return {number}
 */
var lengthOfLIS = function(nums) {
  // dp[i]表示以nums[i]这个数结尾的最长递增子序列长度，初始化为1
  const dp = new Array(nums.length).fill(1)
  for (let i = 0; i < nums.length; i ++) {
      for (let j = 0; j < i; j++) {
          // dp[i] 的值等于第j个递增子序列长度 + 1(j要满足nums[i] > nums[j])
          // 存在多种情况，取最大值
          if (nums[i] > nums[j]) {
              dp[i] = Math.max(dp[i], dp[j] + 1)
          }
      }
  }

  return Math.max(...dp)
};
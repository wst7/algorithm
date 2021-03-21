// https://leetcode-cn.com/problems/single-number/

/**
 * @param {number[]} nums
 * @return {number}
 * 采用异或运算
 */
 var singleNumber = function(nums) {
  let ans = nums[0]
  for (let i = 1; i < nums.length; i++) {
      ans = ans ^ nums[i]
  }
  return ans
};
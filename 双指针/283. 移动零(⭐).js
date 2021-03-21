// https://leetcode-cn.com/problems/move-zeroes/

/**
 * @param {number[]} nums
 * @return {void} Do not return anything, modify nums in-place instead.
 */
 var moveZeroes = function(nums) {
  const swap = (list, i, j) => {
      const temp = list[i]
      list[i] = list[j]
      list[j] = temp
  }

  let j = 0
  for (let i = 0; i < nums.length; i++) {
      
      if (nums[j] === 0 && nums[i] !== 0) {
          swap(nums, j, i)
      }
      
      if (nums[j] !== 0) {
          j++
      }
  }
  return nums
};
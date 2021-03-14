// https://leetcode-cn.com/problems/3sum/


/**
 * @param {number[]} nums
 * @return {number[][]}
 */
 var threeSum = function(nums) {
  if (nums.length < 3 || nums === null) {
      return []
  }
 

  nums.sort((a, b) => a - b)
  // [-4, -1, -1, 0, 1, 2]
  const n = nums.length
  const result = []
  
  for (let i = 0; i < n; i++) {
      if (nums[i] > 0) {
          return result
      }
      if(nums[i] === nums[i-1]) {
          continue
      }
              

      let l = i + 1
      let r = n - 1
      while(l < r) {
          const res = nums[i] + nums[l] + nums[r]
          if (res === 0) {
              result.push([nums[i], nums[l], nums[r]])
              while (nums[l] === nums[l + 1] && l < r) {
                  l += 1
              }
              while (nums[r] === nums[r - 1] && l < r) {
                  r -= 1
              }
              r -= 1
              l += 1
              
          } else if (res > 0) {
              r -= 1
          } else if (res < 0) {
              l += 1
          }
      }
  }

  return result


};
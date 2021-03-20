// https://leetcode-cn.com/problems/permutations/

/**
 * @param {number[]} nums
 * @return {number[][]}
 */
 var permute = function(nums) {
  const ans = []
  const output = []

  const dfs = (nums, output) => {
      if (output.length === nums.length) {
          ans.push([...output])
          return
      }

      for (let i = 0; i < nums.length; i++) {
          if (output.includes(nums[i])) {
              continue
          }
          output.push(nums[i])
          dfs(nums, output)
          output.pop()
      }
      
  }

  dfs(nums, output)

return ans
  
};
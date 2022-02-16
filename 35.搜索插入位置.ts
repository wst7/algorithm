/*
 * @lc app=leetcode.cn id=35 lang=typescript
 *
 * [35] 搜索插入位置
 */

// @lc code=start
function searchInsert(nums: number[], target: number): number {
  let left: number = 0;
  let right: number = nums.length - 1;

  while (left <= right) {
    let mid = Math.floor((left + right) / 2)
    
    if (nums[mid] === target) {
      return mid
    } else  if(nums[mid] < target){
      left = mid + 1
    } else {
      right = mid -1
    }
  }
  return left

};
// @lc code=end

const nums = [1, 3, 5, 6]
const target = 5
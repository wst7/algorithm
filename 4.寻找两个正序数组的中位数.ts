/*
 * @lc app=leetcode.cn id=4 lang=typescript
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
function findMedianSortedArrays(nums1: number[], nums2: number[]): number {
   const m = nums1.length
   const n = nums2.length
   const nums = [...nums1, ...nums2].sort((a, b) => a - b)
   if ((m+n) %2 === 0) {
    return nums[(m+n)/2] + nums[(m+n)/2-1] / 2
   }
   return nums[(m+n -1)/2]
};
// @lc code=end


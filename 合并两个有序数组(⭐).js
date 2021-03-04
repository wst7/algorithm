// 题目
// https://leetcode-cn.com/problems/merge-sorted-array/


// 解题
/**
 * @param {number[]} nums1
 * @param {number} m
 * @param {number[]} nums2
 * @param {number} n
 * @return {void} Do not return anything, modify nums1 in-place instead.
 */
// 思路：双指针，利用原数组有序，从后开始比较元素，填充到nums1中
var merge = function(nums1, m, nums2, n) {
    let p1 = m - 1
    let p2 = n - 1
    let p = m + n - 1
    while(p1 >= 0 && p2 >= 0) {
        nums1[p --] = nums1[p1] > nums2[p2] ? nums1[p1 --] : nums2[p2 --] 
    }
    // 如果p2 < 0说明nums2中的元素已经全部放入到nums1中，不需要其他操作
    // 如果p2 不小于0，说明nums1中的元素都往后移动了，将nums2中剩余的元素直接拼到nums1前面
    if (p2 >= 0) {
        nums1.splice(0, p + 1, ...nums2.slice(0, p2 + 1))
    }
};
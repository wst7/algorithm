// 题目
// https://leetcode-cn.com/problems/two-sum/



// 解题
/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(nums, target) {
    const map = new Map()
    for(let i = 0; i < nums.length; i++) {
        const newTarget = target - nums[i]
        if (map.has(newTarget)) {
           return [map.get(newTarget), i]
        }
        map.set(nums[i], i)
    }
};
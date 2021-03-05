// 题目
// https://leetcode-cn.com/problems/happy-number/


/**
 * 解题
 * 循环计算，将计算过且不等于1的值存储到map中
 * 再次出现map中存储的值时返回false，  计算值等于1时返回true
 */


/**
 * @param {number} n
 * @return {boolean}
 */
var isHappy = function(n) {
    let numStr = n + ''
    let m = 0
    
    const  map = {
        [n]: true
    }
    while(n !== 1) {
        n = 0
        for (let i = 0; i < numStr.length; i ++) {
          n += Math.pow(Number(numStr[i]), 2)
        }
        if (map[n]) {
            return false
        }
        map[n] = true
        numStr = n + ''
    }
    return true
};
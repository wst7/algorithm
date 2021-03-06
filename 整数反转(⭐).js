// 题目
// https://leetcode-cn.com/problems/reverse-integer/


// 解题

/**
 * @param {number} x
 * @return {number}
 */
var reverse = function(x) {
    let res = 0
    let pop
    while(x !== 0) {
        //pop operation:
        pop =  x % 10
        x = parseInt(x / 10)
        //push operation:
        res = res * 10 + pop
    }
    if (res > Math.pow(2, 31) - 1 || res < Math.pow(-2, 31)) {
        return 0
    }
    return res
};
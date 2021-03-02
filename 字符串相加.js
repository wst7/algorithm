// 题目
// https://leetcode-cn.com/problems/add-strings/

// 解题
/**
 * @param {string} num1
 * @param {string} num2
 * @return {string}
 */
var addStrings = function(num1, num2) {
    let res = []
    let carry = 0 // 进位
    let i = num1.length - 1
    let j = num2.length -1
    while(i >= 0 || j >= 0 || carry) {
        //                             转换成数字
        const x = i >= 0 ? num1.charAt(i) - '0' : 0
        const y = j >= 0 ? num2.charAt(j) - '0' : 0
        const result = x + y + carry
        
        res.unshift(result % 10)
        // 向下取整
        carry = Math.floor(result / 10)
        i --
        j --
    }
    return res.join('')
};
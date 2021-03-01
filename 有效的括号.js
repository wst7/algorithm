// 题目
// https://leetcode-cn.com/problems/valid-parentheses/
// 解题
/**
 * @param {string} s
 * @return {boolean}
 */
let s = ']'
var isValid = function(s) {
    // 1、replace
    // while(s.includes('{}') || s.includes('()') || s.includes('[]')) {
    //     s = s.replace('{}', '')
    //     s = s.replace('[]', '')
    //     s = s.replace('()', '')
    // }
    // return s === ''
  // 2、stack
    if (s.length % 2 !==0) {
        return false
    }
    const items = []
    for (let i = 0; i < s.length; i ++) {
        const letter = items[items.length - 1]
        switch(s[i]) {
            case '{':
                items.push('{')
                break
            case '(':
                items.push('(')
                break
            case '[':
                items.push('[')
                break
            case '}':
                if (letter !== '{' ) {
                    return false
                }
                items.pop()
                break
            case ']':
                if (letter !== '[' ) {
                    return false
                    
                }
                items.pop()
                break
            case ')':
                if (letter !== '(' ) {
                    return false
                    
                }
                items.pop()
                break
        }
    }
    return items.length === 0
};
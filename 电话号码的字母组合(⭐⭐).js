
// https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/


/**
 * @param {string} digits
 * @return {string[]}
 */
 var letterCombinations = function(digits) {
  if(digits === '') {
      return []
  }
  let ans = 0
  const letterMap = {
      2: ['a', 'b', 'c'],
      3: ['d', 'e', 'f'],
      4: ['g', 'h', 'i'],
      5: ['j', 'k', 'l'],
      6: ['m', 'n', 'o'],
      7: ['p', 'q', 'r', 's'],
      8: ['t', 'u', 'v'],
      9: ['w', 'x', 'y', 'z'],
  }
  const result = []
  const stack = [] // 每一种可能
  
  const dfs = (stack, depth) => {

      if (depth === digits.length) { // 出口
          result.push([...stack])
          return
      }
      for (let i = 0; i < letterMap[digits[depth]].length; i++) {
          stack.push(letterMap[digits[depth]][i]) // 先push进去
          dfs(stack, depth + 1) // 遍历完子元素
          stack.pop() // 在pop出来，进行下轮
      }
  }

  dfs(stack, 0)

  return result.map(item => item.join(''))
};
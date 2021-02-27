/**
 * @param {string} s
 * @return {string}
 */
var decodeString = function(s) {
  let multi  = 0 
  let res = ''
  const stack = []

  for (let c of s) {
      if (c === '[') {
          stack.push([multi, res])
          res = '', multi = 0
      } else if (c === ']') {
          const pop = stack.pop()
          const cur_mult = pop[0]
          const last_res = pop[1]
          //    之前结果       // 
          res = last_res + new Array(cur_mult).fill(res).join('')
      } else if (c >= '0' && c <= '9') {
          multi = multi * 10 + Number(c)
      } else {
          // 记录重复单元
          res += c
      }
  }

  return res
  
};
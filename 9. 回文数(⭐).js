// https://leetcode-cn.com/problems/palindrome-number/

/**
 * @param {number} x
 * @return {boolean}
 */
 var isPalindrome = function(x) {
  if (x < 0 || (x % 10 === 0 && x !== 0)) {
      return false
  }

  let reverted = 0
  const ref =  x

  while(x > 0) {
      reverted = reverted * 10 + x % 10
      x = Math.floor(x / 10)
  }

  return ref === reverted
};
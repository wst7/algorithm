// https://leetcode-cn.com/problems/palindrome-linked-list/


/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @return {boolean}
 */
 var isPalindrome = function(head) {
  // 思路：将val放到数组中，使用双指针进行判断
  const values = []
  while(head !== null)  {
      values.push(head.val)
      head = head.next
  }

  for (let i = 0, j = values.length - 1; i < j; i++, j--) {
      if (values[i] !== values[j]) {
          return false
      }
  }
  return true
};
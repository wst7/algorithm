// https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/

/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} head
 * @param {number} k
 * @return {ListNode}
 */
 var getKthFromEnd = function(head, k) {
  // 第一个指针记录链表长度n
  let front = head
  // 第二个指针记录第(n - k)个指针
  let back = head

  // 先把第一个指针移动k步
  while(k > 0) {
      front = front.next
      k--
  }

  // 同时遍历移动front和back，直到front到结尾，back就是倒数第k个


  while(front) {
      front = front.next
      back = back.next
  }

  return back

};
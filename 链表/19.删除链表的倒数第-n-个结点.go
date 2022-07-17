/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0,head}
	first, second := head, dummy

	for i := 0; i < n; i++ {
		first = first.Next
	}

	for first != nil { // for ;first != nil; { 
		first = first.Next
		second = second.Next
	}

	second.Next = second.Next.Next
	return dummy.Next
}
// @lc code=end

/** TS
* function removeNthFromEnd(head: ListNode | null, n: number): ListNode | null {
  let first = head
  const dummy = new ListNode(0, head)
  let second = dummy
 
  for (let i = 0; i < n; i++) {
    first = first.next
  }
  while (first !== null) {
    first = first.next
    second = second.next
  }

  second.next = second.next.next
// 返回head报错？
  return head
};
/
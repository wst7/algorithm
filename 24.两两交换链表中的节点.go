/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	h := &ListNode{-1, head}
	pre := h
	for pre.Next != nil && pre.Next.Next != nil {
		node1 := pre.Next
		node2 := node1.Next
		lst := node2.Next

		pre.Next = node2
		node2.Next = node1
		node1.Next = lst

		pre = node1
	}

	return h.Next

}

// @lc code=end
// 我们首先需要建立pre、node1、node2和lat四个指针即可。
// 然后pre->next=node2;node2.next=node1;node1.next=lat
/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] 环形链表
 */

/*
 * 思路：双指针，快指针fast，慢指针slow
 * 让fast优先一步出发，没次走两步，slow每次走一步，如果有环，
 * 那么在循环几次后，fast必然追上slow
 * 如果fast.Next 不存在，则说明没有环
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

// @lc code=end

// Testcase
// [3,2,0,-4]
// 1
// 3,2  2,-4 0,0
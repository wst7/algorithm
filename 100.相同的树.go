/*
 * @lc app=leetcode.cn id=100 lang=golang
 *
 * [100] 相同的树
 *
 * https://leetcode.cn/problems/same-tree/description/
 *
 * algorithms
 * Easy (60.94%)
 * Likes:    1139
 * Dislikes: 0
 * Total Accepted:    542.3K
 * Total Submissions: 889.2K
 * Testcase Example:  '[1,2,3]\n[1,2,3]'
 *
 * 给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。
 *
 * 如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：p = [1,2,3], q = [1,2,3]
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：p = [1,2], q = [1,null,2]
 * 输出：false
 *
 *
 * 示例 3：
 *
 *
 * 输入：p = [1,2,1], q = [1,1,2]
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 两棵树上的节点数目都在范围 [0, 100] 内
 * -10^4
 *
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 深度优先迭代法
// func isSameTree(p *TreeNode, q *TreeNode) bool {
// 	// 都是nil
// 	if p == nil && q == nil {
// 		return true
// 	}
// 	// 其中一个是nil
// 	if p == nil || q == nil {
// 		return false
// 	}
// 	// 当前节点的Val不相等
// 	if p.Val != q.Val {
// 		return false
// 	}
// 	// 左子树和右子树是否都相等
// 	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
// }

// 广度优先队列法
func isSameTree(p *TreeNode, q *TreeNode) bool {
	queue := []*TreeNode{p, q}
	for len(queue) > 0 {
		p, q = queue[0], queue[1]
		queue = queue[2:]
		if p == nil && q == nil {
			continue
		}
		if p == nil || q == nil {
			return false
		}
		if p.Val != q.Val {
			return false
		}
		queue = append(queue, p.Left, q.Left)
		queue = append(queue, p.Right, q.Right)
	}
	return true
}

// @lc code=end


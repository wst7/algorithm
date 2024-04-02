/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
 *
 * https://leetcode.cn/problems/binary-tree-postorder-traversal/description/
 *
 * algorithms
 * Easy (76.57%)
 * Likes:    1166
 * Dislikes: 0
 * Total Accepted:    753.2K
 * Total Submissions: 983.5K
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给你一棵二叉树的根节点 root ，返回其节点值的 后序遍历 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,null,2,3]
 * 输出：[3,2,1]
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = [1]
 * 输出：[1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点的数目在范围 [0, 100] 内
 * -100 <= Node.val <= 100
 *
 *
 *
 *
 * 进阶：递归算法很简单，你可以通过迭代算法完成吗？
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
// 递归
// func postorderTraversal(root *TreeNode) []int {
// 	res := []int{}
// 	var dfs func(root *TreeNode)
// 	dfs = func(root *TreeNode) {
// 		if root == nil {
// 			return
// 		}
// 		dfs(root.Left)
// 		dfs(root.Right)
// 		res = append(res, root.Val)
// 	}
// 	dfs(root)
// 	return res
// }

// 迭代: 左右根
// 前序：根左右，交换左右顺序-> 根右左，反转 -> 左右根（后序）
func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	stack := []*TreeNode{root}

	for len(stack) > 0 {

		// pop
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		res = append([]int{node.Val}, res...)

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	return res
}

// @lc code=end


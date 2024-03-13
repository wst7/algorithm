/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 *
 * https://leetcode.cn/problems/binary-tree-inorder-traversal/description/
 *
 * algorithms
 * Easy (76.60%)
 * Likes:    2048
 * Dislikes: 0
 * Total Accepted:    1.4M
 * Total Submissions: 1.8M
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,null,2,3]
 * 输出：[1,3,2]
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
 * 树中节点数目在范围 [0, 100] 内
 * -100 <= Node.val <= 100
 *
 *
 *
 *
 * 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
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
// func inorderTraversal(root *TreeNode) []int {
// 	list := make([]int, 0)
// 	var dfs func(*TreeNode)
// 	dfs = func(root *TreeNode) {
// 		if root == nil {
// 			return
// 		}
// 		dfs(root.Left)
// 		list = append(list, root.Val)
// 		dfs(root.Right)
// 	}
// 	dfs(root)
// 	return list
// }

// 使用迭代实现
func inorderTraversal(root *TreeNode) []int {
	list := make([]int, 0)
	stack := make([]*TreeNode, 0)
	p := root

	for p != nil || len(stack) > 0 {
		for p != nil {
			// 将所有节点的左节点全部推入栈中
			stack = append(stack, p)
			p = p.Left
		}
		// 逐一访问每个节点及其右节点
		// pop
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		list = append(list, node.Val)
		p = node.Right
	}

	return list
}

// @lc code=end

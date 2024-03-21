/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
 *
 * https://leetcode.cn/problems/balanced-binary-tree/description/
 *
 * algorithms
 * Easy (58.16%)
 * Likes:    1495
 * Dislikes: 0
 * Total Accepted:    591K
 * Total Submissions: 1M
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，判断它是否是 平衡二叉树
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [3,9,20,null,null,15,7]
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [1,2,2,3,3,null,null,4,4]
 * 输出：false
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = []
 * 输出：true
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中的节点数在范围 [0, 5000] 内
 * -10^4 <= Node.val <= 10^4
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

// 平衡二叉树: 一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var height func(root *TreeNode) int
	height = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftHeight := height(root.Left) // 左
		if leftHeight == -1 {
			return -1
		}
		rightHeight := height(root.Right) // 右
		if rightHeight == -1 {
			return -1
		}
		if abs(leftHeight-rightHeight) > 1 { // 中
			return -1
		}
		return max(leftHeight, rightHeight) + 1
	}
	return !(height(root) == -1)
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

// @lc code=end


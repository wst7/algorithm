/*
 * @lc app=leetcode.cn id=1382 lang=golang
 *
 * [1382] 将二叉搜索树变平衡
 *
 * https://leetcode.cn/problems/balance-a-binary-search-tree/description/
 *
 * algorithms
 * Medium (73.98%)
 * Likes:    201
 * Dislikes: 0
 * Total Accepted:    27.9K
 * Total Submissions: 37.7K
 * Testcase Example:  '[1,null,2,null,3,null,4]'
 *
 * 给你一棵二叉搜索树，请你返回一棵 平衡后 的二叉搜索树，新生成的树应该与原来的树有着相同的节点值。如果有多种构造方法，请你返回任意一种。
 *
 * 如果一棵二叉搜索树中，每个节点的两棵子树高度差不超过 1 ，我们就称这棵二叉搜索树是 平衡的 。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：root = [1,null,2,null,3,null,4,null,null]
 * 输出：[2,1,3,null,null,null,4]
 * 解释：这不是唯一的正确答案，[3,1,4,null,2,null,null] 也是一个可行的构造方案。
 *
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入: root = [2,1,3]
 * 输出: [2,1,3]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树节点的数目在 [1, 10^4] 范围内。
 * 1 <= Node.val <= 10^5
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
func balanceBST(root *TreeNode) *TreeNode {
	// 中序结果
	list := []int{}
	p := root
	stack := []*TreeNode{}
	var inOrder func(root *TreeNode)
	inOrder = func(root *TreeNode) {
		for p != nil || len(stack) > 0 {
			for p != nil {
				// 所有左节点入栈
				stack = append(stack, p)
				p = p.Left
			}
			// pop
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 访问
			list = append(list, node.Val)
			p = node.Right
		}
	}
	inOrder(root)

	// 将list转换成AVL
	var buildAVL func(start, end int) *TreeNode
	buildAVL = func(start, end int) *TreeNode {
		if start > end {
			return nil
		}
		mid := (start + end) >> 1
		root := &TreeNode{Val: list[mid]}
		root.Left = buildAVL(start, mid-1)
		root.Right = buildAVL(mid+1, end)
		return root
	}
	return buildAVL(0, len(list)-1)
}

// @lc code=end


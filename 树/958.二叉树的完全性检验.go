/*
 * @lc app=leetcode.cn id=958 lang=golang
 *
 * [958] 二叉树的完全性检验
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

type PNode struct {
	Node *TreeNode
	Num  int
}

func isCompleteTree(root *TreeNode) bool {
	nodes := []PNode{PNode{root, 1}}
	i := 0
	for i < len(nodes) {
		iNode := nodes[i]
		i++

		if iNode.Node != nil {
			node := iNode.Node
			num := iNode.Num
			nodes = append(nodes, PNode{node.Left, num * 2})
			nodes = append(nodes, PNode{node.Right, num*2 + 1})
		}
	}

	return nodes[i-1].Num == len(nodes)

}

// @lc code=end


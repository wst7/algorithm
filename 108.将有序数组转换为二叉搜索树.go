/*
 * @lc app=leetcode.cn id=108 lang=golang
 *
 * [108] 将有序数组转换为二叉搜索树
 *
 * https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/description/
 *
 * algorithms
 * Easy (78.14%)
 * Likes:    1486
 * Dislikes: 0
 * Total Accepted:    444K
 * Total Submissions: 567.7K
 * Testcase Example:  '[-10,-3,0,5,9]'
 *
 * 给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 平衡 二叉搜索树。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [-10,-3,0,5,9]
 * 输出：[0,-3,9,-10,null,5]
 * 解释：[0,-10,5,null,-3,null,9] 也将被视为正确答案：
 *
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,3]
 * 输出：[3,1]
 * 解释：[1,null,3] 和 [3,1] 都是高度平衡二叉搜索树。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^4
 * -10^4 <= nums[i] <= 10^4
 * nums 按 严格递增 顺序排列
 *
 *
 */

/*
 二叉搜索树（Binary Sort Tree）：
 1.若他的左子树不为空，则左子树上所有节点的值都小于根节点的值
 2.若它的右子树不为空，则右子树上所有节点的值都大于根节点的值
 3.它的左右子树也分别是二叉搜索树

二叉平衡搜索树（AVL）：
1.左子树与右子树高度之差的绝对值不超过1
2.树的每个左子树和右子树都是AVL树
3.每一个节点都有一个平衡因子（balance factor），任一节点的平衡因子是-1、0、1（每一个节点的平衡因子 = 右子树高度 - 左子树高度）
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

/**
* 思路:
* [1382] 将二叉搜索树变平衡的 的子问题
 */
func sortedArrayToBST(nums []int) *TreeNode {
	var buildAVL func(start, end int) *TreeNode
	buildAVL = func(start, end int) *TreeNode {
		if start > end {
			return nil
		}
		mid := (start + end) >> 1
		root := &TreeNode{Val: nums[mid]}
		root.Left = buildAVL(start, mid-1)
		root.Right = buildAVL(mid+1, end)
		return root
	}
	return buildAVL(0, len(nums)-1)
}

// @lc code=end


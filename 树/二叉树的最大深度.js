/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number}
 */
var maxDepth = function(root) {
  if (root === null) 0
  let depth = 1
  const dfs = (root, level = 0) => {
      if (root === null) return
      depth = Math.max(depth, level + 1)
      dfs(root.left, level + 1)
      dfs(root.right, level + 1)
  }
  dfs(root)
  return depth
};
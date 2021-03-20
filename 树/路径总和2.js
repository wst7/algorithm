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
 * @param {number} targetSum
 * @return {number[][]}
 */
var pathSum = function(root, targetSum) {
  if (root === null) {
      return []
  }
  const result = []
  const path = []

  const dfs = (root, targetSum) => {
      if (root === null) return
      path.push(root.val)
      targetSum -= root.val
      if (!root.left && !root.right && targetSum === 0) { // 叶子节点
          result.push([...path])
      }
      

      dfs(root.left, targetSum)
      dfs(root.right, targetSum)

      path.pop()
  }
  dfs(root, targetSum)

  return result
  
};
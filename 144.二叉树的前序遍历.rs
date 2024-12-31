/*
 * @lc app=leetcode.cn id=144 lang=rust
 *
 * [144] 二叉树的前序遍历
 */

// @lc code=start
// Definition for a binary tree node.
// #[derive(Debug, PartialEq, Eq)]
// pub struct TreeNode {
//   pub val: i32,
//   pub left: Option<Rc<RefCell<TreeNode>>>,
//   pub right: Option<Rc<RefCell<TreeNode>>>,
// }
// 
// impl TreeNode {
//   #[inline]
//   pub fn new(val: i32) -> Self {
//     TreeNode {
//       val,
//       left: None,
//       right: None
//     }
//   }
// }
use std::rc::Rc;
use std::cell::RefCell;
impl Solution {
    pub fn preorder_traversal(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
       if root.is_none() {
            return vec![];
        }
        let mut res = vec![];
        let root = root.unwrap();
        res.push(root.borrow().val);
        if root.borrow().left.is_some() {
            res.append(&mut Solution::preorder_traversal(root.borrow().left.clone()));
        }
        if root.borrow().right.is_some() {
            res.append(&mut Solution::preorder_traversal(root.borrow().right.clone()));
        }
        res 
    }
}
// @lc code=end


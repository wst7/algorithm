/*
 * @lc app=leetcode.cn id=66 lang=rust
 *
 * [66] 加一
 */

// @lc code=start
impl Solution {
    pub fn plus_one(digits: Vec<i32>) -> Vec<i32> {
       
        let mut result = digits.clone();
         let mut pos = result.len() - 1;
        loop {
            // 如果动态数组最后一个数小于 9，则直接加 1 并返回
            if (result[pos] < 9) {
                result[pos] += 1;
                return result;
            }
            // 否则将当前位置重置为0，高一位数进去循环
            result[pos] = 0;
            if (pos > 0) {
                pos -= 1;
            } else if (pos == 0) {
                break;
            }
        }
        // 重置数组，数组长度因进位而加 1，除第一个元素为 1 外，其余元素为 0
        result = vec![0; result.len() + 1];
        result[0] = 1;
        result
    }
}
// @lc code=end


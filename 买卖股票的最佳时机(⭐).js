// 题目
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/


// 解题
/**
 * @param {number[]} prices
 * @return {number}
 */
var maxProfit = function(prices) {
    let maxprofit = 0
    // 前i天最小价
    let minprice = prices[0]
    for (let i = 1; i < prices.length; i++) {
        
        maxprofit = Math.max(maxprofit, prices[i] - minprice)
        minprice = Math.min(minprice, prices[i])
    }
    return maxprofit
};
// 题目

// https://leetcode-cn.com/problems/last-stone-weight/


// 解题


/**
 * @param {number[]} stones
 * @return {number}
 */
 var lastStoneWeight = function(stones) {

    
  if (stones.length === 0) {
      return 0
  }
  if (stones.length === 1) {
      return stones[0]
  }

  stones.sort((a, b) => a - b)

  const y = stones[stones.length - 1]
  const x = stones[stones.length - 2]


  if (x === y) {
      stones.length = stones.length - 2
  } else {
      stones.length = stones.length - 1
      stones[stones.length - 1] = y - x

  }

  return lastStoneWeight(stones)
};

// https://leetcode-cn.com/problems/compare-version-numbers/

/**
 * @param {string} version1
 * @param {string} version2
 * @return {number}
 */
 var compareVersion = function(version1, version2) {
  const version1List = version1.split('.')
  const version2List = version2.split('.')
  const length1 = version1List.length
  const length2 = version2List.length

  const maxLength = Math.max(length1, length2)
  for(let i = 0; i < maxLength; i++) {
      let i1 = i < length1 ? Number(version1List[i]) : 0
      let i2 = i < length2 ? Number(version2List[i]) : 0

      if (i1 !== i2) {
          return i1 > i2 ? 1 : -1
      }
  }

  return 0
};
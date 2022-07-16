// 通用函数
function swap(array, left, right) {
    let rightValue = array[right]
    array[right] = array[left]
    array[left] = rightValue
}

/* 1冒泡排序
*  原理：
*  从第一个元素开始，把当前元素current与下一个索引对应的元素next进行比较，如果current > next（升序），
*  则交换位置，重复此操作直到最后一个元素，此时最后一个元素就是这个数组中最大的元素，下一轮比较中已经确定顺序的
*  元素就不需要参与了
*  时间复杂度：平方阶O(n^2)
*/
function bubbleSort(array) {
    for (let i = array.length - 1; i > 0; i--) {
        for (let j = 0; j < i; j++) {
            if (array[j] > array[j + 1]) {
                swap(array, j, j + 1)
            }
        }
        
    }
}


/* 2插入排序
*  原理：
*  默认第一个元素已经排好序了，取下一个元素和当前元素作比较，如果当前元素大（升序）就交换位置，
*  此时前两个元素已经排好序了，再取下一个元素和前面已经排好序的元素在进行比较(从后往前比较， 
*  因为下一个元素是和前面已经排好序的的最后一个元素是连着的)
*  时间复杂度：平方阶O(n^2)
*/

function insertSort(array) {
    for (let i = 1; i < array.length; i++) {
        for (let j = i - 1; j >= 0; j--) {
           if (array[j] > array[j + 1]) {
              swap(array, j, j + 1)
           }
        }
    }
}


/* 3选择排序
    （选择最小的）
*  原理：
*  从索引值0开始，与后面元素依次进行比较，如果索引值0的元素比某个元素大，则交换位置（升序），
*  第一轮比较完之后，索引值0位置的元素就是最小的，下一轮取索引值是1的元素重复操作，寻找下一
*  个相对比较小的值
*  
*  时间复杂度：平方阶O(n^2)
*/

function selectionSort(array) {
    for (let i = 0; i < array.length - 1; i++) {
        for (let j = i + 1; j < array.length; j++) {
           if (array[i] > array[j]) {
              swap(array, i, j)
           }
        }
    }
}

/**
 * 4快速排序
 * 原理：
 * （1）在数据集之中，选择一个元素作为"基准"（pivot）。
 * （2）所有小于"基准"的元素，都移到"基准"的左边；所有大于"基准"的元素，都移到"基准"的右边。
 * （3）对"基准"左边和右边的两个子集，不断重复第一步和第二步，直到所有子集只剩下一个元素为止。
 *  
 * 时间复杂度：平均时间复杂度O(nlogn)  最差复杂度O( n^2 )
 */
function quickSort(array) {
    // 递归只剩一个元素的时候直接返回
    if (array.length <= 1) { 
        return array; 
    }

    const pivotIndex = Math.floor(array.length / 2)
    const pivot = array.splice(pivotIndex, 1)[0];
    const left = [];
    const right = [];
  
    for (let i = 0; i < array.length; i++) {
　　　　if (array[i] < pivot) {
　　　　　　left.push(array[i]);
　　　　} else {
　　　　　　right.push(array[i]);
　　　　}
　　}

    return [...quickSort(left), pivot, ...quickSort(right)]
}
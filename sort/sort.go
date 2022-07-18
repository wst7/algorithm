package main

import (
	"fmt"
	"math"
)

func swap(arr []int, left int, right int) {
	rightV := arr[right]
	arr[right] = arr[left]
	arr[left] = rightV
}

/*
 * 冒泡排序
 * i代表最后停放元素的位置
 * j和j+1代表比较的元素的下标，蒋遇到的最大的元素移动到i位置，移动过程中，最大元素会变化，有点像🐒掰苞米
 */

func bubbleSort(arr []int) {
	n := len(arr)
	for i := n - 1; i > 0; i-- { //
		fmt.Println("第", n-i, "次冒泡")
		for j := 0; j < i; j++ { //
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
			}
		}
		fmt.Println(arr)
	}
	fmt.Println("结果：", arr)
}

// 2.插入排序
// i指向待排序的元素
// j遍历已排序的元素，寻找第i个元素的位置
func insertSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		fmt.Println("插入第", i, "个位置")
		for j := 0; j < i; j++ {
			if arr[i] < arr[j] {
				swap(arr, i, j)
			}
		}
		fmt.Println(arr)
	}
	fmt.Println("结果：", arr)
}

// 3.选择排序
// 选择要放置元素的位置，将其与其他元素进行比较大小，判断是否要交换
func selectSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		fmt.Println("第", i, "个位置选择排序")
		for j := i + 1; j < n; j++ {
			if arr[i] > arr[j] {
				swap(arr, i, j)
			}
		}
		fmt.Println(arr)
	}
	fmt.Println("结果：", arr)
}

// 4.快速排序
// 选择基准元素，把小于基准元素的都放在基准元素左边，大于基准元素的都放在基准元素右边
// 左右两边的子集重复上述动作

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pIndex := int(math.Floor(float64(len(arr) / 2)))
	pItem := arr[pIndex]
	arr = remove(arr, pIndex)
	left, right, result := []int{}, []int{}, []int{}
	for i := 0; i < len(arr); i++ {
		if arr[i] < pItem {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	result = append(result, quickSort(left)...)
	result = append(result, pItem)
	result = append(result, quickSort(right)...)
	fmt.Println(result)
	return result
}

func main() {
	var arr = []int{3, 4, 2, 5, 1, 9, 8, 7}
	// bubbleSort(arr)
	// insertSort(arr)
	// selectSort(arr)
	res := quickSort(arr)
	fmt.Println(res)
}

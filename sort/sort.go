package main

import "fmt"

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

func main() {
	var arr = []int{3, 4, 2, 5, 1}
	// bubbleSort(arr)
	insertSort(arr)
}

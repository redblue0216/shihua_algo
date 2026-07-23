package main

import (
	"fmt"
	"shihua_algo/sorting"
)

func main() {
	// 测试选择排序
	fmt.Println("===== 选择排序测试 =====")
	selectNums := []int{4, 1, 3, 2, 7, 5}
	fmt.Printf("排序前: %v\n", selectNums)
	sorting.SelectionSort(selectNums)
	fmt.Printf("排序后: %v\n\n", selectNums)

	// 测试基础冒泡排序
	fmt.Println("===== 基础冒泡排序测试 =====")
	bubbleNums := []int{9, 2, 5, 1, 6, 3}
	fmt.Printf("排序前: %v\n", bubbleNums)
	sorting.BubbleSort(bubbleNums)
	fmt.Printf("排序后: %v\n\n", bubbleNums)

	// 测试带优化标志的冒泡排序
	fmt.Println("===== 带标志优化冒泡排序测试 =====")
	bubbleFlagNums := []int{8, 4, 7, 3, 1, 2}
	fmt.Printf("排序前: %v\n", bubbleFlagNums)
	sorting.BubbleSortWithFlag(bubbleFlagNums)
	fmt.Printf("排序后: %v\n\n", bubbleFlagNums)

	// 测试插入排序
	fmt.Println("===== 插入排序测试 =====")
	insertNums := []int{5, 7, 2, 9, 1, 4}
	fmt.Printf("排序前: %v\n", insertNums)
	sorting.InsertionSort(insertNums)
	fmt.Printf("排序后: %v\n", insertNums)
}
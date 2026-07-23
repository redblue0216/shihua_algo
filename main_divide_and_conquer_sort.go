package main

import (
	"fmt"
	"shihua_algo/sorting"
)

func main() {
	// ===== 已有基础排序（保留方便一次性全跑）=====
	fmt.Println("===== 选择排序测试 =====")
	selectNums := []int{4, 1, 3, 2, 7, 5}
	fmt.Printf("排序前: %v\n", selectNums)
	sorting.SelectionSort(selectNums)
	fmt.Printf("排序后: %v\n\n", selectNums)

	fmt.Println("===== 基础冒泡排序测试 =====")
	bubbleNums := []int{9, 2, 5, 1, 6, 3}
	fmt.Printf("排序前: %v\n", bubbleNums)
	sorting.BubbleSort(bubbleNums)
	fmt.Printf("排序后: %v\n\n", bubbleNums)

	fmt.Println("===== 带标志优化冒泡排序测试 =====")
	bubbleFlagNums := []int{8, 4, 7, 3, 1, 2}
	fmt.Printf("排序前: %v\n", bubbleFlagNums)
	sorting.BubbleSortWithFlag(bubbleFlagNums)
	fmt.Printf("排序后: %v\n\n", bubbleFlagNums)

	fmt.Println("===== 插入排序测试 =====")
	insertNums := []int{5, 7, 2, 9, 1, 4}
	fmt.Printf("排序前: %v\n", insertNums)
	sorting.InsertionSort(insertNums)
	fmt.Printf("排序后: %v\n\n", insertNums)

	// ===== 新增高级排序：快速排序三版本 =====
	fmt.Println("===== 基础快速排序（left基准）测试 =====")
	quickBasic := []int{7, 3, 5, 1, 9, 2, 8}
	fmt.Printf("排序前: %v\n", quickBasic)
	sorting.QuickSort(quickBasic)
	fmt.Printf("排序后: %v\n\n", quickBasic)

	fmt.Println("===== 快速排序（三数取中优化）测试 =====")
	quickMedian := []int{7, 3, 5, 1, 9, 2, 8}
	fmt.Printf("排序前: %v\n", quickMedian)
	sorting.QuickSortMedian(quickMedian)
	fmt.Printf("排序后: %v\n\n", quickMedian)

	fmt.Println("===== 快速排序（递归深度优化）测试 =====")
	quickTail := []int{7, 3, 5, 1, 9, 2, 8}
	fmt.Printf("排序前: %v\n", quickTail)
	sorting.QuickSortTailCall(quickTail)
	fmt.Printf("排序后: %v\n\n", quickTail)

	// ===== 归并排序 =====
	fmt.Println("===== 归并排序测试 =====")
	mergeNums := []int{6, 2, 9, 1, 5, 3}
	fmt.Printf("排序前: %v\n", mergeNums)
	sorting.MergeSort(mergeNums)
	fmt.Printf("排序后: %v\n\n", mergeNums)

	// ===== 堆排序 =====
	fmt.Println("===== 堆排序测试 =====")
	heapNums := []int{10, 4, 6, 3, 8, 2}
	fmt.Printf("排序前: %v\n", heapNums)
	sorting.HeapSort(heapNums)
	fmt.Printf("排序后: %v\n", heapNums)
}
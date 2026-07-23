package main

import (
	"fmt"
	"shihua_algo/sorting"
)

func main() {
	// ===== 基础简单排序 =====
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

	// ===== 分治类排序：快排、归并、堆排序 =====
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

	fmt.Println("===== 归并排序测试 =====")
	mergeNums := []int{6, 2, 9, 1, 5, 3}
	fmt.Printf("排序前: %v\n", mergeNums)
	sorting.MergeSort(mergeNums)
	fmt.Printf("排序后: %v\n\n", mergeNums)

	fmt.Println("===== 堆排序测试 =====")
	heapNums := []int{10, 4, 6, 3, 8, 2}
	fmt.Printf("排序前: %v\n", heapNums)
	sorting.HeapSort(heapNums)
	fmt.Printf("排序后: %v\n\n", heapNums)

	// ===== 新增线性排序：桶排序、计数排序、基数排序 =====
	fmt.Println("===== 桶排序测试（输入必须[0,1)浮点数） =====")
	bucketNums := []float64{0.42, 0.32, 0.33, 0.52, 0.37, 0.47, 0.51}
	fmt.Printf("排序前: %v\n", bucketNums)
	sorting.BucketSort(bucketNums)
	fmt.Printf("排序后: %v\n\n", bucketNums)

	fmt.Println("===== 简易计数排序（不稳定）测试 =====")
	countSimple := []int{1, 0, 4, 1, 3, 2}
	fmt.Printf("排序前: %v\n", countSimple)
	sorting.CountingSortNaive(countSimple)
	fmt.Printf("排序后: %v\n\n", countSimple)

	fmt.Println("===== 完整版稳定计数排序测试 =====")
	countStable := []int{1, 0, 4, 1, 3, 2}
	fmt.Printf("排序前: %v\n", countStable)
	sorting.CountingSort(countStable)
	fmt.Printf("排序后: %v\n\n", countStable)

	fmt.Println("===== 基数排序测试 =====")
	radixNums := []int{23, 12, 45, 9, 345, 89, 187}
	fmt.Printf("排序前: %v\n", radixNums)
	sorting.RadixSort(radixNums)
	fmt.Printf("排序后: %v\n", radixNums)
}
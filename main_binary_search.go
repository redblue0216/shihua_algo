package main

import (
	"fmt"
	"shihua_algo/search"
)

func main() {
	fmt.Println("===== 1. 基础二分查找（双闭区间 BinarySearch）测试 =====")
	// 有序测试数组（含重复元素）
	nums := []int{1, 3, 4, 4, 4, 6, 7, 9, 12, 15}
	target1 := 7
	idx1 := search.BinarySearch(nums, target1)
	fmt.Printf("数组：%v，查找值 %d，索引位置：%d\n", nums, target1, idx1)

	targetNoExist := 8
	idxNo := search.BinarySearch(nums, targetNoExist)
	fmt.Printf("数组：%v，查找不存在值 %d，索引位置：%d\n\n", nums, targetNoExist, idxNo)

	fmt.Println("===== 2. 二分查找（左闭右开 BinarySearchLCRO）测试 =====")
	target2 := 4
	idx2 := search.BinarySearchLCRO(nums, target2)
	fmt.Printf("数组：%v，查找值 %d，左闭右开结果索引：%d\n\n", nums, target2, idx2)

	fmt.Println("===== 3. 插入点算法（无重复 BinarySearchInsertionSimple）测试 =====")
	uniqueNums := []int{2, 5, 8, 11, 14}
	insertTarget := 9
	insertPosSimple := search.BinarySearchInsertionSimple(uniqueNums, insertTarget)
	fmt.Printf("无重复数组：%v，值 %d 插入位置：%d\n\n", uniqueNums, insertTarget, insertPosSimple)

	fmt.Println("===== 4. 插入点算法（支持重复 BinarySearchInsertion）测试 =====")
	dupInsertTarget := 4
	insertPosDup := search.BinarySearchInsertion(nums, dupInsertTarget)
	fmt.Printf("含重复数组：%v，值 %d 左侧插入位置：%d\n\n", nums, dupInsertTarget, insertPosDup)

	fmt.Println("===== 5. 查找左边界 BinarySearchLeftEdge 测试 =====")
	leftTarget := 4
	leftIdx := search.BinarySearchLeftEdge(nums, leftTarget)
	fmt.Printf("数组：%v，值 %d 最左侧索引：%d\n\n", nums, leftTarget, leftIdx)

	fmt.Println("===== 6. 查找右边界 BinarySearchRightEdge 测试 =====")
	rightTarget := 4
	rightIdx := search.BinarySearchRightEdge(nums, rightTarget)
	fmt.Printf("数组：%v，值 %d 最右侧索引：%d\n", nums, rightTarget, rightIdx)
}

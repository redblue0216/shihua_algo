package main

import (
	"fmt"
	"shihua_algo/backtracking"
)

func main() {
	// 测试子集和I：元素可重复选取
	fmt.Println("===== 子集和I（元素可重复选择）测试 =====")
	sum1Nums := []int{2, 3, 5}
	target1 := 8
	fmt.Printf("原始数组: %v，目标和: %d\n", sum1Nums, target1)
	ret1 := backtracking.SubsetSumI(sum1Nums, target1)
	fmt.Printf("满足条件子集: %v\n\n", ret1)

	// 测试子集和II：元素不可重复、数组含重复数字，子集自动去重
	fmt.Println("===== 子集和II（元素不可重复，含重复数字去重）测试 =====")
	sum2Nums := []int{2, 2, 3, 4}
	target2 := 5
	fmt.Printf("原始数组: %v，目标和: %d\n", sum2Nums, target2)
	ret2 := backtracking.SubsetSumII(sum2Nums, target2)
	fmt.Printf("满足条件子集: %v\n", ret2)
}
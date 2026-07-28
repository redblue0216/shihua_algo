package main

import (
	"fmt"
	"shihua_algo/backtracking"
)

func main() {
	// 测试全排列I：无重复数字
	fmt.Println("===== 全排列I（无重复元素）测试 =====")
	perm1Nums := []int{1, 2, 3}
	fmt.Printf("原始数组: %v\n", perm1Nums)
	ret1 := backtracking.PermutationsI(perm1Nums)
	fmt.Printf("全排列结果: %v\n\n", ret1)

	// 测试全排列II：包含重复数字，自动去重
	fmt.Println("===== 全排列II（含重复元素去重）测试 =====")
	perm2Nums := []int{1, 1, 2}
	fmt.Printf("原始数组: %v\n", perm2Nums)
	ret2 := backtracking.PermutationsII(perm2Nums)
	fmt.Printf("去重全排列结果: %v\n", ret2)
}
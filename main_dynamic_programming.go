package main

import (
	"fmt"
	"shihua_algo/dynamic_programming"
)

func main() {
	// ====================== 第一部分：最小路径和全方案测试 ======================
	fmt.Println("==================== 最小路径和 全部解法测试 ====================")
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	m, n := len(grid), len(grid[0])
	fmt.Printf("测试网格：%v\n", grid)

	// 1.暴力DFS
	resDFS := dynamic_programming.MinPathSumDFS(grid, m-1, n-1)
	fmt.Printf("暴力DFS结果: %d\n", resDFS)

	// 2.记忆化DFS
	mem := make([][]int, m)
	for i := range mem {
		mem[i] = make([]int, n)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	resMem := dynamic_programming.MinPathSumDFSMem(grid, mem, m-1, n-1)
	fmt.Printf("记忆化DFS结果: %d\n", resMem)

	// 3.标准DP二维数组
	resDP := dynamic_programming.MinPathSumDP(grid)
	fmt.Printf("二维DP结果: %d\n", resDP)

	// 4.空间压缩DP一维数组
	resComp := dynamic_programming.MinPathSumDPComp(grid)
	fmt.Printf("空间优化DP结果: %d\n\n", resComp)

	// ====================== 第二部分：01背包全方案测试 ======================
	fmt.Println("==================== 01背包 全部解法测试 ====================")
	wgt01 := []int{2, 3, 4, 5}
	val01 := []int{3, 4, 5, 6}
	cap01 := 8
	totalItem := len(wgt01)
	fmt.Printf("物品重量:%v, 物品价值:%v, 背包容量:%d\n", wgt01, val01, cap01)

	// 暴力DFS
	knapDFS := dynamic_programming.KnapsackDFS(wgt01, val01, totalItem, cap01)
	fmt.Printf("01背包暴力DFS最大价值: %d\n", knapDFS)

	// 记忆化DFS
	memKnap := make([][]int, totalItem+1)
	for i := range memKnap {
		memKnap[i] = make([]int, cap01+1)
		for j := range memKnap[i] {
			memKnap[i][j] = -1
		}
	}
	knapMem := dynamic_programming.KnapsackDFSMem(wgt01, val01, memKnap, totalItem, cap01)
	fmt.Printf("01背包记忆化DFS最大价值: %d\n", knapMem)

	// 标准二维DP
	knapDP := dynamic_programming.KnapsackDP(wgt01, val01, cap01)
	fmt.Printf("01背包二维DP最大价值: %d\n", knapDP)

	// 空间优化一维DP
	knapComp := dynamic_programming.KnapsackDPComp(wgt01, val01, cap01)
	fmt.Printf("01背包空间优化DP最大价值: %d\n\n", knapComp)

	// ====================== 第三部分：完全背包全方案测试 ======================
	fmt.Println("==================== 完全背包 全部解法测试 ====================")
	wgtFull := []int{2, 3, 4}
	valFull := []int{3, 4, 5}
	capFull := 8
	fmt.Printf("物品重量:%v, 物品价值:%v, 背包容量:%d\n", wgtFull, valFull, capFull)

	// 标准二维DP
	fullDP := dynamic_programming.UnboundedKnapsackDP(wgtFull, valFull, capFull)
	fmt.Printf("完全背包二维DP最大价值: %d\n", fullDP)

	// 空间优化一维DP
	fullComp := dynamic_programming.UnboundedKnapsackDPComp(wgtFull, valFull, capFull)
	fmt.Printf("完全背包空间优化DP最大价值: %d\n", fullComp)
}
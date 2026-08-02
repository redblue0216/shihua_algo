package main

import (
	"fmt"
	"shihua_algo/greedy"
)

func main() {
	// ====================== 第一部分：分数背包 fractionalKnapsack 测试 ======================
	fmt.Println("==================== 分数背包（贪心）测试 ====================")
	wgtFrac := []int{2, 3, 4, 5}
	valFrac := []int{3, 4, 5, 6}
	capFrac := 8
	fmt.Printf("物品重量: %v\n物品价值: %v\n背包总容量: %d\n", wgtFrac, valFrac, capFrac)

	maxValFrac := greedy.FractionalKnapsack(wgtFrac, valFrac, capFrac)
	fmt.Printf("分数背包可装入最大总价值: %.2f\n\n", maxValFrac)

	// 补充另一组测试用例，验证部分取物品逻辑
	wgtFrac2 := []int{10, 20, 30}
	valFrac2 := []int{60, 100, 120}
	capFrac2 := 50
	fmt.Printf("【额外用例】重量%v，价值%v，容量%d\n", wgtFrac2, valFrac2, capFrac2)
	maxValFrac2 := greedy.FractionalKnapsack(wgtFrac2, valFrac2, capFrac2)
	fmt.Printf("额外用例分数背包最大价值: %.2f\n\n", maxValFrac2)

	// ====================== 第二部分：盛最多水容器 maxCapacity 测试 ======================
	fmt.Println("==================== 盛水最大容量（双指针贪心）测试 ====================")
	heightArr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Printf("挡板高度数组: %v\n", heightArr)

	maxArea := greedy.MaxCapacity(heightArr)
	fmt.Printf("可以容纳的最大水量: %d\n\n", maxArea)

	// 极简边界用例
	heightEdge := []int{1, 2}
	fmt.Printf("边界高度数组%v，最大容量：%d\n", heightEdge, greedy.MaxCapacity(heightEdge))
}
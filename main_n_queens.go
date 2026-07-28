package main

import (
	"fmt"
	"shihua_algo/backtracking"
)

func main() {
	// 测试4皇后问题
	fmt.Println("===== 4皇后问题测试 =====")
	n4 := 4
	result4 := backtracking.NQueens(n4)
	fmt.Printf("棋盘阶数：%d，总合法方案数：%d\n", n4, len(result4))
	for idx, board := range result4 {
		fmt.Printf("方案 %d:\n", idx+1)
		for _, line := range board {
			fmt.Println(line)
		}
		fmt.Println()
	}

	// 测试8皇后问题，仅打印方案数量（方案过多不完整打印棋盘）
	fmt.Println("===== 8皇后问题测试 =====")
	n8 := 8
	result8 := backtracking.NQueens(n8)
	fmt.Printf("棋盘阶数：%d，总合法方案数：%d\n", n8, len(result8))
}
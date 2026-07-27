package main

import (
	"container/list"
	"fmt"
	"shihua_algo/divide_and_conquer"
)

// 辅助打印list，方便查看柱子圆盘状态
func printList(name string, l *list.List) {
	fmt.Printf("%s: [", name)
	elem := l.Front()
	for elem != nil {
		fmt.Printf("%d ", elem.Value)
		elem = elem.Next()
	}
	fmt.Println("]")
}

func main() {
	fmt.Println("===== 汉诺塔分治算法测试 =====")
	// 初始化三根柱子 A源柱、B缓冲柱、C目标柱
	A := list.New()
	B := list.New()
	C := list.New()

	// 初始化4个圆盘，大的放最底部，数字越大圆盘越大
	diskCount := 4
	for i := diskCount; i >= 1; i-- {
		A.PushBack(i)
	}

	fmt.Println("移动前三根柱子状态：")
	printList("A(源)", A)
	printList("B(缓冲)", B)
	printList("C(目标)", C)

	// 调用分治包暴露的汉诺塔函数
	divide_and_conquer.SolveHanota(A, B, C)

	fmt.Println("\n移动完成后三根柱子状态：")
	printList("A(源)", A)
	printList("B(缓冲)", B)
	printList("C(目标)", C)
}
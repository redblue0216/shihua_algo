package main

import (
	"fmt"
	"shihua_algo/heap"
)

func main() {
	fmt.Println("===== 1. 基于空堆初始化大顶堆 =====")
	// 创建空大顶堆
	h := heap.NewHeap()
	// 元素入堆
	h.Push(1)
	h.Push(3)
	h.Push(2)
	h.Push(5)
	h.Push(4)
	fmt.Println("入堆 1,3,2,5,4 后堆结构：")
	h.Print()

	fmt.Println("\n===== 2. 访问堆顶元素 Peek =====")
	top := h.Peek()
	fmt.Printf("当前堆顶：%d\n", top.(int))

	fmt.Println("\n===== 3. 出堆 Pop 操作 =====")
	popVal := h.Pop()
	fmt.Printf("出堆元素：%d\n", popVal.(int))
	fmt.Println("出堆后堆结构：")
	h.Print()

	fmt.Println("\n===== 4. 获取堆大小 Size =====")
	fmt.Printf("当前堆元素数量：%d\n", h.Size())

	fmt.Println("\n===== 5. 判断堆是否为空 IsEmpty =====")
	fmt.Printf("堆是否为空：%t\n", h.IsEmpty())

	fmt.Println("\n===== 6. 基于已有切片直接建堆 NewMaxHeap =====")
	nums := []any{7, 1, 9, 2, 6}
	h2 := heap.NewMaxHeap(nums)
	fmt.Println("切片 [7,1,9,2,6] 直接建堆结果：")
	h2.Print()

	fmt.Println("\n===== 7. 循环出堆清空堆 =====")
	for !h2.IsEmpty() {
		val := h2.Pop()
		fmt.Printf("弹出：%d ", val.(int))
	}
	fmt.Printf("\n清空后堆是否为空：%t\n", h2.IsEmpty())
}

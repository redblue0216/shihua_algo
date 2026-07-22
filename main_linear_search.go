// main_linear_search.go
package main

import (
	"fmt"
	"shihua_algo/search"
)

func main() {
	fmt.Println("===== 1. 线性查找：数组测试 =====")
	arr := []int{5, 2, 9, 1, 5, 6}
	targetExist := 9
	idx := search.LinearSearchArray(arr, targetExist)
	fmt.Printf("数组：%v\n查找 %d，索引 = %d\n", arr, targetExist, idx)

	targetNotExist := 99
	idxNone := search.LinearSearchArray(arr, targetNotExist)
	fmt.Printf("查找不存在值 %d，索引 = %d\n\n", targetNotExist, idxNone)

	fmt.Println("===== 2. 线性查找：链表测试 =====")
	// 构建链表 10 -> 20 -> 30 -> 40
	n0 := search.CreateListNode(10)
	n1 := search.CreateListNode(20)
	n2 := search.CreateListNode(30)
	n3 := search.CreateListNode(40)
	n0.Next = n1
	n1.Next = n2
	n2.Next = n3

	findVal := 30
	resultNode := search.LinearSearchLinkedList(n0, findVal)
	if resultNode != nil {
		fmt.Printf("找到值 %d，节点Val = %d\n", findVal, resultNode.Val)
	} else {
		fmt.Printf("未找到 %d\n", findVal)
	}

	nilVal := 99
	emptyNode := search.LinearSearchLinkedList(n0, nilVal)
	if emptyNode == nil {
		fmt.Printf("查找不存在值 %d，返回 nil\n", nilVal)
	}
}
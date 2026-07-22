package main

import (
	"fmt"
	"shihua_algo/search"
)

func main() {
	fmt.Println("===== 1. 哈希查找：数组索引映射测试 =====")
	// 构造 value -> index 哈希表
	numMap := map[int]int{
		1:  0,
		3:  1,
		5:  2,
		7:  3,
		9:  4,
		11: 5,
	}
	target1 := 7
	idx := search.HashingSearchArray(numMap, target1)
	fmt.Printf("哈希表: %v\n查找值 %d，对应索引: %d\n", numMap, target1, idx)

	targetNo := 99
	idxNo := search.HashingSearchArray(numMap, targetNo)
	fmt.Printf("查找不存在值 %d，返回: %d\n\n", targetNo, idxNo)

	fmt.Println("===== 2. 哈希查找：链表节点映射测试 =====")
	// 创建链表节点
	n0 := search.CreateListNode(10)
	n1 := search.CreateListNode(20)
	n2 := search.CreateListNode(30)
	// 构造 value -> ListNode 哈希表
	nodeMap := map[int]*search.ListNode{
		10: n0,
		20: n1,
		30: n2,
	}
	findVal := 20
	resNode := search.HashingSearchLinkedList(nodeMap, findVal)
	if resNode != nil {
		fmt.Printf("找到值 %d 的链表节点，节点Val=%d\n", findVal, resNode.Val)
	} else {
		fmt.Printf("未找到值 %d 的节点\n", findVal)
	}

	findNilVal := 99
	nilNode := search.HashingSearchLinkedList(nodeMap, findNilVal)
	if nilNode == nil {
		fmt.Printf("查找不存在值 %d，返回 nil\n", findNilVal)
	}
}
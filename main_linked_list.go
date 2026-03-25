package main

import (
	"fmt"
	"shihua_algo/linked_list"
)

func main() {
	fmt.Println("初始化链表")
	/* 初始化链表 1 -> 3 -> 2 -> 5 -> 4 */
	// 初始化操作--初始化各个节点(create)
	n0 := linked_list.CreateNewLinkedListNode(1)
	n1 := linked_list.CreateNewLinkedListNode(3)
	n2 := linked_list.CreateNewLinkedListNode(2)
	n3 := linked_list.CreateNewLinkedListNode(5)
	n4 := linked_list.CreateNewLinkedListNode(4)
	n5 := linked_list.CreateNewLinkedListNode(6)
	// 构建节点之间的引用
	n0.Next = n1
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	fmt.Println(n0.Value, n1.Value, n2.Value, n3.Value, n4.Value)
	// 插入操作--在n3和n4之间加入n5(create)
	linked_list.InterNewLinkedListNode(n3, n5)
	fmt.Println(n0.Value, n1.Value, n2.Value, n3.Value, n4.Value, n3.Next.Value)
	// 删除操作--删除n5(delete)
	linked_list.RemoveLinkedListNode(n3)
	fmt.Println(n0.Value, n1.Value, n2.Value, n3.Value, n4.Value, n3.Next.Value)
	// 访问操作--访问索引为3的节点,从n0开始遍历(read,update)
	access_node := linked_list.AccessLinkedListNode(n0, 3)
	access_node.Value = 55
	fmt.Println(access_node.Value, n3.Value)
	// 查找操作--在链表中查找值为名目标值的首个节点,从n0开始遍历(read)
	query_index := linked_list.FindLinkedListNode(n0, 4)
	fmt.Println(query_index)
}

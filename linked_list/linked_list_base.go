/*
这是一个链表相关模块
*/
package linked_list

/*
载入依赖包
*/

/*
基本数据结构--单向链表
技术实现--指针
*/
// 链表节点结构体--LinkedListNode
type LinkedListNode struct {
	Value int             // 节点值
	Next  *LinkedListNode // 指向下一节点的指针
}

// 创建新的链表节点--CreateNewLinkedListNode(create)
func CreateNewLinkedListNode(value int) *LinkedListNode {
	return &LinkedListNode{ // 返回节点地址
		Value: value,
		Next:  nil,
	}
}

// 在执行链表节点后插入新的链表节点--InterNewLinkedListNode(create)
func InterNewLinkedListNode(target_node *LinkedListNode, new_node *LinkedListNode) {
	latter_node := target_node.Next // 确定目标节点的后一个节点
	new_node.Next = latter_node     // 调整新节点的指针，指向目标节点的后一个节点
	target_node.Next = new_node     // 调整目标节点的指针，指向新的插入节点
}

// 删除链表的目标节点之后的首个节点--RemoveLinkedListNode(delete)
func RemoveLinkedListNode(target_node *LinkedListNode) {
	if target_node.Next == nil { // 确定目标节点是否是尾节点，即其后一个节点是否存在
		return
	}
	RemovedNode := target_node.Next // 确定目标节点的后一个节点，即需要删除的节点
	latter_node := RemovedNode.Next // 确定需要删除节点后的节点
	target_node.Next = latter_node  // 调整目标节点的指针，指向需要删除节点后的节点
}

// 根据索引访问链表中的节点--AccessLinkedListNode(read,update)
func AccessLinkedListNode(head *LinkedListNode, index int) *LinkedListNode {
	for i := 0; i < index; i++ { // 从头开始遍历索引
		if head == nil { // 判断索引是否已经到达尾部
			return nil
		}
		head = head.Next // 在n-1轮后，获得目标索引的节点
	}
	return head
}

// 查找值为目标值的首个节点(read)
func FindLinkedListNode(head *LinkedListNode, target_value int) int {
	index := 0
	for head != nil { // 判断索引是否已经达到尾部
		if head.Value == target_value { // 对比目标值
			return index // 满足条件则返回索引
		}
		head = head.Next // 不满足条件则继续遍历
		index++
	}
	return -1 // 发生错误则返回-1
}

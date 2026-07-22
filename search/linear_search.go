/*
这是一个线性查找相关模块
*/
package search

// 线性查找（数组）
func LinearSearchArray(nums []int, target int) int {
	// 遍历数组
	for i := 0; i < len(nums); i++ {
		// 找到目标元素，返回其索引
		if nums[i] == target {
			return i
		}
	}
	// 未找到目标元素，返回 -1
	return -1
}

// 线性查找（链表）
func LinearSearchLinkedList(node *ListNode, target int) *ListNode {
	// 遍历链表
	for node != nil {
		// 找到目标节点，返回之
		if node.Val == target {
			return node
		}
		node = node.Next
	}
	// 未找到目标元素，返回 nil
	return nil
}

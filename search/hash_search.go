/*
这是一个哈希查找相关模块
*/
package search

// 哈希查找（数组）
// HashingSearchArray 哈希查找-数组映射表
func HashingSearchArray(m map[int]int, target int) int {
	// 哈希表的 key: 目标元素，value: 索引
	// 若哈希表中无此 key ，返回 -1
	if index, ok := m[target]; ok {
		return index
	}
	return -1
}

// 哈希查找（链表）
// HashingSearchLinkedList 哈希查找-链表节点映射表
func HashingSearchLinkedList(m map[int]*ListNode, target int) *ListNode {
	// 哈希表的 key: 目标节点值，value: 节点对象
	// 若哈希表中无此 key ，返回 nil
	if node, ok := m[target]; ok {
		return node
	}
	return nil
}

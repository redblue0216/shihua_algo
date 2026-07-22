/*
这是一个搜索算法辅助功能相关模块
*/
package search

// ListNode 单向链表节点，对标 hello-algo/pkg
type ListNode struct {
	Val  int
	Next *ListNode
}

// CreateListNode 新建链表节点（对外暴露）
func CreateListNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

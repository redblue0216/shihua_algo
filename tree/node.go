/*
这是一个树辅助功能相关模块
*/
package tree

import "fmt"


// TreeNode 通用二叉树节点（带高度，兼容AVL/BST）
type TreeNode struct {
	Val    int
	Height int
	Left   *TreeNode
	Right  *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val:    val,
		Height: 0,
	}
}

// PrintTree 层序打印二叉树
func PrintTree(root *TreeNode) {
	if root == nil {
		fmt.Println("空二叉树")
		return
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			fmt.Printf("%d ", node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		fmt.Println()
	}
}

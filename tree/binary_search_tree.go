/*
这是一个二叉搜索树相关模块
*/
package tree


/*
基本数据结构--二叉搜索树
技术实现--链表二叉树
*/
type BinarySearchTree struct {
	root *TreeNode
}

// NewBinarySearchTree 构造二叉搜索树
func NewBinarySearchTree() *BinarySearchTree {
	bst := &BinarySearchTree{}
	bst.root = nil
	return bst
}

// GetRoot 获取根节点
func (bst *BinarySearchTree) GetRoot() *TreeNode {
	return bst.root
}

// Search 查找节点
func (bst *BinarySearchTree) Search(num int) *TreeNode {
	node := bst.root
	for node != nil {
		if node.Val < num {
			node = node.Right
		} else if node.Val > num {
			node = node.Left
		} else {
			break
		}
	}
	return node
}

// Insert 插入节点
func (bst *BinarySearchTree) Insert(num int) {
	cur := bst.root
	if cur == nil {
		bst.root = NewTreeNode(num)
		return
	}
	var pre *TreeNode = nil
	for cur != nil {
		if cur.Val == num {
			return
		}
		pre = cur
		if cur.Val < num {
			cur = cur.Right
		} else {
			cur = cur.Left
		}
	}
	node := NewTreeNode(num)
	if pre.Val < num {
		pre.Right = node
	} else {
		pre.Left = node
	}
}

// Remove 删除节点
func (bst *BinarySearchTree) Remove(num int) {
	cur := bst.root
	if cur == nil {
		return
	}
	var pre *TreeNode = nil
	for cur != nil {
		if cur.Val == num {
			break
		}
		pre = cur
		if cur.Val < num {
			cur = cur.Right
		} else {
			cur = cur.Left
		}
	}
	if cur == nil {
		return
	}

	if cur.Left == nil || cur.Right == nil {
		var child *TreeNode = nil
		if cur.Left != nil {
			child = cur.Left
		} else {
			child = cur.Right
		}
		if cur != bst.root {
			if pre.Left == cur {
				pre.Left = child
			} else {
				pre.Right = child
			}
		} else {
			bst.root = child
		}
	} else {
		tmp := cur.Right
		for tmp.Left != nil {
			tmp = tmp.Left
		}
		bst.Remove(tmp.Val)
		cur.Val = tmp.Val
	}
}

// Print 打印二叉树
func (bst *BinarySearchTree) Print() {
	PrintTree(bst.root)
}

// 中序遍历（BST中序天然升序）
func (bst *BinarySearchTree) InOrder() []int {
	var res []int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		res = append(res, node.Val)
		dfs(node.Right)
	}
	dfs(bst.root)
	return res
}

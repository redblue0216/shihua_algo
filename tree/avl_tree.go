/*
这是一个AVL二叉树相关模块
*/
package tree


/*
基本数据结构--AVL二叉树
技术实现--AVL二叉树
*/
// AVL 树
type AVLTree struct {
	// 根节点
	root *TreeNode
}

// NewAVLTree 构造AVL树
func NewAVLTree() *AVLTree {
	return &AVLTree{root: nil}
}

// GetRoot 获取根节点
func (t *AVLTree) GetRoot() *TreeNode {
	return t.root
}

// 获取节点高度
func (t *AVLTree) Height(node *TreeNode) int {
	// 空节点高度为 -1 ，叶节点高度为 0
	if node != nil {
		return node.Height
	}
	return -1
}

// 更新节点高度
func (t *AVLTree) UpdateHeight(node *TreeNode) {
	lh := t.Height(node.Left)
	rh := t.Height(node.Right)
	// 节点高度等于最高子树高度 + 1
	if lh > rh {
		node.Height = lh + 1
	} else {
		node.Height = rh + 1
	}
}

// 获取平衡因子
func (t *AVLTree) BalanceFactor(node *TreeNode) int {
	// 空节点平衡因子为 0
	if node == nil {
		return 0
	}
	// 节点平衡因子 = 左子树高度 - 右子树高度
	return t.Height(node.Left) - t.Height(node.Right)
}

// 右旋操作
func (t *AVLTree) RightRotate(node *TreeNode) *TreeNode {
	child := node.Left
	grandChild := child.Right
	// 以 child 为原点，将 node 向右旋转
	child.Right = node
	node.Left = grandChild
	// 更新节点高度
	t.UpdateHeight(node)
	t.UpdateHeight(child)
	// 返回旋转后子树的根节点
	return child
}

// 左旋操作
func (t *AVLTree) LeftRotate(node *TreeNode) *TreeNode {
	child := node.Right
	grandChild := child.Left
	// 以 child 为原点，将 node 向左旋转
	child.Left = node
	node.Right = grandChild
	// 更新节点高度
	t.UpdateHeight(node)
	t.UpdateHeight(child)
	// 返回旋转后子树的根节点
	return child
}

// 执行旋转操作，使该子树重新恢复平衡
func (t *AVLTree) Rotate(node *TreeNode) *TreeNode {
	// 获取节点 node 的平衡因子
	bf := t.BalanceFactor(node)
	// 左偏树
	if bf > 1 {
		if t.BalanceFactor(node.Left) >= 0 {
			// 右旋
			return t.RightRotate(node)
		} else {
			// 先左旋后右旋
			node.Left = t.LeftRotate(node.Left)
			return t.RightRotate(node)
		}
	}
	// 右偏树
	if bf < -1 {
		if t.BalanceFactor(node.Right) <= 0 {
			// 左旋
			return t.LeftRotate(node)
		} else {
			// 先右旋后左旋
			node.Right = t.RightRotate(node.Right)
			return t.LeftRotate(node)
		}
	}
	// 平衡树，无须旋转，直接返回
	return node
}

// 插入节点
func (t *AVLTree) Insert(val int) {
	t.root = t.insertHelper(t.root, val)
}

// 递归插入节点（辅助函数
func (t *AVLTree) insertHelper(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return NewTreeNode(val)
	}
	/* 1. 查找插入位置并插入节点 */
	if val < node.Val {
		node.Left = t.insertHelper(node.Left, val)
	} else if val > node.Val {
		node.Right = t.insertHelper(node.Right, val)
	} else {
		// 重复节点不插入，直接返回
		return node
	}
	// 更新节点高度
	t.UpdateHeight(node)
	/* 2. 执行旋转操作，使该子树重新恢复平衡 */
	node = t.Rotate(node)
	// 返回子树的根节点
	return node
}

// 删除节点
func (t *AVLTree) Remove(val int) {
	t.root = t.removeHelper(t.root, val)
}

// 递归删除节点（辅助函数
func (t *AVLTree) removeHelper(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return nil
	}
	/* 1. 查找节点并删除 */
	if val < node.Val {
		node.Left = t.removeHelper(node.Left, val)
	} else if val > node.Val {
		node.Right = t.removeHelper(node.Right, val)
	} else {
		if node.Left == nil || node.Right == nil {
			child := node.Left
			if node.Right != nil {
				child = node.Right
			}
			if child == nil {
				// 子节点数量 = 0 ，直接删除 node 并返回
				return nil
			} else {
				// 子节点数量 = 1 ，直接删除 node
				node = child
			}
		} else {
			// 子节点数量 = 2 ，则将中序遍历的下个节点删除，并用该节点替换当前节点
			temp := node.Right
			for temp.Left != nil {
				temp = temp.Left
			}
			node.Right = t.removeHelper(node.Right, temp.Val)
			node.Val = temp.Val
		}
	}
	// 更新节点高度
	t.UpdateHeight(node)
	/* 2. 执行旋转操作，使该子树重新恢复平衡 */
	node = t.Rotate(node)
	// 返回子树的根节点
	return node
}

// 查找节点
func (t *AVLTree) Search(val int) *TreeNode {
	cur := t.root
	// 循环查找，越过叶节点后跳出
	for cur != nil {
		if cur.Val < val {
			// 目标节点在 cur 的右子树中
			cur = cur.Right
		} else if cur.Val > val {
			// 目标节点在 cur 的左子树中
			cur = cur.Left
		} else {
			// 找到目标节点，跳出循环
			break
		}
	}
	// 返回目标节点
	return cur
}

// InOrder 中序遍历（AVL同BST，升序）
func (t *AVLTree) InOrder() []int {
	var res []int
	var dfs func(n *TreeNode)
	dfs = func(n *TreeNode) {
		if n == nil {
			return
		}
		dfs(n.Left)
		res = append(res, n.Val)
		dfs(n.Right)
	}
	dfs(t.root)
	return res
}

func (t *AVLTree) Print() {
	PrintTree(t.root)
}
